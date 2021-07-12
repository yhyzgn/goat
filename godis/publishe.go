// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-02 17:36
// version: 1.0.0
// desc   : 主题订阅模式

package godis

import (
	"github.com/gomodule/redigo/redis"
)

// publisher ...
type publisher struct {
	pool  *RedisPool
	topic Key
}

func newPublisher(pool *RedisPool, topic Key) *publisher {
	return &publisher{pool: pool, topic: topic}
}

// Publish ...
func (p *publisher) Publish(data string) error {
	_, err := p.pool.call(p.topic, func(conn redis.Conn, realKey string) (interface{}, error) {
		if e := conn.Send("PUBLISH", realKey, data); nil != e {
			return nil, e
		}
		if e := conn.Flush(); nil != e {
			return nil, e
		}
		return conn.Receive()
	})
	return err
}
