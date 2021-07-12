// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 9:46
// version: 1.0.0
// desc   : yaml文件读取

package yaml

import (
	"bytes"

	"github.com/yhyzgn/goat/config"
	"gopkg.in/yaml.v3"
)

// Reader Yaml文件读取器
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
		return
	}
	return r.Parse(bs, value)
}

// Parse 解析数据
func (r *Reader) Parse(data []byte, value interface{}) (err error) {
	return yaml.NewDecoder(bytes.NewBuffer(data)).Decode(value)
}
