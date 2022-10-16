// author pengchengbai@shopee.com
// date 2021/11/5

package main

import (
	"fmt"
	"sync"
)

type Person struct {
	Name string
}

var personPool *sync.Pool

func initPool() {
	personPool = &sync.Pool{
		New: func() interface {} {
			fmt.Printf("Creating a new Person")
			return new(Person)
		},
	}
}

func main() {
	initPool()


	p := personPool.Get().(*Person)
	fmt.Println("首次从 pool 里获取：", p)

	p.Name = "first"
	fmt.Printf("设置 p.Name = %s\n", p.Name)

	personPool.Put(p)

	fmt.Println("Pool 里已有一个对象：&{first}，调用 Get: ", personPool.Get().(*Person))
	fmt.Println("Pool 没有对象了，调用 Get: ", personPool.Get().(*Person))
}
