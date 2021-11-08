package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

var depth int
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		elementPrefix := fmt.Sprintf("%*s<%s", depth*2, "", n.Data)
		attrs := ""
		for _, a := range n.Attr {
			attrVal := a.Val
			if a.Key == "href" || a.Key == "src" {
				attrVal = `"` + attrVal + `"`
			}
			attrs = attrs + " " + a.Key + "=" + attrVal
		}

		if n.FirstChild == nil {
			fmt.Printf("%s /> \n", elementPrefix + attrs)
		} else {
			fmt.Printf("%s> \n", elementPrefix + attrs)
		}
		depth++
		if n.FirstChild != nil && (n.FirstChild.Type == html.CommentNode || n.FirstChild.Type == html.TextNode) {
			if len(strings.TrimSpace(n.FirstChild.Data)) > 0 {
				fmt.Printf("%*s<%s>\n", depth*2 + 2, "", strings.TrimSpace(n.FirstChild.Data))
			}
		}
	}
}
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		if n.FirstChild != nil {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

//练习 5.7： 完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，
//文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。
//编写测试，验证程序输出的格式正确。（详见11章）
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
	forEachNode(doc, startElement, endElement)
}

