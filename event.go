package satori

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/RomiChan/websocket"
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

func (cli *Client) connect() {
	log.Infof("start connecting to the satori: %s", cli.api)
	network, address := resolveURI(cli.api + "/v1/events")
	dialer := websocket.Dialer{
		NetDial: func(_, addr string) (net.Conn, error) {
			if network == "unix" {
				host, _, err := net.SplitHostPort(addr)
				if err != nil {
					host = addr
				}
				filepath, err := base64.RawURLEncoding.DecodeString(host)
				if err == nil {
					addr = string(filepath)
				}
			}
			return net.Dial(network, addr) // support unix socket transport
		},
	}
	identify, _ := json.Marshal(
		&Signal[Identify]{
			Op:   OpCodeIdentify,
			Body: Identify{Token: cli.token},
		},
	)
	for {
		var ready = Signal[Ready]{}
		conn, res, err := dialer.Dial(address, nil) // nolint
		if err != nil {
			goto ERROR
		}
		_ = res.Body.Close()
		if err = conn.WriteMessage(websocket.TextMessage, identify); err != nil {
			_ = conn.Close()
			goto ERROR
		}
		if err = conn.ReadJSON(&ready); err != nil {
			_ = conn.Close()
			goto ERROR
		}
		if ready.Op != OpCodeReady || len(ready.Body.Logins) == 0 {
			err = errors.New("unknown satori ready signal")
			_ = conn.Close()
			goto ERROR
		}
		cli.platform = ready.Body.Logins[0].Platform
		cli.selfID = ready.Body.Logins[0].SelfID
		log.Infof("successfully connected to satori: %s, platform: %s, self_id: %s",
			cli.api, cli.platform, cli.selfID)
		cli.ws = conn
		go cli.doheartbeat()
		break
	ERROR:
		log.Warnf("failed to connect to satori: %s %v", cli.api, err)
		time.Sleep(5 * time.Second) // 等待两秒后重新连接
	}
}

// Listen 监听 satori 事件.
func (cli *Client) Listen(handler func(*Event)) {
	cli.connect()
	for {
		t, payload, err := cli.ws.ReadMessage()
		if err != nil { // reconnect
			log.Warnf("lost connection to satori: %s %v", cli.api, err)
			cli.cancel <- true
			time.Sleep(time.Millisecond * time.Duration(3))
			cli.connect()
			continue
		}
		if t != websocket.TextMessage {
			continue
		}
		rsp := gjson.ParseBytes(payload)
		if !rsp.Get("op").Exists() {
			continue
		}
		var event = Signal[Event]{}
		switch rsp.Get("op").Int() {
		case OpCodeEvent:
			err = json.Unmarshal(payload, &event)
			if err != nil {
				continue
			}
			handler(&event.Body)
		case OpCodePing, OpCodePong, OpCodeIdentify, OpCodeReady:
			//
		default:
			//
		}
	}
}

func (cli *Client) doheartbeat() {
	var ping = []byte(`{"op": 1}`)
	for {
		select {
		case <-time.After(time.Duration(5) * time.Second):
			err := cli.ws.WriteMessage(websocket.TextMessage, ping)
			if err != nil {
				log.Warnf("an error occurred while sending heartbeat to satori: %v", err)
			}
		case <-cli.cancel:
			return
		}
	}
}

func resolveURI(addr string) (network, address string) {
	network, address = "tcp", addr
	uri, err := url.Parse(addr)
	if err == nil && uri.Scheme != "" {
		scheme, ext, _ := strings.Cut(uri.Scheme, "+")
		switch scheme {
		case "http":
			scheme = "ws"
		case "https":
			scheme = "wss"
		}
		if ext != "" {
			network = ext
			if ext == "unix" {
				uri.Host, uri.Path, _ = strings.Cut(uri.Path, ":")
				uri.Host = base64.StdEncoding.EncodeToString([]byte(uri.Host))
			}
		}
		uri.Scheme = scheme // remove `+unix`/`+tcp4`
		address = uri.String()
	}
	return
}
