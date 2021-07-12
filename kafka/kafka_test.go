// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-30 下午6:23
// version: 1.0.0
// desc   :

package kafka

import (
	"fmt"
	"testing"
	"time"
)

func TestKfkService_Write(t *testing.T) {
	kfk, err := InitKafka([]string{"172.22.96.1:9092"}, map[KfkTopicName]string{"kgo": "test-kgo"})
	if err != nil {
		fmt.Println(err)
		return
	}
	err = kfk.InitWriter("kgo")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = kfk.Write("Hello gateway kafka ...")
	if err != nil {
		fmt.Println(err)
		return
	}

	time.Sleep(3 * time.Second)
}
