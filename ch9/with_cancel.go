// author pengchengbai@shopee.com
// date 2021/11/12

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func(cancelFunc context.CancelFunc) {
		time.Sleep(3*time.Second)
		fmt.Println("cannel operation")
		cancel()
	}(cancel)


	go func(ctx2 context.Context) {
		select {
		case <-ctx.Done():
			fmt.Println("cancel processing")
		default:
			for i := 0; i < 10; i++ {
				time.Sleep(1*time.Second)
				fmt.Println(i)
			}
			fmt.Println("woken up")
		}
	}(ctx)

	select {}

}
