// author pengchengbai@shopee.com
// date 2021/3/21

package ch9

import "testing"



func Test(t *testing.T) {
	incomingURLs = func() []string {
		return []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		}
	}


	runWithMemo()

}
