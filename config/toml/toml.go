// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 11:05
// version: 1.0.0
// desc   : toml文件读取

package toml

import (
	"github.com/BurntSushi/toml"
	"github.com/yhyzgn/goat/config"
)

// Reader Toml文件读取器
type Reader struct {
	config.AbsReader
}

// NewReader 创建一个文件读取器
func NewReader() *Reader {
	return new(Reader)
}

// Decode 文件读取解码
func (r *Reader) Decode(filename string, value interface{}) (err error) {
	bs, err := r.Read(filename)
	if err != nil {
		return err
	}
	return r.Parse(bs, value)
}

// Parse 解析数据
func (r *Reader) Parse(data []byte, value interface{}) (err error) {
	_, err = toml.Decode(string(data), value)
	return err
}
