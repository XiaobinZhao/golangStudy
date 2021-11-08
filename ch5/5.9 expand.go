package main

import (
	"fmt"
	"strings"
)

// 练习 5.9： 编写函数expand，将s中的"foo"替换为f("foo")的返回值
// func expand(s string, f func(string) string) string

func fun (s string) string {
	return "blablabla"
}

func expand (s string) string {
	fmt.Printf("befor replase, s is :%s \n", s)
	s = strings.ReplaceAll(s, "foo", fun("foo"))
	fmt.Printf("after replase, s is :%s \n", s)
	return s
}

func main() {
	expand("abcfoo")
}
