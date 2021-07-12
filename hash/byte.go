package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

// MD5Byte 获取字节数组 md5 值
func MD5Byte(s []byte) string {
	h := md5.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA1Byte 获取节数组 sha1 值
func SHA1Byte(s []byte) string {
	h := sha1.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA256Byte 获取节数组 sha256 值
func SHA256Byte(s []byte) string {
	h := sha256.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}

// SHA512Byte 获取节数组 sha512 值
func SHA512Byte(s []byte) string {
	h := sha512.New()
	h.Write(s)
	return hex.EncodeToString(h.Sum(nil))
}
