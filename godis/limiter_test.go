// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-07-12 20:55
// version: 1.0.0
// desc   :

package godis

import (
	"fmt"
	"testing"
	"time"
)

func TestLimiter_AcquireWith(t *testing.T) {
	lmt := Instance.Limiter("test-limiter", 3, time.Minute)

	if lmt.AcquireWith("18313889251") {
		t.Log("成功啦")
	} else {
		fmt.Println("限流啦")
	}
}
