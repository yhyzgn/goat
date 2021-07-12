// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 9:46
// version: 1.0.0
// desc   :

package config

// Reader 文件读取器
type Reader interface {
	// Read 读取文件
	Read(filename string) (data []byte, err error)

	// Decode 文件读取解码
	Decode(filename string, value interface{}) (err error)
}
