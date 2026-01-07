package snowflake

import (
	"sync"
	"time"
)

const (
	workerBits     = 5  // 机器ID位数
	datacenterBits = 5  // 机房ID位数
	sequenceBits   = 12 // 序列号位数

	maxWorkerID     = -1 ^ (-1 << workerBits)
	maxDatacenterID = -1 ^ (-1 << datacenterBits)
	maxSequence     = -1 ^ (-1 << sequenceBits)

	workerShift     = sequenceBits
	datacenterShift = sequenceBits + workerBits
	timestampShift  = sequenceBits + workerBits + datacenterBits

	epoch = int64(1577808000000) // 自定义起始时间（2020-01-01）
)

type Snowflake struct {
	mu           sync.Mutex
	lastStamp    int64
	workerID     int64
	datacenterID int64
	sequence     int64
}

var IdGen *Snowflake

func InitSnowflake(workerID int64, datacenterID int64) {
	IdGen = newSnowflake(workerID, datacenterID)
}

func newSnowflake(workerID, datacenterID int64) *Snowflake {
	if workerID > maxWorkerID || workerID < 0 {
		panic("workerID out of range")
	}
	if datacenterID > maxDatacenterID || datacenterID < 0 {
		panic("datacenterID out of range")
	}
	return &Snowflake{
		workerID:     workerID,
		datacenterID: datacenterID,
	}
}

func (s *Snowflake) NextID() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()

	now := time.Now().UnixNano() / 1e6 // 毫秒

	if now == s.lastStamp {
		s.sequence = (s.sequence + 1) & maxSequence
		if s.sequence == 0 {
			for now <= s.lastStamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		s.sequence = 0
	}

	s.lastStamp = now

	return ((now - epoch) << timestampShift) |
		(s.datacenterID << datacenterShift) |
		(s.workerID << workerShift) |
		s.sequence
}
