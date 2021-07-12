// author : 颜洪毅
// e-mail : yhyzgn@gmail.com
// time   : 2021-05-20 17:30
// version: 1.0.0
// desc   : 雪花算法

package snowflake

import (
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/yhyzgn/gog"
)

const (
	epoch             = int64(1577808000000)                           // 设置起始时间(时间戳/毫秒)：2020-01-01 00:00:00，有效期69年
	timestampBits     = uint(41)                                       // 时间戳占用位数
	dataCenterIDBits  = uint(5)                                        // 数据中心id所占位数
	workerIDBits      = uint(5)                                        // 机器id所占位数
	sequenceBits      = uint(12)                                       // 序列所占的位数
	timestampMax      = int64(-1 ^ (-1 << timestampBits))              // 时间戳最大值
	dataCenterIDMax   = int64(-1 ^ (-1 << dataCenterIDBits))           // 支持的最大数据中心id数量
	workerIDMax       = int64(-1 ^ (-1 << workerIDBits))               // 支持的最大机器id数量
	sequenceMask      = int64(-1 ^ (-1 << sequenceBits))               // 支持的最大序列id数量
	workerIDShift     = sequenceBits                                   // 机器id左移位数
	dataCenterIDShift = sequenceBits + workerIDBits                    // 数据中心id左移位数
	timestampShift    = sequenceBits + workerIDBits + dataCenterIDBits // 时间戳左移位数
)

// Snowflake 雪花算法
type Snowflake struct {
	sync.Mutex
	timestamp    int64
	dataCenterID int64
	workerID     int64
	sequence     int64
}

// NewSnowflake 创建一个对象
func NewSnowflake(dataCenterID, workerID int64) (*Snowflake, error) {
	if dataCenterID < 0 || dataCenterID > dataCenterIDMax {
		return nil, fmt.Errorf("dataCenterID must be between 0 and %d", dataCenterIDMax-1)
	}
	if workerID < 0 || workerID > workerIDMax {
		return nil, fmt.Errorf("workerID must be between 0 and %d", workerIDMax-1)
	}
	return &Snowflake{
		timestamp:    0,
		dataCenterID: dataCenterID,
		workerID:     workerID,
		sequence:     0,
	}, nil
}

// NextStr 获取一个ID
func (sf *Snowflake) NextStr() string {
	return strconv.FormatInt(sf.Next(), 10)
}

// Next 获取一个ID
func (sf *Snowflake) Next() int64 {
	sf.Lock()
	now := time.Now().UnixNano() / 1000000 // 转毫秒
	if sf.timestamp == now {
		// 当同一时间戳（精度：毫秒）下多次生成id会增加序列号
		sf.sequence = (sf.sequence + 1) & sequenceMask
		if sf.sequence == 0 {
			// 如果当前序列超出12bit长度，则需要等待下一毫秒
			// 下一毫秒将使用sequence:0
			for now <= sf.timestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		// 不同时间戳（精度：毫秒）下直接使用序列号：0
		sf.sequence = 0
	}
	t := now - epoch
	if t > timestampMax {
		sf.Unlock()
		gog.ErrorF("epoch must be between 0 and %d", timestampMax-1)
		return 0
	}
	sf.timestamp = now
	r := (t << timestampShift) | (sf.dataCenterID << dataCenterIDShift) | (sf.workerID << workerIDShift) | (sf.sequence)
	sf.Unlock()
	return r
}

// GetDeviceID 获取数据中心ID和机器ID
func GetDeviceID(sid int64) (datacenterid, workerid int64) {
	datacenterid = (sid >> dataCenterIDShift) & dataCenterIDMax
	workerid = (sid >> workerIDShift) & workerIDMax
	return
}

// GetTimestamp 获取时间戳
func GetTimestamp(sid int64) (timestamp int64) {
	timestamp = (sid >> timestampShift) & timestampMax
	return
}

// GetGenTimestamp 获取创建ID时的时间戳
func GetGenTimestamp(sid int64) (timestamp int64) {
	timestamp = GetTimestamp(sid) + epoch
	return
}

// GetGenTime 获取创建ID时的时间字符串(精度：秒)
func GetGenTime(sid int64) (t string) {
	// 需将GetGenTimestamp获取的时间戳/1000转换成秒
	t = time.Unix(GetGenTimestamp(sid)/1000, 0).Format("2006-01-02 15:04:05")
	return
}

// GetTimestampStatus 获取时间戳已使用的占比：范围（0.0 - 1.0）
func GetTimestampStatus() (state float64) {
	state = float64(time.Now().UnixNano()/1000000-epoch) / float64(timestampMax)
	return
}
