package crypto

import (
	"bytes"
	"encoding/base64"
	"encoding/hex"
	"strings"
)

// Encoder 编码器
type Encoder interface {
	Encode(src []byte) (encoded []byte)
}

// Decoder 解码器
type Decoder interface {
	Decode(src []byte) (decoded []byte)
}

// Encryptor 加密器
type Encryptor interface {
	Encrypt(key, src string) (encrypted string)

	EncryptWithIV(key, src, iv string) (encrypted string)
}

// Decryptor 解密器
type Decryptor interface {
	Decrypt(key, src string) (decrypted string)

	DecryptWithIV(key, src, iv string) (decrypted string)
}

// Base64Encoder Base64编码器
type Base64Encoder struct {
}

// NewBase64Encoder 新实例
func NewBase64Encoder() *Base64Encoder {
	return new(Base64Encoder)
}

// Encode 编码方法
func (*Base64Encoder) Encode(src []byte) (encoded []byte) {
	encoded = make([]byte, base64.StdEncoding.EncodedLen(len(src)))
	base64.StdEncoding.Encode(encoded, src)
	return
}

// Base64URLEncoder Base64URL编码器
type Base64URLEncoder struct {
}

// NewBase64URLEncoder 新实例
func NewBase64URLEncoder() *Base64URLEncoder {
	return new(Base64URLEncoder)
}

// Encode 编码
func (*Base64URLEncoder) Encode(src []byte) (encoded []byte) {
	encoded = make([]byte, base64.URLEncoding.EncodedLen(len(src)))
	base64.URLEncoding.Encode(encoded, src)
	// 把末尾的 = 去掉吧
	encoded = bytes.ReplaceAll(encoded, []byte("="), []byte(""))
	return
}

// HexEncoder 十六进制编码器
type HexEncoder struct {
}

// NewHexEncoder 新实例
func NewHexEncoder() *HexEncoder {
	return new(HexEncoder)
}

// Encode 编码
func (he *HexEncoder) Encode(src []byte) (encoded []byte) {
	encoded = make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(encoded, src)
	return
}

// Base64Decoder Base64解码器
type Base64Decoder struct {
}

// NewBase64Decoder 新实例
func NewBase64Decoder() *Base64Decoder {
	return new(Base64Decoder)
}

// Decode 解码
func (*Base64Decoder) Decode(src []byte) (decoded []byte) {
	decoded, _ = base64.StdEncoding.DecodeString(string(src))
	return
}

// Base64URLDecoder Base64URL解码器
type Base64URLDecoder struct {
}

// NewBase64URLDecoder 新实例
func NewBase64URLDecoder() *Base64URLDecoder {
	return new(Base64URLDecoder)
}

// Decode 解码
func (*Base64URLDecoder) Decode(src []byte) (decoded []byte) {
	// Go 语言 Base64URL 不处理 ==，所以解码前先判断长度是否为 4 的倍数，填充缺失的 =
	// 如果不是 4 的整数倍字符长度，解码就会报错
	data := string(src)
	padding := 4 - (len(data) % 4)
	data += strings.Repeat("=", padding)
	decoded, _ = base64.URLEncoding.DecodeString(data)
	return
}

// HexDecoder 十六进制解码器
type HexDecoder struct {
}

// NewHexDecoder 新实例
func NewHexDecoder() *HexDecoder {
	return new(HexDecoder)
}

// Decode 解码
func (*HexDecoder) Decode(src []byte) (decoded []byte) {
	decoded, _ = hex.DecodeString(string(src))
	return
}
