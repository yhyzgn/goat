// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-12 11:00
// version: 1.0.0
// desc   : 协程

package concurrent

import (
	"runtime"
	"strings"
)

// GoroutineId 获取当前协程ID
func GoroutineId() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	stk := strings.TrimPrefix(string(buf[:n]), "goroutine ")
	return strings.Fields(stk)[0]
}
