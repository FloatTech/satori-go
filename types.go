package satori

type OpCode = int64

const (
	OpCodeEvent OpCode = iota
	OpCodePing
	OpCodePong
	OpCodeIdentify
	OpCodeReady
)

// Signal defines model for Signal.
type Signal[T Event | Identify | Ready] struct {
	// 信令类型
	Op OpCode `json:"op"`

	// 信令数据
	Body T `json:"body,omitempty"`
}

// Event defines model for Event.
type Event struct {
	// ID 事件 ID
	ID int64 `json:"id"`

	// Type 事件类型
	Type string `json:"type"`

	// Platform 接收者的平台名称
	Platform string `json:"platform"`

	// SelfID 接收者的平台账号
	SelfID string `json:"self_id"`

	// Timestamp 事件的时间戳
	Timestamp int64 `json:"timestamp"`

	// Channel 事件所属的频道
	Channel *Channel `json:"channel,omitempty"`

	// Guild 事件所属的群组
	Guild *Guild `json:"guild,omitempty"`

	// Login 事件的登录信息
	Login *Login `json:"login,omitempty"`

	// Member 事件的目标成员
	Member *GuildMember `json:"member,omitempty"`

	// Message 事件的消息
	Message *Message `json:"message,omitempty"`

	// Operator 事件的操作者
	Operator *User `json:"operator,omitempty"`

	// Role 事件的目标角色
	Role *GuildRole `json:"role,omitempty"`

	// User 事件的目标用户
	User *User `json:"user,omitempty"`
}

type Identify struct {
	Token string `json:"token"`
}

// Ready defines model for Ready.
type Ready struct {
	Logins []Login `json:"logins"`
}

// List defines model for List.
type List[T Channel | Message] struct {
	// Data 数据数组
	Data []T `json:"data"`

	// Next 分页
	Next string `json:"next"`
}

// Channel defines model for Channel.
type Channel struct {
	// Avatar 不安全的频道头像
	Avatar string `json:"avatar"`

	// ID 频道 ID
	ID string `json:"id"`

	// Name 频道名称
	Name string `json:"name"`

	// ParentID 父频道 ID
	ParentID string      `json:"parent_id"`
	Type     ChannelType `json:"type"`
}

// ChannelType defines model for ChannelType.
type ChannelType float32

// Guild defines model for Guild.
type Guild struct {
	// Avatar 不安全的群组头像
	Avatar string `json:"avatar"`

	// ID 群组 ID
	ID string `json:"id"`

	// Name 群组名称
	Name string `json:"name"`
}

// GuildMember defines model for GuildMember.
type GuildMember struct {
	// Avatar 用户在群组中的头像
	Avatar string `json:"avatar,omitempty"`

	// Name 用户在群组中的名称
	Name string `json:"name,omitempty"`
	User *User  `json:"user,omitempty"`
}

// GuildRole defines model for GuildRole.
type GuildRole struct {
	// ID 群组 ID
	ID string `json:"id"`

	// Name 群组名称
	Name string `json:"name"`
}

// Status 登录状态
type Status int8

const (
	StatusOffline    = iota // 离线
	StatusOnline            // 在线
	StatusConnect           // 连接中
	StatusDisconnect        // 断开连接
	StatusReconnect         // 重新连接
)

// Login defines model for Login.
type Login struct {
	// User 用户对象
	User *User `json:"user,omitempty"`

	// SelfID 平台账号
	SelfID string `json:"self_id"`

	// Platform 平台名称
	Platform string `json:"platform,omitempty"`

	// Status 登录状态
	Status Status `json:"status"`
}

// Message defines model for Message.
type Message struct {
	Channel *Channel `json:"channel,omitempty"`

	// Content 消息内容
	Content string `json:"content"`

	// CreatedAt 消息发送的时间戳
	CreatedAt float32 `json:"created_at,omitempty"`
	Guild     *Guild  `json:"guild,omitempty"`

	// ID 消息 ID
	ID     string       `json:"id"`
	Member *GuildMember `json:"member,omitempty"`

	// UpdatedAt 消息修改的时间戳
	UpdatedAt float32 `json:"updated_at,omitempty"`
	User      *User   `json:"user,omitempty"`
}

// MessageCreatePayload defines model for MessageCreatePayload.
type MessageCreatePayload struct {
	// ChannelID 消息要发送到的频道。
	//
	// 在 Chronocat，群聊对应的频道为群号，
	// 私聊对应的频道为 private: 后跟 QQ 号。
	ChannelID string `json:"channel_id"`

	// Content 消息的内容。
	//
	// 格式为 Satori 消息元素字符串。
	Content string `json:"content"`
}

// User defines model for User.
type User struct {
	// Avatar 用户头像
	Avatar string `json:"avatar,omitempty"`

	// ID 用户 ID
	ID string `json:"id"`

	// IsBot 是否为机器人
	IsBot bool `json:"is_bot,omitempty"`

	// Name 用户名称
	Name string `json:"name"`
}
