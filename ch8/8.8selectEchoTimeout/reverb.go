
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

//!+
func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	inputChan := make(chan string)
	go func() {
		ticker := time.NewTicker(1 * time.Second)

		for countdown := 10; countdown > 0; countdown-- {
			fmt.Printf("countdown: %d \n", countdown)
			select {
			case  <-ticker.C:
				// do nothing
			case in := <-inputChan:
				fmt.Println("input something, countdown reset!")
				countdown = 10
				go echo(c, in, 1*time.Second)
			}
		}
		ticker.Stop()
		fmt.Println("countdown over, close conn!")
		c.Close()
	}()


	for input.Scan() {
		inputChan <- input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()



}

//!-
// 练习 8.8： 使用select来改造8.3节中的echo服务器，为其增加超时，
// 这样服务器可以在客户端10秒中没有任何喊话时自动断开连接。
// 分析：
// 服务器启动之后，等待客户端连接。客户端连接输入之后，触发handle。handle执行时，先启动一个goroutine进行倒计时，
// 然后一直等待输入(for死循环)，有输入时，把输入送入chan,在goroutine从chan取数据并重置倒计时。如此循环。直到没有输入，倒计时结束，
// 关闭客户端的连接。
func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
