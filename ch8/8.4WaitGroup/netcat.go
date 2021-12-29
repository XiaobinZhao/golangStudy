package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})
	go func() {
		_,err := io.Copy(os.Stdout, conn)
		log.Println("done")
		log.Println(err)
		done <- struct{}{} // signal the main goroutine
	}()
	mustCopy(conn, os.Stdin)
	//cw := conn.(*net.TCPConn)
	//cw.CloseWrite()
	conn.Close()
	<-done // wait for background goroutine to finish
}

//!-

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
