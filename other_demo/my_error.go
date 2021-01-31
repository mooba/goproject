// Copyright 2021 Shopee, Inc.
// author pengchengbai
// date 2021/1/26

package main

import (
	"fmt"
	"strconv"
	"strings"
)
import "errors"

//自定义的出错结构
type myError struct {
	arg    int
	errMsg string
}

//实现Error接口
func (e *myError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.errMsg)
}

//两种出错
func error_test(arg int) (int, error) {
	if arg < 0 {
		return -1, errors.New("Bad Arguments - negtive!")
	} else if arg > 256 {
		return -1, &myError{arg, "Bad Arguments - too large!"}
	}
	return arg * arg, nil
}

//相关的测试
func main() {
	for _, i := range []int{-1, 4, 1000} {
		if r, e := error_test(i); e != nil {
			fmt.Println("failed:", e)
		} else {
			fmt.Println("success:", r)
		}
	}

	strbuilder()
}

func strbuilder() {
	var sb strings.Builder
	for i := 0; i < 10; i++ {
		sb.WriteString(strconv.Itoa(i))
	}
	fmt.Println(sb.String())
	sb.Reset()
	fmt.Println(sb.String())
	sb.WriteString("AB")
	sb.WriteString("CD")
	fmt.Println(sb.String())

	s := "ABV" + "CD"
	fmt.Println(s)
}


type FSSCategoryType int

const (
	FSS FSSCategoryType = 1
	CCB FSSCategoryType = 2
)

