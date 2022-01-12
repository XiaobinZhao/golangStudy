package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

//!+
// 练习 8.9： 编写一个du工具，每隔一段时间将root目录下的目录大小计算并显示出来。
// 分析：使用time.Tick 5s秒钟输出一次images目录的大小。程序执行时立刻运行一次，然后等待tick到来。
// 使用WaitGroup等待每一次统计结束，输出结果.
var nfiles, nbytes int64

func main() {

	root := "e:\\images"
	fileSizes := make(chan int64)
	tick := time.Tick(10 * time.Second)  // 5s打印一次指定目录的total size
	var n sync.WaitGroup
	n.Add(1)
	go waitToFinish(&n, fileSizes)
	go walkDir(root, &n, fileSizes)
	fmt.Printf("after 5s du will start ..., %s \n", time.Now())


	for  {
		select {  // select 会一直等待直到其case任何一个执行。执行之后select结束。所以在select外需要一直死循环。
		case <-tick:
			fmt.Printf("it is time to du..., %s \n", time.Now())
			nbytes = 0
			nfiles = 0
			fileSizes = make(chan int64)
			n.Add(1)
			go waitToFinish(&n, fileSizes)
			go walkDir(root, &n, fileSizes)
		case size, ok := <-fileSizes:
			if !ok {
				break
			} else {
				//fmt.Printf("received size: %d.\n", size)
				nfiles++
				nbytes += size
			}
		}
	}
}

func waitToFinish(n *sync.WaitGroup, fileSizes chan int64) {
	n.Wait()
	close(fileSizes)
	fmt.Println("congratulation！ du finish!")
	printDiskUsage(nfiles, nbytes) // final totals
}
//!-

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

// walkDir recursively walks the file tree rooted at dir
// and sends the size of each found file on fileSizes.
//!+walkDir
func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

//!-walkDir

//!+sema
// sema is a counting semaphore for limiting concurrency in dirents.
var sema = make(chan struct{}, 20)

// dirents returns the entries of directory dir.
func dirents(dir string) []os.FileInfo {
	sema <- struct{}{}        // acquire token
	defer func() { <-sema }() // release token
	// ...
	//!-sema

	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
