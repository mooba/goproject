// author pengchengbai@shopee.com
// date 2021/2/10

package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
)

// FormatAny format any value as a string
func FormatAny(i interface{}) string {
	return formatAtom(reflect.ValueOf(i))
}

func formatAtom(value reflect.Value) string {
	fmt.Printf("value kind : %s \n", value.Kind().String())
	switch value.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(value.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(value.Uint(), 10)
	case reflect.String:
		return strconv.Quote(value.String())
	case reflect.Bool:
		return strconv.FormatBool(value.Bool())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return value.Type().String() + " 0x" + strconv.FormatUint(uint64(value.Pointer()), 16)
	default:
		return value.Type().String() + " value"
	}
}

func main() {
	str := "hello world"
	fmt.Println(FormatAny(str))

	ptr := &str
	fmt.Println(FormatAny(ptr))

	user := User2{
		nickName: "foo",
		age:      10,
	}
	fmt.Println(FormatAny(user))

	var p *User2
	fmt.Println(reflect.ValueOf(p).Kind() == reflect.Ptr) // should be true
	fmt.Println(reflect.ValueOf(p).IsNil()) // should be true

	var i interface{} = 3
	fmt.Println(reflect.ValueOf(i).Kind())

	toJsonStr()
}

type inter interface {
	hello()
}

type User2 struct {
	nickName string
	age  int
}

func (u User2) hello() {
	fmt.Println("hello")
}

func toJsonStr() {
	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))
}
