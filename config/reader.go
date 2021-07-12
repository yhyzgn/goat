// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-14 11:09
// version: 1.0.0
// desc   :

package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"github.com/yhyzgn/goat/file"
)

// AbsReader ...
type AbsReader struct {
	Reader
}

// Read 读取文件
func (r *AbsReader) Read(filename string) (data []byte, err error) {
	if filename == "" {
		err = errors.New("filename can not be empty")
		return
	}
	if !file.Exists(filename) {
		err = errors.New("no such config file '" + filename + "'")
		return
	}
	data, err = ioutil.ReadFile(filename)
	return
}
