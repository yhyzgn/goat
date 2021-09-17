package crypto

import (
	"testing"
)

const (
	aesKey = "a80a28b4f10c4794ae32815b47bc50e7"
	//aesKey = "a80a28b4f10c4794"
	aesSrc = "李万姬"
	aesIV  = "c6320b4f17334fbd"
)

func TestNewAES(t *testing.T) {
	aes := NewAES()

	encrypted, _ := aes.Encrypt(aesKey, aesSrc)
	t.Log("ECB Base64Encoder - encrypt: " + encrypted)
	decrypted, _ := aes.Decrypt(aesKey, encrypted)
	t.Log("ECB Base64Decoder - decrypt: " + decrypted)
	t.Log("------------------------------------")

	aes = NewAES()
	aes.ModeCBC()
	aes.PaddingPKCS7()
	aes.Encoder(NewBase64URLEncoder())
	aes.Decoder(NewBase64URLDecoder())
	encrypted, _ = aes.EncryptWithIV(aesKey, aesSrc, aesIV)
	t.Log("CBC Base64Encoder - encrypt: " + encrypted)
	decrypted, _ = aes.DecryptWithIV(aesKey, encrypted, aesIV)
	t.Log("CBC Base64Decoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	aes = NewAES()
	aes.ModeECB()
	aes.Encoder(NewHexEncoder())
	aes.Decoder(NewHexDecoder())
	encrypted, _ = aes.Encrypt(aesKey, aesSrc)
	t.Log("ECB HexEncoder - encrypt: " + encrypted)
	decrypted, _ = aes.Decrypt(aesKey, encrypted)
	t.Log("ECB HexDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	aes = NewAES()
	aes.ModeCTR()
	aes.PaddingNo()
	aes.Encoder(NewBase64URLEncoder())
	aes.Decoder(NewBase64URLDecoder())
	encrypted, _ = aes.EncryptWithIV(aesKey, aesSrc, aesIV)
	t.Log("CTR NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = aes.DecryptWithIV(aesKey, encrypted, aesIV)
	t.Log("CTR NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	aes = NewAES()
	aes.ModeOFB()
	aes.PaddingNo()
	aes.Encoder(NewBase64URLEncoder())
	aes.Decoder(NewBase64URLDecoder())
	encrypted, _ = aes.EncryptWithIV(aesKey, aesSrc, aesIV)
	t.Log("OFB NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = aes.DecryptWithIV(aesKey, encrypted, aesIV)
	t.Log("OFB NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	aes = NewAES()
	aes.ModeCFB()
	aes.PaddingNo()
	aes.Encoder(NewBase64URLEncoder())
	aes.Decoder(NewBase64URLDecoder())
	encrypted, _ = aes.EncryptWithIV(aesKey, aesSrc, aesIV)
	t.Log("CFB NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = aes.DecryptWithIV(aesKey, encrypted, aesIV)
	t.Log("CFB NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")
}
