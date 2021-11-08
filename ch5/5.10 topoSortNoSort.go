package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

// 练习5.10： 重写topoSort函数，用map代替切片并移除对key的排序代码。验证结果的正确性（结果不唯一）。

func main() {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	//sort.Strings(keys)
	visitAll(keys)
	return order
}

// 结果一
//1:      linear algebra
//2:      calculus
//3:      intro to programming
//4:      discrete math
//5:      data structures
//6:      formal languages
//7:      computer organization
//8:      operating systems
//9:      networks
//10:     programming languages
//11:     algorithms
//12:     compilers
//13:     databases
// 结果二
//1:      intro to programming
//2:      discrete math
//3:      data structures
//4:      computer organization
//5:      programming languages
//6:      linear algebra
//7:      calculus
//8:      operating systems
//9:      formal languages
//10:     networks
//11:     algorithms
//12:     compilers
//13:     databases

// 结果三
//1:      intro to programming
//2:      discrete math
//3:      data structures
//4:      formal languages
//5:      computer organization
//6:      compilers
//7:      operating systems
//8:      algorithms
//9:      databases
//10:     networks
//11:     programming languages
//12:     linear algebra
//13:     calculus