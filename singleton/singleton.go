package singleton

import (
	"sync"
	"time"
)

type Singleton struct {
	CreateTime time.Time
}

var instance *Singleton = nil
var once sync.Once

func New() *Singleton {
	once.Do(func() {
		instance = new(Singleton)
		instance.CreateTime = time.Now()
	})
	return instance
}
