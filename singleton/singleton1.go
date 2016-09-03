package singleton

import (
	"sync"
	"time"
)

type Singleton1 struct {
	CreateTime time.Time
}

var instance *Singleton1 = nil
var once sync.Once

func New() *Singleton1 {
	once.Do(func() {
		instance = new(Singleton1)
		instance.CreateTime = time.Now()
	})
	return instance
}
