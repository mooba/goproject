// author pengchengbai@shopee.com
// date 2021/3/20

package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")
	naturals := make(chan<-int)
	naturals <- 10
	fmt.Println(len(naturals))
	close(naturals)

	//
	////m := make(map[int]int)
	//done := make(chan struct{})
	//go func() {
	//	time.Sleep(1*time.Second)
	//	fmt.Println("Hello Channel")
	//	done <- struct{}{}
	//}()
	//fmt.Println("Hello World")
	//<-done // wait for background goroutine to finish
}

func fibo(x int) int {
	if x < 2 {
		return x
	}
	return fibo(x-1) + fibo(x-2)
}
