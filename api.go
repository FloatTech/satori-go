package satori

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type GuildList List[Channel]

func (cli *Client) call(endpoint string, request any, response any) (err error) {
	var buf *bytes.Buffer
	if request != nil {
		buf = bytes.NewBuffer(make([]byte, 0))
		ec := json.NewEncoder(buf)
		ec.Encode(request)
	}
	req, err := http.NewRequest("POST", cli.api+endpoint, buf)
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+cli.token)
	req.Header.Set("X-Platform", cli.platform)
	req.Header.Set("X-Self-ID", cli.selfID)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New("http: " + resp.Status)
	}
	if response != nil {
		dc := json.NewDecoder(resp.Body)
		err = dc.Decode(response)
	}
	return
}

func (cli *Client) GetChannel(channelID string) (channel *Channel, err error) {
	err = cli.call("/v1/channel.get",
		&map[string]any{"channel_id": channelID}, channel,
	)
	return
}

func (cli *Client) ListChannel(guildID, next string) (channels *List[Channel], err error) {
	err = cli.call("/v1/channel.list",
		map[string]any{"guild_id": guildID, "next": next}, channels)
	return
}

func (cli *Client) CreateChannel(guildID string, data *Channel) (channel *Channel, err error) {
	err = cli.call("/v1/channel.create",
		map[string]any{"guild_id": guildID, "data": channel}, channel)
	return
}

func (cli *Client) UpdateChannel(channelID string, data *Channel) (err error) {
	err = cli.call("/v1/channel.update",
		map[string]any{"channel_id": channelID, "data": data}, nil)
	return
}

func (cli *Client) DeleteChannel(channelID string) (err error) {
	err = cli.call("/v1/channel.delete",
		map[string]any{"channel_id": channelID}, nil)
	return
}

func (cli *Client) CreateUserChannel(userID string) (channel *Channel, err error) {
	err = cli.call("/v1/user.channel.create",
		map[string]any{"user_id": userID}, channel)
	return
}

func (cli *Client) GetGuild(guildID string) (guild *Guild, err error) {
	err = cli.call("/v1/guild.get",
		map[string]any{"guild_id": guildID}, guild)
	return
}

func (cli *Client) ListGuild(next string) (guilds *List[Guild], err error) {
	err = cli.call("/v1/guild.list",
		map[string]any{"next": next}, guilds)
	return
}

func (cli *Client) ApproveGuild(messageID string, approve bool, comment string) (err error) {
	err = cli.call("/v1/guild.approve",
		map[string]any{"message_id": messageID, "approve": approve, "comment": comment}, nil)
	return
}

func (cli *Client) GetGuildMember(guildID, userID string) (member *GuildMember, err error) {
	err = cli.call("/v1/guild.member.get",
		map[string]any{"guild_id": guildID, "user_id": userID}, member)
	return
}

func (cli *Client) ListGuildMember(guildID, next string) (members *List[GuildMember], err error) {
	err = cli.call("/v1/guild.member.list",
		map[string]any{"guild_id": guildID, "next": next}, members)
	return
}

func (cli *Client) KickGuildMember(guildID, userID string, permanent bool) (err error) {
	err = cli.call("/v1/guild.member.kick",
		map[string]any{"guild_id": guildID, "user_id": userID, "permanent": permanent}, nil)
	return
}

func (cli *Client) ApproveGuildMember(messageID string, approve bool, comment string) (err error) {
	err = cli.call("/v1/guild.member.approve",
		map[string]any{"message_id": messageID, "approve": approve, "comment": comment}, nil)
	return
}

func (cli *Client) SetGuildMemberRole(guildID, userID, roleID string) (err error) {
	err = cli.call("/v1/guild.member.role.set",
		map[string]any{"guild_id": guildID, "user_id": userID, "role_id": roleID}, nil)
	return
}

