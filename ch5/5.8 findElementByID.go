package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"os"
	"strings"
)

func startElement2(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {

		attrs := ""
		for _, a := range n.Attr {
			attrs = attrs + " " + a.Key + "=" + a.Val
		}
		if strings.Contains(attrs, "class=dropdown-toggle") {
			fmt.Printf("<%s%s>", n.Data, attrs)
			return true
		}
	}
	return false
}

func endElement2(n *html.Node, className string) bool {
	if n.Type == html.ElementNode {
		for _, a := range n.Attr {
			if a.Key == "class" && a.Val == "dropdown-toggle" {
				fmt.Printf("</%s>\n", n.Data)
				return true
			}
		}

	}
	return false
}

var isExists = false

func forEachNode2(n *html.Node, pre, post func(n *html.Node, className string) bool, className string) {
	fmt.Printf("遍历节点：%s \n", n.Data)
	if pre != nil {
		isExists = pre(n, className)
	}
	for c := n.FirstChild; c != nil && !isExists; c = c.NextSibling {
		forEachNode2(c, pre, post, className)
	}

	if post != nil {
		post(n, className)
	}
}


func ElementByID(doc *html.Node, className string) {
	forEachNode2(doc, startElement2, endElement2, className)
}


//练习 5.8： 修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，
//根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。
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
	ElementByID(doc, "dropdown-toggle")
}

