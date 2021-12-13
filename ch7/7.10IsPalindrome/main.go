package main

import (
	"fmt"
	"sort"
	"strings"
)

type s struct {
	s []string
}

func (s s) Len() int { return len(s.s) }
func (s s) Less(i, j int) bool { return s.s[i] < s.s[j]}
func (s s) Swap(i, j int) { s.s[i], s.s[j] = s.s[j], s.s[i]}

func IsPalindrome(s sort.Interface) bool {
	var i,j = 0, s.Len() - 1

	for i<j {
		if !s.Less(i, j) && !s.Less(j, i) {  // 索引i和j上的元素相等
			i++
			j--
		} else {
			return false
		}
	}
	return true
}



// 练习 7.10： sort.Interface类型也可以适用在其它地方。编写一个IsPalindrome(s sort.Interface) bool函数表明序列s是否是回文序列，
// 换句话说反向排序不会改变这个序列。假设如果!s.Less(i, j) && !s.Less(j, i)则索引i和j上的元素相等。

func main() {
	fmt.Println(IsPalindrome(s{strings.Split("abba", "")}))
	fmt.Println(IsPalindrome(s{strings.Split("xabcbax", "")}))
	fmt.Println(IsPalindrome(s{strings.Split("aaaa", "")}))
	fmt.Println(IsPalindrome(s{strings.Split("abc", "")}))
	fmt.Println(IsPalindrome(s{strings.Split("aabbc", "")}))
}
