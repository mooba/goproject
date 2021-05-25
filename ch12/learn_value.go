// author pengchengbai@shopee.com
// date 2021/2/3

package main

import (
	"fmt"
	"reflect"
)

func main() {
	//testValue()
	//testStructValue()
	//testStructValue2()

	testNewByReflect()
}

func testValue() {
	a := 100

	va, vp := reflect.ValueOf(a), reflect.ValueOf(&a).Elem()

	fmt.Println(va, va.CanAddr(), va.CanSet())
	fmt.Println(vp, vp.CanAddr(), vp.CanSet())
}

func testStructValue() {
	var foo interface{}
	foo = mockItem1{
		Name: "foo",
		Priority1: 1,
		Priority2: 2,
	}
	fmt.Println(reflect.TypeOf(foo).Name(), reflect.TypeOf(foo).Kind())
	va, vp := reflect.ValueOf(foo), reflect.ValueOf(&foo).Elem()
	fmt.Println(va, vp)

	// 非指针得到的value去set值是不成功的
	p1 := va.FieldByName("Priority1")
	p2 := va.FieldByName("Priority2")
	fmt.Println(p1, p2)
	if p1.CanSet() {
		p1.SetInt(4)
	}
	fmt.Println(p1)
	fmt.Println(foo)
}

func testStructValue2()  {
	var foo interface{}
	foo = &mockItem1{
		Name: "foo",
		Priority1: 1,
		Priority2: 2,
	}
	fmt.Println(reflect.TypeOf(foo).Name(), reflect.TypeOf(foo).Kind())
	vp := reflect.ValueOf(foo).Elem()
	fmt.Println(vp, vp.Type(), vp.Kind())

	// 指针得到的value去set是可以成功的
	// refer:https://stackoverflow.com/questions/63421976/panic-reflect-call-of-reflect-value-fieldbyname-on-interface-value
	pp1 := vp.FieldByName("Priority1")
	pp2 := vp.FieldByName("Priority2")
	fmt.Println(pp1, pp2)
	if pp1.CanSet() {
		pp1.SetInt(4)
	}
	if pp2.CanSet() {
		pp2.SetInt(5)
	}
	
	fmt.Println(pp1)
	fmt.Println(foo)
}

type mockItem1 struct {
	Name      string
	Priority1 int
	Priority2 int
}

func testNewByReflect() {
	t := reflect.TypeOf(mockItem1{})

	p := reflect.New(t)
	pf := p.Elem().FieldByName("Priority1")
	fmt.Println(pf.Int())
	pfs := p.Elem().FieldByName("Name")
	fmt.Println(pfs.Type())

	if pf.IsValid() && pf.CanSet() {
		pf.SetInt(45)
	}

	p1 := reflect.New(t)
	p1f := p1.Elem().FieldByName("Priority1")

	fmt.Println(pf.Int() > p1f.Int())

	fmt.Println(p.Kind())
	fmt.Println(p.Type())
	fmt.Println(p.Interface())
}





