// author pengchengbai@shopee.com
// date 2021/3/23

package main

import (
	"fmt"
)

func main() {
	learnDefer()
}

func learnDefer()  {
	defer fmt.Println("done1")
	defer fmt.Println("done2")

	printAllOperations(4, 0)
	fmt.Println("Exiting main without any issues")
}

func printAllOperations(x int, y int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovering from panic in printAllOperations error is: %v \n", r)
		}
	}()

	sum, subtract, multiply, divide := x+y, x-y, x*y, x/y
	fmt.Printf("sum=%v, subtract=%v, multiply=%v, divide=%v \n", sum, subtract, multiply, divide)
}