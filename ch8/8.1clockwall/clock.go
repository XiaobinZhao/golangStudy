package main

import (
	"flag"
	"io"
	"log"
	"net"
	"time"
)

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		newZone := time.FixedZone("UTC" + string(*timezone), *timezone*60*60)  // 根据参数设置时区
		_, err := io.WriteString(c, time.Now().In(newZone).Format("15:04:05\n"))  // 返回相应时区的时间
		if err != nil {
			return // e.g., client disconnected
		}
		time.Sleep(1 * time.Second)
	}
}
// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，
// 从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
// 如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
// 这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
//
// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

var port = flag.String("port", "8000", "服务端口")
var timezone = flag.Int("timezone", +8, "时区")

func main() {
	flag.Parse()
	listener, err := net.Listen("tcp", "localhost:" + *port)
	if err != nil {
		log.Fatal(err)
	}
	//!+
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn) // handle connections concurrently
	}
	//!-
}
