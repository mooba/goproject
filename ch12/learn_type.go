// author pengchengbai@shopee.com
// date 2021/2/3

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//testIntType()
	testStringType()
	//testStructType()
	//outPut()

	//testTraversalFieldOfStruct()
}

func testIntType() {
	var a = 45
	typ := reflect.TypeOf(a)

	// output int
	fmt.Println(typ)
	fmt.Printf("type name %s \n", typ.Name())
	fmt.Printf("type kind %s \n", typ.Kind())
}

func testStringType() {
	var a = "hello"
	typ := reflect.TypeOf(a)

	fmt.Println(typ)
	fmt.Printf("type name %s \n", typ.Name())
	fmt.Printf("type kind %s \n", typ.Kind())
}


func testStructType() {
	var foo = mockItem{
		Name: "foo",
	}
	typ := reflect.TypeOf(foo)

	// output main.mockItem
	fmt.Println(typ)
	fmt.Printf("type name %s \n", typ.Name())
	fmt.Printf("type kind %s \n", typ.Kind())
}

func outPut() {
	a := reflect.ArrayOf(10, reflect.TypeOf(byte(0)))
	b := reflect.MapOf(reflect.TypeOf(""), reflect.TypeOf(0))

	fmt.Println(a, b)
}

func testTraversalFieldOfStruct() {
	var m mockItem = mockItem{
		Name:      "foo",
		Priority1: 2,
		Priority2: 1,
	}

	t := reflect.TypeOf(&m)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type, f.Offset, f.Index)

		if f.Anonymous {
			for j := 0; j < f.Type.NumField(); j++ {
				af := f.Type.Field(j)
				fmt.Println(" ", af.Name, af.Type)
			}
		}



	}
}

type mockItem struct {
	Name      string
	Priority1 int
	Priority2 int

	// embedded field
	user
}

type user struct {
	nickName string
	age  int
}
