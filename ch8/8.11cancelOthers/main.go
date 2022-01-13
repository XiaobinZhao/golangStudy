package main

import (
	"context"
	"fmt"
	"net/http"
)
var done = make(chan struct{})

func mirroredQuery() {
	responses := make(chan string, 3)
	go func() { responses <- request("http://books.studygolang.com") }()
	go func() { responses <- request("https://github.com/") }()
	go func() { responses <- request("https://cn.bing.com") }()
	select {
	case res := <-responses:
		close(done)
		fmt.Println(res)
	}
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func request(hostname string) string {
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx,"GET", hostname, nil)
	if cancelled() {  // 执行的太快了么？ cancel已经不能生效了
		fmt.Printf("%s canceled.", hostname)
		cancel()
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Sprintf("%s request failed.", hostname)
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Sprintf("getting %s: %s", hostname, resp.Status)
	}
	defer resp.Body.Close()
	return fmt.Sprintf("%s success.", hostname)
}

// 练习 8.11： 紧接着8.4.4中的mirroredQuery流程，实现一个并发请求url的fetch的变种。
// 当第一个请求返回时，直接取消其它的请求。
func main() {
	mirroredQuery()
}
