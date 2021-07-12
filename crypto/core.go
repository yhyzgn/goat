package crypto

import (
	"crypto/cipher"

	"github.com/yhyzgn/goat/crypto/mode"
)

// encrypt 加密
//
// 模式：ECB
func encrypt(block cipher.Block, src string, padding Padding, encoder Encoder) (encrypted string) {
	bys := DoPadding(padding, []byte(src), block.BlockSize())
	enc := make([]byte, len(bys))
	ecb := mode.NewECBEncrypter(block)
	ecb.CryptBlocks(enc, bys)
	encrypted = string(encoder.Encode(enc))
	return
}

// encryptWithIV 加密
//
// 模式：CBC CTR CFB OFB
func encryptWithIV(block cipher.Block, src, iv string, md Mode, padding Padding, encoder Encoder) (encrypted string) {
	bys := DoPadding(padding, []byte(src), block.BlockSize())
	enc := make([]byte, len(bys))
	ivBys := []byte(iv)[:block.BlockSize()]
	switch md {
	case CBC:
		cipher.NewCBCEncrypter(block, ivBys).CryptBlocks(enc, bys)
	case CTR:
		cipher.NewCTR(block, ivBys).XORKeyStream(enc, bys)
	case CFB:
		cipher.NewCFBEncrypter(block, ivBys).XORKeyStream(enc, bys)
	case OFB:
		cipher.NewOFB(block, ivBys).XORKeyStream(enc, bys)
	default:
		panic("the mode haven't been supported")
	}
	encrypted = string(encoder.Encode(enc))
	return
}

// decrypt 解密
//
// 模式：ECB
func decrypt(block cipher.Block, src string, padding Padding, decoder Decoder) (decrypted string) {
	bys := decoder.Decode([]byte(src))
	enc := make([]byte, len(bys))
	ecb := mode.NewECBDecrypter(block)
	ecb.CryptBlocks(enc, bys)
	unPadding := UnPadding(padding, enc)
	decrypted = string(unPadding)
	return
}

// decryptWithIV 解密
//
// 模式：CBC CTR CFB OFB
func decryptWithIV(block cipher.Block, src, iv string, md Mode, padding Padding, decoder Decoder) (decrypted string) {
	bys := decoder.Decode([]byte(src))
	enc := make([]byte, len(bys))
	ivBys := []byte(iv)[:block.BlockSize()]
	switch md {
	case CBC:
		cipher.NewCBCDecrypter(block, ivBys).CryptBlocks(enc, bys)
	case CTR:
		cipher.NewCTR(block, ivBys).XORKeyStream(enc, bys)
	case CFB:
		cipher.NewCFBDecrypter(block, ivBys).XORKeyStream(enc, bys)
	case OFB:
		cipher.NewOFB(block, ivBys).XORKeyStream(enc, bys)
	default:
		panic("the mode haven't been supported")
	}
	unPadding := UnPadding(padding, enc)
	decrypted = string(unPadding)
	return
}
