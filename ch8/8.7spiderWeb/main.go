package main

import (
	"fmt"
	"log"
	"studygolang/links"
)

const domainUrl = "https://qumaicai.top"

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url, domainUrl)
	if err != nil {
		log.Print(err)
	}
	return list
}

// 练习 8.7： 完成一个并发程序来创建一个线上网站的本地镜像，把该站点的所有可达的页面都抓取到本地硬盘。
// 为了省事，我们这里可以只取出现在该域下的所有页面（比如golang.org开头，译注：外链的应该就不算了。）
// 当然了，出现在页面里的链接你也需要进行一些处理，使其能够在你的镜像站点上进行跳转，而不是指向原始的链接.
func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	// Add command-line arguments to worklist.
	initLinks := []string{domainUrl}
	go func() { worklist <- initLinks }()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 10; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				unseenLinks <- link
			}
		}
	}
}

//!-
