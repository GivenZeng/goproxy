package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestCache(t *testing.T) {
	datas := make(map[string][]byte)
	for i := 0; i < 100; i++ {
		length := rand.Intn(100)
		content := strings.Repeat("a", length)
		datas[strconv.Itoa(i)] = []byte(content)
	}
	c, _ := NewCache()
	go c.Run()

	for k, content := range datas {
		c.Store(k, content, time.Second)
	}
	t.Log("finish store")

	for i := 0; i < 10; i++ {
		go func() {
			idx := rand.Intn(100)
			content, hitted := c.Get(strconv.Itoa(idx))
			if !hitted {
				t.Error("not found in cache")
			}
			if len(content) != len(datas[strconv.Itoa(idx)]) {
				t.Error("length mismatch")
			}
		}()
	}
}
