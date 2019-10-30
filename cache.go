package main

import (
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

func NewCache() (c *cache, err error) {
	return &cache{
		m:     make(map[string]*item),
		mutex: new(sync.RWMutex),
	}, nil
}

type item struct {
	content []byte
	expire  int64 // 过期时间

	// todo：应该把响应的header也缓存起来
}

// 一个极其简单的缓存器
// 数据缓存在内存中
type cache struct {
	m     map[string]*item // key=md5(req_path), val = expire timestamp
	mutex *sync.RWMutex
}

func (c *cache) Store(key string, content []byte, ttl time.Duration) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	data := &item{
		content: content,
		expire:  time.Now().Add(ttl).Unix(),
	}
	c.m[key] = data
	return nil
}

func (c *cache) Get(key string) (content []byte, hitted bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, hitted := c.m[key]
	if !hitted {
		return nil, false
	}
	if item.expire > time.Now().Unix() {
		return item.content, hitted
	}
	return nil, false
}

func (c *cache) clean() error {
	newm := make(map[string]*item)
	c.mutex.RLock()
	for key, item := range c.m {
		// has not expired
		if item.expire > time.Now().Unix() {
			newm[key] = item
		}
	}
	c.mutex.RUnlock()

	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.m = newm
	return nil
}

func (c *cache) Run() {
	t := time.NewTicker(time.Second * 10)
	for {
		<-t.C
		if err := c.clean(); err != nil {
			logrus.WithField("method", "Cache.run").Error(err)
		} else {
			logrus.WithField("method", "Cache.run").Info("success")
		}
	}
}