func (cli *Client) UnsetGuildMemberRole(guildID, userID, roleID string) (err error) {
	err = cli.call("/v1/guild.member.role.unset",
		map[string]any{"guild_id": guildID, "user_id": userID, "role_id": roleID}, nil)
	return
}

func (cli *Client) ListGuildRole(guildID, next string) (roles *List[GuildRole], err error) {
	err = cli.call("/v1/guild.role.list",
		map[string]any{"guild_id": guildID, "next": next}, roles)
	return
}

func (cli *Client) CreateGuildRole(guildID string, role *GuildRole) (role2 *GuildRole, err error) {
	err = cli.call("/v1/guild.role.create",
		map[string]any{"guild_id": guildID, "role": role}, role2)
	return
}

func (cli *Client) UpdateGuildRole(guildID, roleID string, role *GuildRole) (err error) {
	err = cli.call("/v1/guild.role.update",
		map[string]any{"guild_id": guildID, "role_id": roleID, "role": role}, nil)
	return
}

func (cli *Client) DeleteGuildRole(guildID, roleID string) (err error) {
	err = cli.call("/v1/guild.role.delete",
		map[string]any{"guild_id": guildID, "role_id": roleID}, nil)
	return
}

func (cli *Client) GetLogin() (login *Login, err error) {
	err = cli.call("/v1/login.get",
		nil, login)
	return
}

func (cli *Client) GetMessage(channelID, messageID string) (message *Message, err error) {
	err = cli.call("/v1/message.get",
		map[string]any{"channel_id": channelID, "message_id": messageID}, message)
	return
}

func (cli *Client) CreateMessage(channelID, content string) (messages []Message, err error) {
	err = cli.call("/v1/message.create",
		map[string]any{"channel_id": channelID, "content": content}, &messages)
	return
}

func (cli *Client) DeleteMessage(channelID, messageID string) (err error) {
	err = cli.call("/v1/message.delete",
		map[string]any{"channel_id": channelID, "message_id": messageID}, nil)
	return
}

func (cli *Client) UpdateMessage(channelID, messageID, content string) (err error) {
	err = cli.call("/v1/message.update",
		map[string]any{"channel_id": channelID, "message_id": messageID, "content": content}, nil)
	return
}

func (cli *Client) ListMessage(channelID, next string) (messages *List[Message], err error) {
	err = cli.call("/v1/message.list",
		map[string]any{"channel_id": channelID, "next": next}, messages)
	return
}

func (cli *Client) CreateReaction(channelID, messageID, emoji string) (err error) {
	err = cli.call("/v1/reaction.create",
		map[string]any{"channel_id": channelID, "message_id": messageID, "emoji": emoji}, nil)
	return
}

func (cli *Client) DeleteReaction(channelID, messageID, emoji, userID string) (err error) {
	err = cli.call("/v1/reaction.delete",
		map[string]any{"channel_id": channelID, "message_id": messageID, "emoji": emoji, "user_id": userID}, nil)
	return
}

func (cli *Client) ClearReaction(channelID, messageID, emoji string) (err error) {
	err = cli.call("/v1/reaction.clear",
		map[string]any{"channel_id": channelID, "message_id": messageID, "emoji": emoji}, nil)
	return
}

func (cli *Client) ListReaction(channelID, messageID, emoji, next string) (users *List[User], err error) {
	err = cli.call("/v1/reaction.list",
		map[string]any{"channel_id": channelID, "message_id": messageID, "emoji": emoji, "next": next}, users)
	return
}

func (cli *Client) GetUser(userID string) (user *User, err error) {
	err = cli.call("/v1/user.get",
		map[string]any{"user_id": userID}, user)
	return
}

func (cli *Client) ListFriend(next string) (users *List[User], err error) {
	err = cli.call("/v1/friend.list",
		map[string]any{"next": next}, users)
	return
}

func (cli *Client) ApproveFriend(messageID string, approve bool, comment string) (err error) {
	err = cli.call("/v1/friend.approve",
		map[string]any{"message_id": messageID, "approve": approve, "comment": comment}, nil)
	return
}
