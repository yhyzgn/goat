// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-04-02 17:36
// version: 1.0.0
// desc   : 主题订阅模式

package godis

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"strings"
)

// subscriber ...
type subscriber struct {
	pool  *RedisPool
	topic []Key
	psc   *redis.PubSubConn
}

// SubCh ...
type SubCh struct {
	Channel string
	Data    []byte
	Err     error
}

func newSubscriber(pool *RedisPool, topic ...Key) (*subscriber, error) {
	if nil == topic {
		return nil, errors.New("请指定要订阅的频道")
	}
	sb := &subscriber{pool: pool, topic: topic}
	err := sb.connect()
	return sb, err
}

func (s *subscriber) connect() error {
	psc := &redis.PubSubConn{Conn: s.pool.Pool.Get()}

	// 可订阅多个频道
	channels := make([]interface{}, len(s.topic))
	for i, tp := range s.topic {
		channels[i] = s.pool.withPrefix(tp)
	}

	if e := psc.Subscribe(channels...); nil != e {
		return e
	}
	s.psc = psc
	return nil
}

// Subscribe 开始订阅
func (s *subscriber) Subscribe(onStart func(channel string) error, onMessage func(channel string, data []byte) error, onError func(err error)) {
	go func() {
		for {
			switch v := s.psc.Receive().(type) {
			case redis.Message:
				if err := onMessage(v.Channel, v.Data); nil != err && nil != onError {
					onError(err)
				}
				continue
			case redis.Subscription:
				switch v.Count {
				case 1:
					if err := onStart(v.Channel); nil != err && nil != onError {
						onError(err)
					}
					continue
				}
			case error:
				err := v
				if strings.Contains(err.Error(), "connection closed") {
					// 断线了，重连
					err = s.connect()
					if nil == err {
						continue
					}
				}
				if nil != err && nil != onError {
					onError(err)
				}
			}
		}
	}()
}

// Unsubscribe 取消订阅
func (s *subscriber) Unsubscribe() error {
	return s.psc.Unsubscribe()
}
