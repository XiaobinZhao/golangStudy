package main

import (
	"fmt"
	"net/http"
	"sync"
)

var count = 0
var  mu sync.Mutex

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", countHandler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count ++
	mu.Unlock()
	fmt.Fprintf(w,"URL.path: %q", r.URL.Path)
}

func countHandler( w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "URL.path:%q, count: %q", r.URL.Path, count)
	mu.Unlock()
}