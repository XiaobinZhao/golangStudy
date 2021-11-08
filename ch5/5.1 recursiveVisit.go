package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visit1(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	if n.FirstChild != nil {
		links = visit1(links, n.FirstChild)
	}
	if n.NextSibling != nil {
		links = visit1(links, n.NextSibling)
	}

	//for c := n.FirstChild; c != nil; c = c.NextSibling {
	//	links = visit1(links, c)
	//}
	return links
}


func main() {
	doc, err := html.Parse(os.Stdin)
	//doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit1(nil, doc) {
		fmt.Println(link)
	}
}