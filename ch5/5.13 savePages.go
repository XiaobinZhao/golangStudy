package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func saveAsFile(s, url string) {
	fileName := ""
	url = strings.Replace(url, "http://books.studygolang.com", "", -1)
	// 此处没有创建目录，以url path为文件名，以_隔开
	if strings.Contains(url, "/") {
		paths := append([]string{"index"}, strings.Split(url, "/")...)
		fileName = strings.Join(paths, "_")
	} else {
		fileName = "index"
	}
	file, err := os.OpenFile(fileName + ".html", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	write.WriteString(s)
}

func extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	bodyStr := string(body)
	resp.Body.Close()
	doc, err := html.Parse(strings.NewReader(bodyStr))
	saveAsFile(bodyStr, url)

	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				// 去除一些页内锚点链接
				if a.Key == "href" && strings.Contains(a.Val, "#") {
					break
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue // ignore bad URLs
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode4(doc, visitNode, nil)
	return links, nil
}

//!-Extract

// Copied from gopl.io/ch5/outline2.
func forEachNode4(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode4(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 && len(worklist) < 50 { // 阻止无限查找下去
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item]{
				if strings.HasPrefix(item, "http://books.studygolang.com") {
					seen[item] = true
					worklist = append(worklist, f(item)...)
				}
			}
		}
	}
}

//!-breadthFirst

//!+crawl
func crawl(url string) []string {
	fmt.Printf("crawl: %s \n", url)
	list, err := extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

//!+main
// 练习5.13： 修改crawl，使其能保存发现的页面，必要时，可以创建目录来保存这些页面。
// 只保存来自原始域名下的页面。假设初始页面在golang.org下，就不要保存vimeo.com下的页面。
func main() {
	breadthFirst(crawl, []string{"http://books.studygolang.com"})
}

//!-main
