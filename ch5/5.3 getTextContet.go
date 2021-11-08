package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visit3(links []string, n *html.Node) []string {
	if n.Type == html.TextNode && n.Data != "script" && n.Data != "style" {
		links = append(links, n.Data)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit3(links, c)
	}
	return links
}


func main() {
	doc, err := html.Parse(os.Stdin)
	//doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit3(nil, doc) {
		fmt.Println(link)
	}
}