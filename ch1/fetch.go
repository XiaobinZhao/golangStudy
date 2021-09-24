package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	url := os.Args[1]
	if !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Get http failed: %v", err)
		os.Exit(1)
	}
	_, err = io.Copy(os.Stdout, resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Printf("接受response 错误： %v", err)
		os.Exit(2)
	}
	fmt.Println(resp.Status)
}
