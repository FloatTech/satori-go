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
	Channel *Channel `json:"channel"`

	// Guild 事件所属的群组
	Guild *Guild `json:"guild"`

	// Login 事件的登录信息
	Login *Login `json:"login"`

	// Member 事件的目标成员
	Member *GuildMember `json:"member"`

	// Message 事件的消息
	Message *Message `json:"message"`

	// Operator 事件的操作者
	Operator *User `json:"operator"`

	// Role 事件的目标角色
	Role *GuildRole `json:"role"`

	// User 事件的目标用户
	User *User `json:"user"`
}

type Identify struct {
	Token string `json:"token"`
}

// Ready defines model for Ready.
type Ready struct {
	Logins []Login `json:"logins"`
}

// List defines model for List.
type List[T Channel | Guild | GuildMember | GuildRole | Message | User] struct {
	// Data 数据数组
	Data []T `json:"data"`

	// Next 分页
	Next string `json:"next"`
}

// Channel 频道.
type Channel struct {
	// ID 频道 ID
	ID string `json:"id"`

	// Type 频道类型
	Type ChannelType `json:"type"`

	// Name 频道名称
	Name string `json:"name"`

	// ParentID 父频道 ID
	ParentID string `json:"parent_id"`

	// Avatar 频道头像
	Avatar string `json:"avatar"`
}

// ChannelType 频道类型.
type ChannelType = int64

const (
	// ChannelTypeText 文本频道.
	ChannelTypeText = iota

	// ChannelTypeVoice 语音频道.
	ChannelTypeVoice

	// ChannelTypeCategory 分类频道.
	ChannelTypeCategory

	// ChannelTypeDirect 私聊频道.
	ChannelTypeDirect
)

// Guild 群组.
type Guild struct {
	// ID 群组 ID
	ID string `json:"id"`

	// Name 群组名称
	Name string `json:"name"`

	// Avatar 群组头像
	Avatar string `json:"avatar"`
}

// GuildMember 群组成员.
type GuildMember struct {
	// 用户对象
	User *User `json:"user"`

	// Name 用户在群组中的名称
	Name string `json:"name"`

	// Avatar 用户在群组中的头像
	Avatar string `json:"avatar"`

	// JoinAt 加入时间
	JoinAt int64 `json:"joined_at"`
}

// GuildRole 群组角色.
type GuildRole struct {
	// ID 群组 ID
	ID string `json:"id"`

	// Name 群组名称
	Name string `json:"name"`
}

// Login 登录信息.
type Login struct {
	// User 用户对象
	User *User `json:"user"`

	// SelfID 平台账号
	SelfID string `json:"self_id"`

	// Platform 平台名称
	Platform string `json:"platform"`

	// Status 登录状态
	Status LoginStatus `json:"status"`
}

// LoginStatus 登录状态.
type LoginStatus int64

const (
	// LoginStatusOffline 离线.
	LoginStatusOffline = iota
	// LoginStatusOnline 在线.
	LoginStatusOnline
	// LoginStatusConnect 连接中.
	LoginStatusConnect
	// LoginStatusDisconnect 断开连接.
	LoginStatusDisconnect
	// LoginStatusReconnect 重新连接.
	LoginStatusReconnect
)

// Message 消息.
type Message struct {
	// ID 消息 ID
	ID string `json:"id"`

	// Content 消息内容
	Content string `json:"content"`

	// Channel 频道对象
	Channel *Channel `json:"channel"`

	// Guild 群组对象
	Guild *Guild `json:"guild"`

	// Member 成员对象
	Member *GuildMember `json:"member"`

	// User 用户对象
	User *User `json:"user"`

	// CreatedAt 消息发送的时间戳
	CreatedAt int64 `json:"created_at"`

	// UpdatedAt 消息修改的时间戳
	UpdatedAt int64 `json:"updated_at"`
}

// User 用户.
type User struct {
	// ID 用户 ID
	ID string `json:"id"`

	// Name 用户名称
	Name string `json:"name"`

	// Avatar 用户头像
	Avatar string `json:"avatar"`

	// IsBot 是否为机器人
	IsBot bool `json:"is_bot"`
}
