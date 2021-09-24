package main

import (
	"fmt"
	"strings"
	"unicode"
)

func myNonempty1(s string) string{
	stemp := []byte(s)
	var spaceCount int

	for i, _ := range stemp {
		if unicode.IsSpace(int32(stemp[i])) {
			spaceCount ++
		} else {
			if spaceCount > 1 {
				var empty  []byte
				for j:=0;j<spaceCount-1;j++ {
					empty = append(empty, 32) // 32 是空格的int32 格式
				}
				stemp = append(empty, append(stemp[:i+1-spaceCount], stemp[i:]...)...)
				spaceCount = 0
			}
			continue
		}

	}
	return strings.TrimSpace(string(stemp))
}


func main() {
	input := "xx  xx   xx xx"
	out := myNonempty1(input)
	fmt.Println(out)
}