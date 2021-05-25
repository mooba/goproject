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

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()


	ticker := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0 ; countdown-- {
		select {
		case <-abort:
			fmt.Println("Launch abort!")
			//close(ticker) ticker is a receive-only channel, and cannot be closed, will cause goroutine leak
			return
		case <-ticker:
			fmt.Println(countdown)
			// do nothing
		}
	}

	launch2()
}


func launch2() {
	fmt.Println("launch")
}

