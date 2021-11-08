package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

func visit4(tagRefMap map[string][]string, n *html.Node)  map[string][]string {
	if n.Type == html.ElementNode &&(n.Data == "a" || n.Data == "link" || n.Data == "script" || n.Data == "img"){
		for _, a := range n.Attr {
			if a.Key == "href" {
				tagRefMap[n.Data] = append(tagRefMap[n.Data], a.Val)
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		tagRefMap = visit4(tagRefMap, c)
	}
	return tagRefMap
}


func main() {
	doc, err := html.Parse(os.Stdin)
	//doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	tagRefMap := make(map[string][]string)
	for key, value := range visit4(tagRefMap, doc) {
		fmt.Println("key: ", key)
		fmt.Println("value: ", value)
	}
}