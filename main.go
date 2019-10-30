package main

import (
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
)

type Srv struct {
	// server config
	Port int
	// 最大并发数，默认5000
	Parallel int

	Root            string
	CachePath       string        `yaml:"cache_path"`
	CacheExpiration time.Duration `yaml:"cache_expiration"`

	Locations []*location `yaml:"locations"`

	limiter chan struct{}
}

func main() {
	srv := new(Srv)
	gpanic(parseConfig(srv, ""))

	cache, err := NewCache()
	gpanic(err)
	for _, l := range srv.Locations {
		l.Cache = cache
		logrus.WithField("path", l.Path).Info("init cache")
	}

	gpanic(srv.Validate())
	go cache.Run()

	serving := ":" + strconv.Itoa(srv.Port)
	go http.ListenAndServe(serving, srv)
	logrus.Info("serving at" + serving)
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
	logrus.Info("stopping ...")
	time.Sleep(time.Second * 3)
}
