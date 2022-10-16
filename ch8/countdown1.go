// author pengchengbai@shopee.com
// date 2021/3/21

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Commencing countdown")

	//tick := time.Tick(1*time.Second)
	//for countdown := 10; countdown > 0 ; countdown-- {
	//	fmt.Println(countdown)
	//	<-tick
	//}

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	select {
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	case <-time.After(10 * time.Second):
		// do nothing
	}

	launch()
}

func launch() {
	fmt.Println("launch")
}

func launch1() {
	fmt.Println("commencing countdown...")
	tick := time.Tick(1 * time.Second)

	for countDown := 10; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-tick
	}
	launch()
}
