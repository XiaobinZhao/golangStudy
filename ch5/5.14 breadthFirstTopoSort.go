package main

import (
	"fmt"
	"sort"
)


var prereqs3 = map[string][]string{
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

var seen = make(map[string]bool)

func preFunc(item string) bool{
	if !seen[item] {
		seen[item] = true
		return true
	}
	return false
}


func forEachGraph(items []string, order []string, pre, post func(item string)bool) []string {
	for _, item := range items {
		if pre(item) {
			order= forEachGraph(prereqs3[item], order, pre, post)
			order = append(order, item)
		}
	}
	return order
}


func breadthFirst2(f func(item []string, order []string, pre, post func(item string)bool) []string, m map[string][]string) {
	var order []string
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	order = f(keys, order, preFunc, nil)
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}


func main() {
	breadthFirst2(forEachGraph, prereqs3)
}
