// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-02 18:39
// version: 1.0.0
// desc   :

package godis

import (
	"fmt"
	"testing"
)

func TestPublish(t *testing.T) {
	if nil != Instance {
		if err := Instance.Publisher("topic-name").Publish("data"); nil != err {
			fmt.Println(err)
		}
	}
}
