package hash

// MD5String 获取字符串 md5 值
func MD5String(s string) string {
	return MD5Byte([]byte(s))
}

// SHA1String 获取字符串 sha1 值
func SHA1String(s string) string {
	return SHA1Byte([]byte(s))
}

// SHA256String 获取字符串 sha256 值
func SHA256String(s string) string {
	return SHA256Byte([]byte(s))
}

// SHA512String 获取字符串 sha512 值
func SHA512String(s string) string {
	return SHA512Byte([]byte(s))
}
