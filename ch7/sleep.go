// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/24

package main

import (
	"flag"
	"fmt"
	"time"
)

var period = flag.Duration("period", 1*time.Second, "sleep period")


func main() {
	//flag.Parse()
	//
	//fmt.Printf("Sleeping for %v...", *period)
	//time.Sleep(*period)
	//fmt.Println()

	fibo(20)
	time.Sleep(3600 * time.Second)
}

func fibo(n int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		tmp := b
		b = a + b
		a = tmp
		fmt.Println(b)
	}

}