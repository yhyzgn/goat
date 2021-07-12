package crypto

import (
	"crypto/aes"
)

// AES AES对称加密解密
type AES struct {
	Crypt
}

// NewAES 新实例
func NewAES() *AES {
	a := new(AES)
	a.mode = ECB
	a.padding = PKCS5
	a.encoder = NewBase64Encoder()
	a.decoder = NewBase64Decoder()
	return a
}

// Encrypt 加密
//
// 模式：ECB
func (a *AES) Encrypt(key, src string) (encrypted string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encrypted = encrypt(block, src, a.padding, a.encoder)
	return
}

// EncryptWithIV 加密
//
// 模式：CBC CTR CFB OFB
func (a *AES) EncryptWithIV(key, src, iv string) (encrypted string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	encrypted = encryptWithIV(block, src, iv, a.mode, a.padding, a.encoder)
	return
}

// Decrypt 解密
//
// 模式：ECB
func (a *AES) Decrypt(key, src string) (decrypted string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return
	}
	decrypted = decrypt(block, src, a.padding, a.decoder)
	return
}

// DecryptWithIV 解密
//
// 模式：CBC CTR CFB OFB
func (a *AES) DecryptWithIV(key, src, iv string) (decrypted string, err error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		panic(err)
	}
	decrypted = decryptWithIV(block, src, iv, a.mode, a.padding, a.decoder)
	return
}
