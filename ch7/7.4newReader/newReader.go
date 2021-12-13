package main

import (
	"io"
)




// 练习 7.4： strings.NewReader函数通过读取一个string参数返回一个满足io.Reader接口类型的值（和其它值）。
// 实现一个简单版本的NewReader，用它来构造一个接收字符串输入的HTML解析器（§5.2）
// 分析：strings.NewReader返回一个 io.Reader接口类型的值：也就是说要返回一个实现了`Read(p []byte) (n int, err error)`方法的实例

type MyHTML struct {
	strIn string
}

func (m *MyHTML) Read(p []byte) (n int, err error) {
	return
}

func NewReader(s string) io.Reader {
	var myHTML *MyHTML
	myHTML.strIn = s

	return myHTML

}


