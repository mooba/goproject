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
	flag.Parse()

	fmt.Printf("Sleeping for %v...", *period)
	time.Sleep(*period)
	fmt.Println()
}
