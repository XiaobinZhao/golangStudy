package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var tagCount = map[string]int{}

func visit2(links []string, n *html.Node) {
	if n.Type == html.ElementNode {
		tagCount[n.Data]++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit2(links, c)
	}
}



func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit2(nil, doc)
	fmt.Printf("tagCount: %v", tagCount)
}