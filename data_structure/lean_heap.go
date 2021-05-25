// author pengchengbai@shopee.com
// date 2021/2/2

package main

import (
	"container/heap"
	"fmt"
)

//定义一个正方形的结构体
type Rectangle struct {
	width  int
	height int
}

// 计算面积
func (rec *Rectangle) Area() int {
	return rec.width * rec.width
}

// 定义一个堆结构体
type RectHeap []Rectangle

// 实现sort.Iterface
func (rech RectHeap) Len() int {
	return len(rech)
}

func (rech RectHeap) Swap(i, j int) {
	rech[i], rech[j] = rech[j], rech[i]
}

// 小根堆
func (rech RectHeap) Less(i, j int) bool {
	return rech[i].Area() < rech[j].Area()
}

// 实现heap.Interface接口定义的额外方法
func (rech *RectHeap) Push(h interface{}) {
	*rech = append(*rech, h.(Rectangle))
}
func (rech *RectHeap) Pop() (x interface{}) {
	n := len(*rech)
	x = (*rech)[n-1]      // 返回删除的元素
	*rech = (*rech)[:n-1] // [n:m]不包括下标为m的元素
	return x
}

func main() {

}

func testHeap() {
	hp := &RectHeap{}
	for i := 2; i < 6; i++ {
		*hp = append(*hp, Rectangle{i, i})
	}

	fmt.Println("原始slice: ", hp)

	// 堆操作
	heap.Init(hp)
	heap.Push(hp, Rectangle{100, 10})
	fmt.Println("push 之后 : ", hp)
	fmt.Println("top元素：", (*hp)[0])
	fmt.Println("删除并返回最后一个：", heap.Pop(hp)) // 最后 一个元素
	fmt.Println("最终slice: ", hp)
}
