package treesort

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
	//appendValues(make([]int, 0), root)
}
// 为在gopl.io/ch4/treesort（§4.4）中的*tree类型实现一个String方法去展示tree类型的值序列。
func (t *tree) String1() string {
	var values []int
	values = appendValues(values, t)
	var buf bytes.Buffer
	buf.WriteByte('[')
	for _,v := range values {
		fmt.Fprintf(&buf, "%d,", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func (t *tree) String2() string {
	str := ""
	if t == nil {
		return str
	} else {
		str += t.left.String2()
		fmt.Sprintf("%s %d", str, t.value)
		str += t.right.String2()
	}
	return str
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

