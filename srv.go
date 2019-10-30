package main

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

func (srv *Srv) Validate() error {
	if srv.Port == 0 {
		srv.Port = 18080
	}
	if srv.Parallel <= 0 {
		srv.Parallel = 5000
	}
	if srv.CachePath == "" {
		srv.CachePath = "/var/goproxy/cache"
	}
	srv.limiter = make(chan struct{}, srv.Parallel)
	return nil
}

func (srv *Srv) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	select {
	case srv.limiter <- struct{}{}:
		path := req.URL.Path
		logrus.WithField("path", path).Info(req.URL.String())
		// TODO use map
		for _, l := range srv.Locations {
			if l.Path == path {
				logrus.Info("location = " + l.Path)
				l.Handle(rw, req)
			}
		}
		<-srv.limiter
	default:
		logrus.Info("parallel limit")
		rw.WriteHeader(http.StatusServiceUnavailable)
	}
}
