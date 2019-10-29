package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

type Srv struct {
	// server config
	Port int
	// 最大并发数，默认5000
	Parallel int

	Root            string
	CachePath       string        `yaml:"cache_path"`
	CacheExpiration time.Duration `yaml:"cache_expiration"`

	Locations []*location

	limiter chan struct{}
}

func main() {
	srv := new(Srv)
	gpanic(parseConfig(srv, ""))

	cache, err := NewCache(srv.CachePath)
	gpanic(err)
	for _, l := range srv.Locations {
		l.cache = cache
	}
	go cache.Run()
	go http.ListenAndServe(":"+strconv.FormatInt(18080, 10), srv)
	waitKill()
}

func gpanic(err error) {
	if err != nil {
		panic(err)
	}
}

func waitKill() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	time.Sleep(time.Second * 30)
}
