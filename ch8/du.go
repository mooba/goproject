// author pengchengbai@shopee.com
// date 2021/3/21

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	start := time.Now() // 获取当前时间

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	du3(roots, fileSizes)

	fmt.Println("该函数执行完成耗时：", time.Since(start))
}

// 也不会快多少
func du3(roots []string, fileSizes chan int64) {
	sema := make(chan struct{}, 20)
	wg := sync.WaitGroup{}

	for _, root := range roots {
		wg.Add(1)
		sema <- struct{}{}
		go func(root string) {
			walkDir(root, fileSizes)

			<-sema
			wg.Done()
		}(root)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick = time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

// 在Mac pro上运行测试显示，这种方式并不比du1的串行方式快，这是因为创建了太多的goroutine
func du2(roots []string, fileSizes chan int64) {

	wg := sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go func(root string) {
			walkDir(root, fileSizes)
			wg.Done()
		}(root)
	}

	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick = time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func du1(roots []string, fileSizes chan int64) {
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	// Print the results periodically.
	var tick = time.Tick(500 * time.Millisecond)
	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop // fileSizes was closed
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes) // final totals
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.5f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
