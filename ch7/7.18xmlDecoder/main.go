package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"strings"
)

// 练习 7.18： 使用基于标记的解码API，编写一个可以读取任意XML文档并构造这个文档所代表的通用节点树的程序。
// 节点有两种类型：CharData节点表示文本字符串，和 Element节点表示被命名的元素和它们的属性。
// 每一个元素节点有一个子节点的切片。
//
// 你可能发现下面的定义会对你有帮助。
//
// import "encoding/xml"
//
// type Node interface{} // CharData or *Element
//
// type CharData string
//
// type Element struct {
//     Type     xml.Name
//     Attr     []xml.Attr
//     Children []Node
// }


type Node interface{} // CharData or *Element

type CharData string

type Element struct {
	Type     xml.Name
	Attr     []xml.Attr
	Children []Node
}


func main() {
	//dec := xml.NewDecoder(os.Stdin)
	xmlStr := `
<xxx>
	<aaa id="z"></aaa>
	<bbb class="b">bbbb</bbb>
	<ccc>
		<ddd></ddd>
		<eee>eee</eee>
	</ccc>
</xxx>
`
	dec := xml.NewDecoder(strings.NewReader(xmlStr))
	var stack []Element
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
			startEle := Element{Type: tok.Name, Attr: tok.Attr, Children: []Node{}}
			stack = append(stack, startEle) // push
		case xml.EndElement:
			node := stack[len(stack)-1]  // top one
			stack = stack[:len(stack)-1] // pop top
			if len(stack) > 0 {
				parentNode := &(stack[len(stack)-1]) // get stack top
				parentNode.Children = append(parentNode.Children,  node)
			} else {
				printStack([]Node{node})
				break
			}
		case xml.CharData:
			charEle := CharData(tok)
			if len(stack) > 0 {
				parentNode := &(stack[len(stack)-1]) // get stack top
				parentNode.Children = append(parentNode.Children,  charEle)
			}
		}
	}
}

func printStack(node []Node) {
	for _,e := range node {
		endTag := ""

		switch ele := e.(type) {
		case Element:
			attrs := ""
			if len(ele.Attr) > 0 {
				for _, attr := range ele.Attr {
					attrs = attrs + attr.Name.Local + "=" + attr.Value + " "
				}
			}
			fmt.Printf("<%s %s>", ele.Type.Local, attrs)
			endTag = fmt.Sprintf("</%s>", ele.Type.Local)
			if len(ele.Children) > 0 {
				//fmt.Println()
				printStack(ele.Children)

			}
			fmt.Printf("</%s>", ele.Type.Local)
		case CharData:
			fmt.Printf("%s%s", ele, endTag)
		}
	}
}

