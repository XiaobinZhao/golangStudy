package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

func startElement6(n *html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if n.Type == html.ElementNode {
		for _, name := range names {
			if name == n.Data {
				nodes = append(nodes, n)
			}
		}
	}
	return nodes
}

func forEachNode6(n *html.Node, pre func(n *html.Node, names ...string) []*html.Node, names ...string) []*html.Node {
	var nodes []*html.Node
	if pre != nil {
		nodes = pre(n, names...)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		nodes = append(nodes, forEachNode6(c, pre, names...)...)
	}
	return nodes
}

func ElementsByTagName(doc *html.Node, names ...string) []*html.Node {
	nodes := forEachNode6(doc, startElement6, names...)
	return nodes
}

//练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：
//
//func ElementsByTagName(doc *html.Node, name...string) []*html.Node
//images := ElementsByTagName(doc, "img")
//headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

func main() {
	url := "http://books.studygolang.com/"
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("get url:%s failed: %s", url, err)
		os.Exit(1)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		os.Exit(1)
	}

	nodes := ElementsByTagName(doc, "meta")
	for _, n := range nodes {
		attrs := ""
		for _, a := range n.Attr {
			attrs = attrs + " " + a.Key + "=" + a.Val
		}
		fmt.Printf("<%s%s></%s> \n", n.Data, attrs, n.Data)
	}
	nodes = ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	fmt.Println("=======================")
	for _, n := range nodes {
		attrs := ""
		for _, a := range n.Attr {
			attrs = attrs + " " + a.Key + "=" + a.Val
		}
		fmt.Printf("<%s%s></%s> \n", n.Data, attrs, n.Data)
	}
}

