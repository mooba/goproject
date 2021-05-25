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

// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

func main() {
	start := time.Now() // 获取当前时间

	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)

	du(roots, fileSizes)

	fmt.Println("该函数执行完成耗时：", time.Since(start))
}


func du(roots []string, fileSizes chan int64) {
	//sema := make(chan struct{}, 20)
	wg := sync.WaitGroup{}
	for _, root := range roots {
		wg.Add(1)
		go walkDir1(root, &wg, fileSizes)
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
			printDiskUsage1(nfiles, nbytes)
		}
	}
	printDiskUsage1(nfiles, nbytes) // final totals
}

func printDiskUsage1(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.5f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
func walkDir1(dir string, wg *sync.WaitGroup, fileSizes chan<- int64) {
	defer wg.Done()

	for _, entry := range dirents1(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir1(subdir, wg, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents returns the entries of directory dir.
func dirents1(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1: %v\n", err)
		return nil
	}
	return entries
}
