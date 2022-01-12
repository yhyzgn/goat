// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-12 11:38
// version: 1.0.0
// desc   :

package concurrent

import (
	"testing"
	"time"
)

func TestGoroutineId(t *testing.T) {
	t.Log(GoroutineID())

	for i := 0; i < 10; i++ {
		go func() {
			t.Log(GoroutineID())
		}()
	}

	time.Sleep(time.Second)

	t.Log(GoroutineID())
	t.Log(GoroutineID())
}
