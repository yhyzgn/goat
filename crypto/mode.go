package crypto

// Mode 加密模式
type Mode uint32

// ECB ...
const (
	ECB Mode = 1 << (16 - 1 - iota) // 电子密码本模式
	CBC                             // 密码分组连接模式
	CTR                             // 计数器模式
	OFB                             // 输出反馈模式
	CFB                             // 密文反馈模式
)
