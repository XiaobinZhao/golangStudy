package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster
type client chan<- string // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
	inputChan = make(chan string)
	clientNum = 0
	clientNameMap = make(map[string]string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

//!-broadcaster

func disconnectIdleClient2(conn net.Conn) { // 另一种倒计时的写法
	for  {
		select {
		case <-time.After(10 * time.Second):
			conn.Close()
		case  str := <-inputChan:
			messages <- str
		}
	}
}


//!+handleConn
func handleConn(conn net.Conn, ch chan string) {
	go clientWriter(conn, ch)

	ip := conn.RemoteAddr().String()
	clientNum++
	name := "client" + fmt.Sprint(clientNum)
	if clientNameMap[ip] == "" {
		clientNameMap[ip] = name
	}
	ch <- "You are " + clientNameMap[ip]
	messages <- clientNameMap[ip] + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		inputChan <- clientNameMap[ip] + ": " + input.Text()
		//messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	messages <- name + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

//!-handleConn

// 练习 8.15： 如果一个客户端没有及时地读取数据可能会导致所有的客户端被阻塞。
// 修改broadcaster来跳过一条消息，而不是等待这个客户端一直到其准备好读写。
// 或者为每一个客户端的消息发送channel建立缓冲区，这样大部分的消息便不会被丢掉；
// broadcaster应该用一个非阻塞的send向这个channel中发消息。
func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		ch := make(chan string, 10)  // 可以缓存10条消息
		go handleConn(conn, ch)
		go disconnectIdleClient2(conn)
	}
}

//!-main
