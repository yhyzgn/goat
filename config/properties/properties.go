// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 11:21
// version: 1.0.0
// desc   : properties文件读取

package properties

import (
	"github.com/magiconair/properties"
	"github.com/yhyzgn/goat/config"
)

// Reader Properties文件读取器
type Reader struct {
	config.AbsReader
}

// NewReader 创建一个文件读取器
func NewReader() *Reader {
	return new(Reader)
}

// Decode 文件读取解码
func (r *Reader) Decode(filename string, value interface{}) (err error) {
	p, err := properties.LoadFile(filename, properties.UTF8)
	if nil != err {
		return
	}
	return p.Decode(value)
}

// Parse 解析数据
func (r *Reader) Parse(data []byte, value interface{}) (err error) {
	p, err := properties.Load(data, properties.UTF8)
	if nil != err {
		return
	}
	return p.Decode(value)
}
