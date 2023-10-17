package satori

import "github.com/RomiChan/websocket"

// Client satori 客户端.
type Client struct {
	api      string
	token    string
	platform string
	selfID   string

	ws     *websocket.Conn
	cancel chan bool
}

// NewClient 创建一个 satori 客户端.
func NewClient(api, token string) *Client {
	return &Client{api: api, token: token}
}

// Platform 获取当前 satori platform.
func (cli *Client) Platform() string {
	return cli.platform
}

// SelfID 获取当前 satori self_id.
func (cli *Client) SelfID() string {
	return cli.selfID
}
