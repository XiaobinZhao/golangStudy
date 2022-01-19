package memo

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)


//!+httpRequestBody
func HttpGetBodyWithCancel(url string, done chan struct{}) (interface{}, error) {
	ctx, cancel := context.WithCancel(context.Background())
	req, _ := http.NewRequestWithContext(ctx,"GET", url, nil)

	go func() {
		for {
			select {
			case <-done:
				fmt.Printf("%s canceled. \n", url)
				cancel()
			}
		}
	}()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)

}
//!-httpRequestBody

//!+httpRequestBody
func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"http://books.studygolang.com",
			"https://github.com/",
			"https://cn.bing.com",
			"http://books.studygolang.com",
			"https://github.com/",
			"https://cn.bing.com",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, done chan struct{}) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M, done chan struct{}) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, done)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M, done chan struct{}) {
	//!+conc
	responses := make(chan string, 3)
	for url := range incomingURLs() {
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url, done)
			if err != nil {
				log.Print(err)
				return
			}
			responses <- url
			fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	select {
	case res := <-responses:  // 跑的最快的获取，其他取消
		close(done)
		fmt.Println(res)
	}
	//!-conc
}
