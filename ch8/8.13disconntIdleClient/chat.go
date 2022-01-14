package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

//!+broadcaster

type clientObject struct {
	name string
	latest time.Time
	message string
}

type client chan<- clientObject // an outgoing message channel

var (
	entering = make(chan client)
	leaving  = make(chan client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- clientObject{message: msg}
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

func disconnectIdleClient(conn net.Conn) {

}


//!+handleConn
func handleConn(conn net.Conn) {
	clientChan := make(chan clientObject)
	//ch := make(chan string) // 用于输出当前client消息
	go clientWriter(conn, clientChan)

	who := conn.RemoteAddr().String()
	client := clientObject{
		name: who,
		latest: time.Now(),
		message: "You are " + who,
	}
	clientChan <- client
	messages <- who + " has arrived"

	entering <- clientChan  // enter事件需要使用clientObject

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- who + ": " + input.Text()
		client.latest = time.Now()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- clientChan
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan clientObject) {
	for client := range ch {
		fmt.Fprintln(conn, client.message) // NOTE: ignoring network errors
	}
}

//!-handleConn

// 练习 8.13： 使聊天服务器能够断开空闲的客户端连接，比如最近五分钟之后没有发送任何消息的那些客户端。
// 提示：可以在其它goroutine中调用conn.Close()来解除Read调用，就像input.Scanner()所做的那样。
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
		go handleConn(conn)
		go disconnectIdleClient(conn)
	}
}

//!-main
