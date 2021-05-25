// author pengchengbai@shopee.com
// date 2021/2/10

package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	//fmt.Printf("%s", "hello")

	a := int32(45)
	t := reflect.TypeOf(a)
	fmt.Println(t)
	fmt.Println(t.Kind())

	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w))

	aa := User1{
		nickName: "hello",
		age:      0,
	}

	fmt.Printf("%T \n", aa)
	fmt.Println(reflect.TypeOf(aa))


	v := reflect.ValueOf(a)
	fmt.Println(v.String())

	fmt.Println(v.Type() == t)

	i := v.Interface()
	s := i.(string)
	fmt.Println(s)
}

type User1 struct {
	nickName string
	age  int
}