package main

import (
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
	"unicode"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(url, doc)
	return
}
func countWordsAndImages(url string, n *html.Node) (words, images int) {
	datas := map[string][]string{
		"word": []string{},
		"image": []string{},
	}
	datas = visit5(datas, n)
	fmt.Printf("==========================from url:%v get worlds: \n", url )
	for index, value := range datas["word"] {
		fmt.Printf("%v: %v\n", index, value)
	}
	_words := strings.Join(datas["word"], "")
	fmt.Printf("===========================After join: \n %s \n", _words)
	wordCounts := make(map[string]int)
	for _, v := range _words {
		if unicode.Is(unicode.Han, v) {
			wordCounts[string(v)]++
		}
	}

	fmt.Printf("============================_words frequency: \n" )
	for index, value := range wordCounts {
		fmt.Printf("%v: %v\n", index, value)
	}

	for _, v := range wordCounts {
		words += v
	}
	images = len(datas["image"])
	return
}

func visit5(datas map[string][]string, n *html.Node) map[string][]string{
	if n.Type == html.TextNode && len(strings.TrimSpace(n.Data)) > 0 {
		datas["word"] = append(datas["word"], strings.TrimSpace(n.Data))
	}
	if n.Data == "img" {
		datas["image"] = append(datas["image"], n.Data)
	}
	if n.Data == "script" {
		return datas
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		datas = visit5(datas, c)
	}
	return datas
}

func main() {
	// 由于网页是中文内容，所以本代码修改为统计汉字个数
	fmt.Print(CountWordsAndImages("http://books.studygolang.com/"))
}
