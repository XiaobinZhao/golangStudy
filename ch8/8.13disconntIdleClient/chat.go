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

func disconnectIdleClient2(conn net.Conn, ch chan string) { // 另一种倒计时的写法
	for  {
		select {
		case <-time.After(5 * time.Minute):
			conn.Close()
		case  str := <-inputChan:
			messages <- str
		}
	}
}

func disconnectIdleClient(conn net.Conn, ch chan string) {
	ticker := time.NewTicker(1 * time.Second)

	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("countdown: %d \n", countdown)
		select {
		case  <-ticker.C:
			// do nothing
		case in := <-inputChan:
			fmt.Println("input something, countdown reset!")
			countdown = 10
			messages <- in
		}
	}
	ticker.Stop()
	fmt.Println("countdown over, close conn!")
	conn.Close()
}


//!+handleConn
func handleConn(conn net.Conn, ch chan string) {
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch

	input := bufio.NewScanner(conn)

	for input.Scan() {
		inputChan <- who + ": " + input.Text()
		//messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()
	leaving <- ch
	messages <- who + " has left"
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
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
		ch := make(chan string)
		go handleConn(conn, ch)
		go disconnectIdleClient(conn, ch)
	}
}

//!-main
