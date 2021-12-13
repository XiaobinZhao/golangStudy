package main

import (
	"bufio"
	"fmt"
	"strings"
)


type WorldCounter int

func (c *WorldCounter) Write(p []byte) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)
	//scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*c ++
	}
	if err := scanner.Err();err != nil {
		return int(*c), err
	}
	return int(*c), nil
}

// 练习 7.1： 使用来自ByteCounter的思路，实现一个针对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。
func main() {
	var c WorldCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Spicy jalapeno pastrami ut ham turducken.\n Lorem sed ullamco, leberkas sint short loin strip steak ut shoulder shankle porchetta venison prosciutto turducken swine.\n Deserunt kevin frankfurter tongue aliqua incididunt tri-tip shank nostrud.\n"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
}
