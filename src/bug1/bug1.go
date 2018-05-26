package bug1

import (
	"time"
	"sync/atomic"
)

type Counter int64

func (c *Counter) Add(x int64) {
	atomic.AddInt64((*int64)(c), 1)
	time.Sleep(time.Nanosecond)
}
