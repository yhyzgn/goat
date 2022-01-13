// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2019-10-30 下午5:47
// version: 1.0.0
// desc   : Kafka 服务

package kafka

import (
	"context"
	"fmt"
	"time"

	"github.com/segmentio/kafka-go"

	"github.com/pkg/errors"
)

// KfkTopicName 话题名称
type KfkTopicName string

// KfkService kafka服务
type KfkService struct {
	hosts  []string
	topics map[KfkTopicName]string
	buff   chan *kafka.Message
	writer *kafka.Writer
}

var (
	service  *KfkService // kafka 服务客户端
	buffSize = 10000     // 缓冲区大小
)

// InitKafka 初始化服务
func InitKafka(hosts []string, topics map[KfkTopicName]string) (*KfkService, error) {
	service = &KfkService{
		hosts:  hosts,
		topics: topics,
		buff:   make(chan *kafka.Message, buffSize),
	}
	go service.start()
	return service, nil
}

// Service 当前服务实例
func Service() (*KfkService, error) {
	if service != nil {
		return service, nil
	}
	return nil, errors.New("invalid kafka service")
}

func (s *KfkService) start() {
	for {
		// 只取有效数据
		ln := len(s.buff)
		if ln > 0 {
			msgList := make([]kafka.Message, 0)
			for i := 0; i < ln; i++ {
				msgList = append(msgList, *(<-s.buff))
			}
			if len(msgList) > 0 {
				s.send(msgList)
			}
		}
		// 3s 发送一次
		time.Sleep(3 * time.Second)
	}
}

func (s *KfkService) send(msgList []kafka.Message) {
	if s.writer != nil {
		if err := s.writer.WriteMessages(context.Background(), msgList...); err != nil {
			fmt.Println("kafka 服务发送数据失败：", err)
		}
	}
}

// InitWriter 初始化输出器
func (s *KfkService) InitWriter(topicName KfkTopicName) (err error) {
	// 匹配 topic
	topic, found := s.topics[topicName]
	if !found {
		err = fmt.Errorf("topic [%s] not found", topicName)
		return
	}
	s.writer = &kafka.Writer{
		Addr:     kafka.TCP(s.hosts...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
	return
}

// Write 输出
func (s *KfkService) Write(msg string) (err error) {
	return s.WriteBytes([]byte(msg))
}

// WriteBytes 输出
func (s *KfkService) WriteBytes(bs []byte) (err error) {
	select {
	case s.buff <- &kafka.Message{Value: bs}:
		break
	default:
		break
	}
	return
}

// Close 关闭
func (s *KfkService) Close() error {
	return nil
}
