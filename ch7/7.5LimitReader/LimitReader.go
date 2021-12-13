package main

import (
	"fmt"
	"io"
	"strings"
)


// 练习 7.5： io包里面的LimitReader函数接收一个io.Reader接口类型的r和字节数n，
// 并且返回另一个从r中读取字节但是当读完n个字节后就表示读到文件结束的Reader。实现这个LimitReader函数：
//
// func LimitReader(r io.Reader, n int64) io.Reader

type MyHTML struct {
	r io.Reader
	limit int64
}

func (m *MyHTML) Read(p []byte) (n int, err error) {
	if m.limit <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > m.limit {
		p = p[:m.limit]
	}
	n ,err = m.r.Read(p)
	m.limit -= int64(n)
	return
}

func LimitReader(r io.Reader, n int64) io.Reader {
	return &MyHTML{r, n}
}

func main() {
	r := strings.NewReader("hello world")
	lr := LimitReader(r, 5)
	every := make([]byte, 3)
	n, _ := lr.Read(every)
	fmt.Println(n, every, string(every))

}