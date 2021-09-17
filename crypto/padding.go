package crypto

import "bytes"

// Padding 填充方式
type Padding uint32

// No ...
const (
	No       Padding = 1 << (32 - 1 - iota) // 无填充
	Zero                                    // 数据长度不对齐时使用0填充，否则不填充
	PKCS1                                   //
	PKCS5                                   // PKCS7Padding的子集，块大小固定为8字节
	PKCS7                                   // 假设数据长度需要填充n(n>0)个字节才对齐，那么填充n个字节，每个字节都是n;如果数据本身就已经对齐了，则填充一块长度为块大小的数据，每个字节都是块大小
	ISO10126                                //
	OAEP                                    //
	SSL3                                    //
)

var paddingMap = map[Padding]paddingAble{
	No:    new(NoPadding),
	Zero:  new(ZeroPadding),
	PKCS5: new(PKCS5Padding),
	PKCS7: new(PKCS7Padding),
}

// DoPadding 加密时填充
func DoPadding(padding Padding, src []byte, blockSize int) []byte {
	if pd, ok := paddingMap[padding]; ok {
		return pd.Padding(src, blockSize)
	}
	return nil
}

// UnPadding 解密时去除填充
func UnPadding(padding Padding, src []byte) []byte {
	if pd, ok := paddingMap[padding]; ok {
		return pd.UnPadding(src)
	}
	return nil
}

// ========================================================== paddingAble ==========================================================

type paddingAble interface {
	Padding(src []byte, blockSize int) []byte

	UnPadding(src []byte) []byte
}

// ========================================================== NoPadding ==========================================================

// NoPadding 无填充
type NoPadding struct {
}

// Padding ...
func (*NoPadding) Padding(src []byte, blockSize int) []byte {
	return src
}

// UnPadding ...
func (*NoPadding) UnPadding(src []byte) []byte {
	return src
}

// ========================================================== ZeroPadding ==========================================================

// ZeroPadding ...
type ZeroPadding struct {
}

// Padding ...
func (*ZeroPadding) Padding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	if paddingCount > 0 {
		return append(src, bytes.Repeat([]byte{byte(0)}, paddingCount)...)
	}
	return src
}

// UnPadding ...
func (*ZeroPadding) UnPadding(src []byte) []byte {
	for i := len(src) - 1; i > 0; i-- {
		if src[i] != 0 {
			return src[:i+1]
		}
	}
	return nil
}

// ========================================================== PKCS5Padding ==========================================================

// PKCS5Padding ...
type PKCS5Padding struct {
	PKCS7Padding
}

// Padding ...
func (*PKCS5Padding) Padding(src []byte, blockSize int) []byte {
	// PKCS5 中 blockSize 固定为 8
	return pkcs57Padding(src, blockSize)
}

// ========================================================== PKCS7Padding ==========================================================

// PKCS7Padding ...
type PKCS7Padding struct {
}

// Padding ...
func (*PKCS7Padding) Padding(src []byte, blockSize int) []byte {
	return pkcs57Padding(src, blockSize)
}

// UnPadding ...
func (*PKCS7Padding) UnPadding(src []byte) []byte {
	length := len(src)
	padding := int(src[length-1])
	return src[:(length - padding)]
}

func pkcs57Padding(src []byte, blockSize int) []byte {
	paddingCount := blockSize - len(src)%blockSize
	padding := bytes.Repeat([]byte{byte(paddingCount)}, paddingCount)
	return append(src, padding...)
}
