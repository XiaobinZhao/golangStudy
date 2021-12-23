package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

// 练习 8.1： 修改clock2来支持传入参数作为端口号，然后写一个clockwall的程序，这个程序可以同时与多个clock服务器通信，
// 从多个服务器中读取时间，并且在一个表格中一次显示所有服务器传回的结果，类似于你在某些办公室里看到的时钟墙。
// 如果你有地理学上分布式的服务器可以用的话，让这些服务器跑在不同的机器上面；或者在同一台机器上跑多个不同的实例，
// 这些实例监听不同的端口，假装自己在不同的时区。像下面这样：
//
// $ TZ=US/Eastern    ./clock2 -port 8010 &
// $ TZ=Asia/Tokyo    ./clock2 -port 8020 &
// $ TZ=Europe/London ./clock2 -port 8030 &
// $ clockwall NewYork=localhost:8010 Tokyo=localhost:8020 London=localhost:8030

var  zoneMap = map[string]int{"NewYork": -5, "Tokyo": +9, "London": 0, "Beijing": +8}


func main() {
	for _, arg := range os.Args[1:] {
		zoneAddress := strings.Split(arg, "=")
		go getServerTime(zoneAddress[1], zoneAddress[0])
	}
	for {
		time.Sleep(1)
	}
}


func getServerTime(address string, zoneName string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	var sb strings.Builder
	buf := make([]byte, 1)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		char := string(buf[:n])
		if char == "\n" {  // 读到换行符，说明读取到一次时间输入结束
			timeStr := sb.String()
			fmt.Printf("%9.9s %2d %s %s \n", zoneName, zoneMap[zoneName], address, timeStr)
			sb.Reset() // 读取成功一次，清空builder
			continue
		}
		sb.Write(buf[:n])
	}

}