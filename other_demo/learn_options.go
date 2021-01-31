// Copyright 2021 Shopee, Inc.
// author pengchengbai@shopee.com
// date 2021/1/31

package main

import (
	"crypto/tls"
	"time"
)

type Server struct {
	Addr     string
	Port     int
	Protocol string
	Timeout  time.Duration
	MaxConns int
	TLS      *tls.Config
}
