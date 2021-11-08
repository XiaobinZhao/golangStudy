package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
)

func main() {
	outline("http://books.studygolang.com/")
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	var depth3 int
	var startElement3, endElement3 func (n *html.Node)
	startElement3 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			fmt.Printf("%*s<%s>\n", depth3*2, "", n.Data)
			depth3++
		}
	}
	endElement3 = func(n *html.Node) {
		if n.Type == html.ElementNode {
			depth3--
			fmt.Printf("%*s</%s>\n", depth3*2, "", n.Data)
		}
	}
	//!+call
	forEachNode3(doc, startElement3, endElement3)
	//!-call

	return nil
}

func forEachNode3(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode3(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

//!-forEachNode

//!+startend


//func startElement3(n *html.Node) {
//	if n.Type == html.ElementNode {
//		fmt.Printf("%*s<%s>\n", depth3*2, "", n.Data)
//		depth3++
//	}
//}
//
//func endElement3(n *html.Node) {
//	if n.Type == html.ElementNode {
//		depth3--
//		fmt.Printf("%*s</%s>\n", depth3*2, "", n.Data)
//	}
//}

//!-startend
