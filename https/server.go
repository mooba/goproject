// Copyright 2021 Shopee, Inc.
// author pengchengbai@shopee.com
// date 2021/1/31
// reference https://tonybai.com/2015/04/30/go-and-https/

package main

import (
	"fmt"
	http "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi, This is an example of https service in golang!")
}

func main() {
	fmt.Println("hello world")
	http.HandleFunc("/", handler)

	_ = http.ListenAndServeTLS(":8081", "./server.crt", "./server.key", nil)
	//_ = http.ListenAndServe(":8080", nil)
}
