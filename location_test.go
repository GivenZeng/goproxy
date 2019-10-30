package main

import (
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestLocation(t *testing.T) {
	l := &location{
		path:            "/a/b",
		proxyPass:       "https://common.givenzeng.cn/mmdb",
		cacheExpiration: time.Minute,
	}

	cache, err := NewCache()
	if err != nil {
		t.Error(err)
		return
	}
	l.cache = cache

	// req, err := http.NewRequest("GET", "http://localhost:18080?ip=123.123.123.123", nil)
	if err != nil {
		t.Error(err)
		return
	}

	go func() {
		http.ListenAndServe(":18080", l)
	}()

	res, err := http.Get("http://localhost:18080/a/b?ip=123.123.123.123")
	if err != nil {
		t.Error(err)
		return
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(body))
}
