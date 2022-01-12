// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2022-01-12 14:56
// version: 1.0.0
// desc   :

package concurrent

import (
	"fmt"
	"sync"
	"testing"
)

func TestLocal(t *testing.T) {
	var wg sync.WaitGroup
	gls := NewLocal()
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(idx int) {
			defer wg.Done()
			defer gls.Clear()

			defer func() {
				fmt.Printf("%d: number = %d\n", idx, gls.Get("number"))
			}()
			gls.Put("number", idx+100)
		}(i)
	}
	wg.Wait()
}
