// author pengchengbai@shopee.com
// date 2021/2/1

package main

import (
	"fmt"
)

func main() {
	fmt.Println(Manual.String())
}


func unpackSlice() {
	a := []int{1,3,4}
	b := []int{4,5,6}

	a = append(a, b...)
	fmt.Println(a)
}



type BoardingType int

const (
	UnKnown BoardingType = iota
	Manual
	Survey
	System
)

func (b BoardingType) String() string {
	return [...]string{"UnKnown","Manual","Survey","System"}[b]
}


