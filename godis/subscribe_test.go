// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-02 18:09
// version: 1.0.0
// desc   :

package godis

import (
	"fmt"
	"testing"
)

func TestSubscribe(t *testing.T) {
	if nil != Instance {
		sb, err := Instance.Subscriber("topic-name")
		if nil != err {
			fmt.Println(err)
			return
		}

		// 开始订阅
		sb.Subscribe(func(channel string) error {
			t.Log(fmt.Sprintf("成功订阅【%s】频道！", channel))
			return nil
		}, func(channel string, data []byte) error {
			t.Log(fmt.Sprintf("通道 %s 接收到消息 %s", channel, data))
			return nil
		}, func(err error) {
			fmt.Println(err)
		})

		select {}
	}
}
