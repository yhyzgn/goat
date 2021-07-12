package crypto

import (
	"testing"
)

const (
	desKey = "a80a28b4f10c4794ae32815b47bc50e7"
	desSrc = "18813875621"
	desIV  = "c6320b4fc6320b4f"
)

func TestNewDES(t *testing.T) {
	des := NewDES()

	encrypted, _ := des.Encrypt(desKey, desSrc)
	t.Log("ECB Base64Encoder - encrypt: " + encrypted)
	decrypted, _ := des.Decrypt(desKey, encrypted)
	t.Log("ECB Base64Decoder - decrypt: " + decrypted)
	t.Log("------------------------------------")

	des = NewDES()
	des.ModeCBC()
	encrypted, _ = des.EncryptWithIV(desKey, desSrc, desIV)
	t.Log("CBC Base64Encoder - encrypt: " + encrypted)
	decrypted, _ = des.DecryptWithIV(desKey, encrypted, desIV)
	t.Log("CBC Base64Decoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	des = NewDES()
	des.ModeECB()
	des.Encoder(NewHexEncoder())
	des.Decoder(NewHexDecoder())
	encrypted, _ = des.Encrypt(desKey, desSrc)
	t.Log("ECB HexEncoder - encrypt: " + encrypted)
	decrypted, _ = des.Decrypt(desKey, encrypted)
	t.Log("ECB HexDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	des = NewDES()
	des.ModeCTR()
	des.Encoder(NewBase64URLEncoder())
	des.Decoder(NewBase64URLDecoder())
	encrypted, _ = des.EncryptWithIV(desKey, desSrc, desIV)
	t.Log("CTR NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = des.DecryptWithIV(desKey, encrypted, desIV)
	t.Log("CTR NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	des = NewDES()
	des.ModeOFB()
	des.Encoder(NewBase64URLEncoder())
	des.Decoder(NewBase64URLDecoder())
	encrypted, _ = des.EncryptWithIV(desKey, desSrc, desIV)
	t.Log("OFB NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = des.DecryptWithIV(desKey, encrypted, desIV)
	t.Log("OFB NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")

	des = NewDES()
	des.ModeCFB()
	des.Encoder(NewBase64URLEncoder())
	des.Decoder(NewBase64URLDecoder())
	encrypted, _ = des.EncryptWithIV(desKey, desSrc, desIV)
	t.Log("CFB NewBase64URLEncoder - encrypt: " + encrypted)
	decrypted, _ = des.DecryptWithIV(desKey, encrypted, desIV)
	t.Log("CFB NewBase64URLDecoder - decrypt: " + decrypted)
	t.Log("====================================\n")
}
