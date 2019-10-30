package main

import (
	"context"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

type location struct {
	path            string
	proxyPass       string        `yaml:"proxy_pass"`
	cacheExpiration time.Duration `yaml:"cache_expiration"`
	cache           *cache
}

func (l *location) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	l.Handle(rw, req)
}

func (l *location) Handle(rw http.ResponseWriter, req *http.Request) {
	// 先判断缓存中是否存在
	logrus.WithField("origin path", req.URL.Path).Info("location handle")
	content, hitted := l.cache.Get(req.URL.Path)
	if hitted {
		logrus.Info("hitted")
		rw.Write(content)
		rw.WriteHeader(http.StatusOK)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	proxyReq, err := http.NewRequest(req.Method, l.proxyPass, req.Body)
	if err != nil {
		return
	}
	proxyReq = proxyReq.WithContext(ctx)
	proxyReq.URL.RawQuery = req.URL.Query().Encode()
	logrus.WithField("proxy url", proxyReq.URL.String()).Info("location handle: rewrite query")

	if req.ContentLength == 0 {
		proxyReq.Body = nil
	}
	if req.Header == nil {
		proxyReq.Header = make(http.Header)
	}
	proxyReq.Close = false

	// TODO 删除或者处理一些header

	if clientIP, _, err := net.SplitHostPort(req.RemoteAddr); err == nil {
		if prior, ok := proxyReq.Header["X-Forwarded-For"]; ok {
			clientIP = strings.Join(prior, ", ") + ", " + clientIP
		}
		proxyReq.Header.Set("X-Forwarded-For", clientIP)
	}
	res, err := http.DefaultTransport.RoundTrip(proxyReq)
	if err != nil {
		return
	}
	// 缓存数据
	// TODO 直接全读不安全，无法知道content有多大
	content, err = ioutil.ReadAll(res.Body)
	if err != nil {
		l.cache.Store(req.URL.Path, content, l.cacheExpiration)
	}
	logrus.WithField("content", string(content)).Info("location handle")

	// TODO 移除和连接相关的header
	copyHeader(res.Header, rw.Header())
	rw.Write(content)
	rw.WriteHeader(res.StatusCode)
}

func copyHeader(src, dst http.Header) {
	for k, vs := range src {
		for _, v := range vs {
			dst.Add(k, v)
		}
	}
}

// 获取所请求的资源的path: $proxypass+"/"+ req.path
func (l *location) getPath(req *http.Request) string {
	return ""
}
