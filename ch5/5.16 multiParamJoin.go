package main

import (
	"fmt"
)

// 练习5.16：编写多参数版本的strings.Join

func join(sep string, vals ...string) string {
	s := ""
	for _, val := range vals {
		s += val + sep
	}
	return s[:len(s) - len(sep)]
}

func main() {
	fmt.Println(join(""))
	fmt.Println(join("hello", "world"))
	fmt.Println(join("hello", "world", "!"))
	fmt.Println(join(",", "hello", "world", "!"))
}


// Join concatenates the elements of its first argument to create a single string. The separator
// string sep is placed between elements in the resulting string.
//func Join(elems []string, sep string) string {
//	switch len(elems) {
//	case 0:
//		return ""
//	case 1:
//		return elems[0]
//	}
//	n := len(sep) * (len(elems) - 1)
//	for i := 0; i < len(elems); i++ {
//		n += len(elems[i])
//	}
//
//	var b Builder
//	b.Grow(n)
//	b.WriteString(elems[0])
//	for _, s := range elems[1:] {
//		b.WriteString(sep)
//		b.WriteString(s)
//	}
//	return b.String()
//}
