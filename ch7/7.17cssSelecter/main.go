package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// 练习 7.17： 扩展xmlselect程序以便让元素不仅可以通过名称选择，也可以通过它们CSS风格的属性进行选择。例如一个像这样
//
// <div id="page" class="wide">
// 的元素可以通过匹配id或者class，同时还有它的名称来进行选择。
// 分析： 支持原先的元素名称查找和支持id、class方式查找。id使用#{id}，比如#page; class使用.{class}，比如.wide
func main() {
	//strReader := strings.NewReader("<div id=\"page\" class=\"wide\"></div>")
	dec := xml.NewDecoder(os.Stdin)
	var stack []string // stack of element names
	attrFilter := os.Args[1]
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Fprintf(os.Stderr, "xmlselect: %v\n", err)
			os.Exit(1)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			stack = append(stack, tok.Name.Local) // push
			if strings.HasPrefix(attrFilter, ".") {  // class
				classFilter := attrFilter[1:]
				for _, attr := range tok.Attr {
					if attr.Name.Local == "class" && attr.Value == classFilter{
						fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
					}
				}
			} else if strings.HasPrefix(attrFilter, "#") {
				idFilter := attrFilter[1:]
				for _, attr := range tok.Attr {
					if attr.Name.Local == "id" && attr.Value == idFilter{
						fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
					}
				}
			}
		case xml.EndElement:
			stack = stack[:len(stack)-1] // pop
		case xml.CharData:
			if containsAll(stack, os.Args[1:]) {
				fmt.Printf("%s: %s\n", strings.Join(stack, " "), tok)
			}
		}
	}
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y []string) bool {
	for len(y) <= len(x) {
		if len(y) == 0 {
			return true
		}
		if x[0] == y[0] {
			y = y[1:]
		}
		x = x[1:]
	}
	return false
}

//!-
