package main

import "fmt"


func main() {
	dirs := []int{1,2,3}
	var rmdirs []func()
	for i := 0; i < len(dirs); i++ {
		i:=i  // declares inner i, initialized to outer i
		rmdirs = append(rmdirs, func() {
			fmt.Println(i)
		})
	}
	for _,f := range rmdirs {
		f()
	}

}
