// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 11:05
// version: 1.0.0
// desc   : ini文件读取

package toml

import (
	"bytes"

	"github.com/yhyzgn/goat/config"
	"gopkg.in/gcfg.v1"
)

// Reader INI文件读取器
type Reader struct {
	config.AbsReader
}

// NewReader 创建一个文件读取器
func NewReader() *Reader {
	return new(Reader)
}

// Decode 文件读取解码
func (r *Reader) Decode(filename string, value interface{}) (err error) {
	return gcfg.ReadFileInto(value, filename)
}

// Parse 解析数据
func (r *Reader) Parse(data []byte, value interface{}) (err error) {
	return gcfg.ReadInto(value, bytes.NewBuffer(data))
}
