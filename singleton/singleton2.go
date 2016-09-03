package singleton

import (
	"sync"
	"time"
)

type Singleton2 struct {
	CreateTime time.Time
}

var inst *Singleton2 = nil
var mu sync.Mutex

func NewInstance() *Singleton2 {
	if inst == nil {
		mu.Lock()
		defer mu.Unlock()
		if inst == nil {
			inst = &Singleton2{time.Now()}
		}
	}
	return inst
}
