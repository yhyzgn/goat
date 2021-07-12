package crypto

import (
	"crypto/des"
)

// TripleDES 3DES对称加密解密
//
// key 和 iv 都必须是8字节
type TripleDES struct {
	Crypt
	keySize int
}

// NewTripleDES 新实例
func NewTripleDES() *TripleDES {
	d := new(TripleDES)
	d.keySize = 24
	d.mode = ECB
	d.padding = PKCS5
	d.encoder = NewBase64Encoder()
	d.decoder = NewBase64Decoder()
	return d
}

// Encrypt 加密
//
// 模式：ECB
func (d *TripleDES) Encrypt(key, src string) (encrypted string, err error) {
	block, err := des.NewTripleDESCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	encrypted = encrypt(block, src, d.padding, d.encoder)
	return
}

// EncryptWithIV 加密
//
// 模式：CBC CTR CFB OFB
func (d *TripleDES) EncryptWithIV(key, src, iv string) (encrypted string, err error) {
	block, err := des.NewTripleDESCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	encrypted = encryptWithIV(block, src, iv, d.mode, d.padding, d.encoder)
	return
}

// Decrypt 解密
//
// 模式：ECB
func (d *TripleDES) Decrypt(key, src string) (decrypted string, err error) {
	block, err := des.NewTripleDESCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	decrypted = decrypt(block, src, d.padding, d.decoder)
	return
}

// DecryptWithIV 解密
//
// 模式：CBC CTR CFB OFB
func (d *TripleDES) DecryptWithIV(key, src, iv string) (decrypted string, err error) {
	block, err := des.NewTripleDESCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	decrypted = decryptWithIV(block, src, iv, d.mode, d.padding, d.decoder)
	return
}
