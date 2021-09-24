package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	fmt.Printf("发起请求： %v \n", url)
	start := time.Now()
	resp , err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("发起http请求失败： %v", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("读取response失败： %v", err)
		return
	}
	dua := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", dua, nbytes, url)
}


func main() {
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
}
