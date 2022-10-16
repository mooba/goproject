// author pengchengbai@shopee.com
// date 2021/3/21

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type result struct {
	value interface{}
	err error
}

type Func func(key string) (interface{}, error)

type Memo struct {
	f Func
	cache map[string]result
}


func NewMemo(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]result),
	}
}



func (m *Memo) Get(key string) (interface{}, error){
	ret, ok := m.cache[key]
	if !ok {
		ret.value, ret.err = m.f(key)
		m.cache[key] = ret
	}
	return ret.value, ret.err
}

func main() {
	runWithMemo()
}

func runWithMemo() {
	m := NewMemo(httpGetBody)
	for _, url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

var incomingURLs = func() []string {
	return nil
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
