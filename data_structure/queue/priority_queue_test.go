// author pengchengbai@shopee.com
// date 2021/2/2

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)


func TestPriorityPut(t *testing.T) {
	q := NewPriorityQueue(1, false)

	q.Put(mockItem(2))

	assert.Len(t, q.items, 1)
	assert.Equal(t, mockItem(2), q.items[0])

	q.Put(mockItem(1))

	if !assert.Len(t, q.items, 2) {
		return
	}

	assert.Equal(t, mockItem(1), q.items[0])
	assert.Equal(t, mockItem(2), q.items[1])
}

func TestPriorityPut_with_mockItem(t *testing.T) {
	q := NewPriorityQueue(1, false)
	cmpFun1 := func(item1, item2 mockItem2) int {
		return item1.Priority1 - item2.Priority1
	}
	i1 := mockItem2{
		Name:      "foo",
		Priority1: 2,
		Priority2: 1,
		CmpFunc:   nil,
	}


	i2 := mockItem2{
		Name:      "bar",
		Priority1: 1,
		Priority2: 2,
		CmpFunc:   nil,
	}

	i3 := mockItem2{
		Name:      "fee",
		Priority1: 3,
		Priority2: 2,
		CmpFunc:   nil,
	}

	i1.CmpFunc = cmpFun1
	i1.CmpFunc = cmpFun1
	i1.CmpFunc = cmpFun1

	q.Put(i1, i2, i3)

	assert.Equal(t, i2, q.items[0])
	assert.Equal(t, i1, q.items[1])
	assert.Equal(t, i3, q.items[3])
}


type mockItem2 struct {
	Name string
	Priority1 int
	Priority2 int

	CmpFunc func(item1, item2 mockItem2) int
}


func (m mockItem2) Compare(other Item) int {
	return m.CmpFunc(m, other.(mockItem2))
}


type mockItem int

func (mi mockItem) Compare(other Item) int {
	omi := other.(mockItem)
	if mi > omi {
		return 1
	} else if mi == omi {
		return 0
	}
	return -1
}