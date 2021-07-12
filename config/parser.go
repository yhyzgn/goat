// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-06-02 15:45
// version: 1.0.0
// desc   : 解析数据能力

package config

// Parser 解析数据能力
type Parser interface {

	// Parse 解析数据
	Parse(data []byte, value interface{}) (err error)
}
