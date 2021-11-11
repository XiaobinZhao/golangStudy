package main

import "fmt"

// 练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
// 不包含return语句但能返回一个非零值的函数 其实就是使用 命名返回值的写法，也即：bare return

func panicAndRecover() (returnVal string) {
	defer func() {
		if p := recover(); p!= nil {
			returnVal = p.(string)
		}
	}()
	panic("i am a panic")
}

func main() {
	returnVal := panicAndRecover()
	fmt.Println(returnVal)
}
