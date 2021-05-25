package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)


// refer to https://www.liwenzhou.com/posts/Go/go_context/
func main() {
	fmt.Println("hello")
	//cancelBasedOnContext()

	//contextWithValue()

	contextWithTimeout()

	//TimeoutOpWithResponse(1*time.Second, TestFunc, 1*time.Second)
}

func contextWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	result, err := handle(ctx)
	if err != nil {
		fmt.Printf("error: %s", err)
	} else {
		fmt.Printf("result:%v", result)
	}
	fmt.Println("父协程运行中")
	time.Sleep(10*time.Second)
	fmt.Println("父协程退出")
}

func TestFunc(sleepTime time.Duration) (interface{}, error) {
	time.Sleep(sleepTime)
	return fmt.Sprintf("I sleeped %s", sleepTime.String()), nil
}

type MyResult struct {
	Result int
}

type Response struct {
	MyResult *MyResult
	Err error
}

func handle(ctx context.Context) (result *MyResult, err error)  {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
		return nil, fmt.Errorf("handle %s", ctx.Err())
	case res := <-handleHelper(ctx):
		return res.MyResult, res.Err
	}
}

func handleHelper(ctx context.Context) <- chan *Response{
	chResp := make(chan *Response, 1)
	go func() {
		result, err := actualHandle()
		if err != nil {
			return
		} else {
			chResp <- &Response{MyResult: result, Err: nil}
		}
		fmt.Println("子协程退出")
	}()
	return chResp
}

func actualHandle() (*MyResult, error) {
	i := 0
	for {
		time.Sleep(500*time.Millisecond)
		fmt.Println("子协程运行中")
		if i == 10 {
			break
		}
	}
	return &MyResult{1}, nil
}


func cancelBasedOnChannel() {
	var exitChan = make(chan struct{})
	wg.Add(1)
	go worker(exitChan)
	time.Sleep(time.Second * 3)
	exitChan <- struct{}{}
	close(exitChan)
	wg.Wait()
	fmt.Println("over")
}


func cancelBasedOnContext()  {
	ctx, cancelFunc := context.WithCancel(context.Background())
	wg.Add(2)
	go worker1(ctx)
	go worker2()
	time.Sleep(time.Second * 3)
	cancelFunc()
	wg.Wait()
	fmt.Println("over")
}



var wg sync.WaitGroup

func worker(exitChan chan struct{}) {
LOOP:
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <-exitChan:
			break LOOP
		default:
		}
	}

	// 如何接收外部命令实现退出
	wg.Done()
}

func worker1(ctx context.Context) {
LOOP:
	for true {
		fmt.Println("worker")
		time.Sleep(time.Second)
		select {
		case <- ctx.Done(): // 等待上级通知
			break LOOP
		default:
		}
	}

	wg.Done()
}

func worker2() {
	time.Sleep(time.Second * 4)
	fmt.Println("worker2 sleep 4 seconds")

	wg.Done()
}




func contextWithValue() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "232244")

	wg.Add(1)
	go worker3(ctx)
	time.Sleep(time.Second * 5)
	cancelFunc()
	wg.Wait()
	fmt.Println("over")
}



type  TraceCode string

func worker3(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	// 在子goroutine中获取trace code
	traceCode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", traceCode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

