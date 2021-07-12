// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2020-11-19 11:05
// version: 1.0.0
// desc   :

package godis

import (
	"testing"
)

func TestNewRedisPool(t *testing.T) {
	//pool := NewRedisPool("localhost", 6379, 0, 20, 0, "root", "aa")
	//// 3min 内 5 个令牌
	//limiter := pool.Limiter("mobile:13313313313", 5, 1*time.Minute)
	//
	//// 模拟持续请求
	//counter := 1
	//for {
	//	if limiter.Acquire() {
	//		fmt.Println(fmt.Sprintf("第%d次令牌获取成功", counter))
	//	} else {
	//		fmt.Println(fmt.Sprintf("第%d次令牌获取失败", counter))
	//	}
	//
	//	// 随便停一下
	//	interval := time.Duration(rand.Intn(6)) * time.Second
	//	fmt.Println(fmt.Sprintf("第%d次后间隔%v", counter, interval))
	//	time.Sleep(interval)
	//	counter++
	//	fmt.Println("===================================================")
	//}

	// 获取一些key
	//fmt.Println(pool.Keys("bb:*"))
}
