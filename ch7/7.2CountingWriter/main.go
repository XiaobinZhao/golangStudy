package main

import (
	"fmt"
	"io"
	"os"
)


type MyWriter struct {
	w io.Writer
	c int64
}


func (c *MyWriter) Write(p []byte) (int, error) {
	c.c = int64(len(p))
	n, err := c.w.Write(p)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	myWriter := &MyWriter{w: w}
	return myWriter, &myWriter.c  // 返回值类型是io.Writer，他是一个interface，所以io.Writer的子类都可以返回
	                              // 但是这里有必须要返回myWriter，这个指针。这样后续的Write操作才会把字节数设置到c上
}

// 练习 7.2： 写一个带有如下函数签名的函数CountingWriter，传入一个io.Writer接口类型，
// 返回一个把原来的Writer封装在里面的新的Writer类型和一个表示新的写入字节数的int64类型指针。
//
// func CountingWriter(w io.Writer) (io.Writer, *int64)
func main() {
	rw, ilen := CountingWriter(os.Stdout)
	rw.Write([]byte("hello,world\n"))
	fmt.Println(*ilen)
}
