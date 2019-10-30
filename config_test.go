package main

import (
	"errors"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	srv := new(Srv)
	err := parseConfig(srv, "./etc_sample/conf.yaml")
	if err != nil {
		t.Error(err)
	}

	if srv.Port != 1088 {
		t.Error(errors.New("port mismatch"))
	}
	if srv.Parallel != 10000 {
		t.Error(errors.New("parallel mismatch"))
	}
	if srv.Root != "./" {
		t.Error(errors.New("root mismatch"))
	}
	if srv.CachePath != "/var/goproxy/cache" {
		t.Error(errors.New("cache path mismatch"))
	}
	if srv.CacheExpiration != time.Minute*3 {
		t.Error(errors.New("cache expiration mismatch"))
	}
	if len(srv.Locations) != 1 {
		t.Error(errors.New("location mismatch"))
	}

	err = srv.Validate()
	if err != nil {
		t.Error(err)
	}
}
