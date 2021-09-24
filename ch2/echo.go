package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit newline end")
var step = flag.String("s", "#", "separator")

func main() {
	flag.Parse() // 初始化命令行参数到声明变量
	fmt.Printf(strings.Join(flag.Args(), *step))
	if !*n {
		fmt.Println()
	}
}
