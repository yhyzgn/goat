package crypto

import (
	"crypto/des"
)

// DES DES对称加密解密
//
// key 和 iv 都必须是8字节
type DES struct {
	Crypt
	keySize int
}

// NewDES 新实例
func NewDES() *DES {
	d := new(DES)
	d.keySize = 8
	d.mode = ECB
	d.padding = PKCS5
	d.encoder = NewBase64Encoder()
	d.decoder = NewBase64Decoder()
	return d
}

// Encrypt 加密
//
// 模式：ECB
func (d *DES) Encrypt(key, src string) (encrypted string, err error) {
	block, err := des.NewCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	encrypted = encrypt(block, src, d.padding, d.encoder)
	return
}

// EncryptWithIV 加密
//
// 模式：CBC CTR CFB OFB
func (d *DES) EncryptWithIV(key, src, iv string) (encrypted string, err error) {
	block, err := des.NewCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	encrypted = encryptWithIV(block, src, iv, d.mode, d.padding, d.encoder)
	return
}

// Decrypt 解密
//
// 模式：ECB
func (d *DES) Decrypt(key, src string) (decrypted string, err error) {
	block, err := des.NewCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	decrypted = decrypt(block, src, d.padding, d.decoder)
	return
}

// DecryptWithIV 解密
//
// 模式：CBC CTR CFB OFB
func (d *DES) DecryptWithIV(key, src, iv string) (decrypted string, err error) {
	block, err := des.NewCipher([]byte(key)[:d.keySize])
	if err != nil {
		return
	}
	decrypted = decryptWithIV(block, src, iv, d.mode, d.padding, d.decoder)
	return
}
