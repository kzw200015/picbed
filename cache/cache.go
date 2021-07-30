package cache

import (
	"errors"
	"sync"
	"time"
)

var cache sync.Map

type Cache struct {
	DstPath string
	Time    time.Time
}

func Set(id, dstPath string) {
	cache.Store(id, Cache{
		DstPath: dstPath,
		Time:    time.Now().UTC(),
	})

	time.AfterFunc(time.Minute, func() {
		cache.Delete(id)
	})
}

func Get(id string) (dstPath string, err error) {
	if c, ok := cache.Load(id); ok {
		return c.(Cache).DstPath, nil
	}

	return "", errors.New("id not exist")
}
