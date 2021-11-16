package ch6

import (
	"bytes"
	"fmt"
)

// 练习 6.5： 我们这章定义的IntSet里的每个字都是用的uintMachineBit类型，但是MachineBit位的数值可能在32位的平台上不高效。修改程序，使其使用uint类型，
// 这种类型对于32位平台来说更合适。当然了，这里我们可以不用简单粗暴地除MachineBit，可以定义一个常量来决定是用32还是MachineBit，
// 这里你可能会用到平台的自动判断的一个智能表达式：32 << (^uint(0) >> 63)
const MachineBit = 32 << (^uint(0) >> 63)

type IntSet2 struct {
	words []uint64
}

func (s *IntSet2) Has(x int) bool {
	word, bit := x/MachineBit, uint(x%MachineBit)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

func (s *IntSet2) Add(x int) {
	word, bit := x/MachineBit, uint(x%MachineBit)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

func (s *IntSet2) UnionWith(t *IntSet2) {

	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet2) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < MachineBit; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", MachineBit*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// return the number of elements
func (s *IntSet2) Len() int {
	sum := 0
	for _, word := range s.words {
		for j := 0; j < MachineBit; j++ {
			if word&(1<<uint(j)) != 0 {
				sum++
			}
		}
	}
	return sum
}
// remove x from the set
func (s *IntSet2) Remove(x int) {
	word, bit := x/MachineBit, uint(x%MachineBit)
	s.words[word] &= s.words[word] ^ (1 << bit)
}
// remove all elements from the set
func (s *IntSet2) Clear() {
	for i, word := range s.words {
		for j := 0; j < MachineBit; j++ {
			if word&(1<<uint(j)) != 0 {
				s.words[i] ^= 1 << uint(j)
			}
		}
	}
}

// return a copy of the set
func (s *IntSet2) Copy() *IntSet2 {
	var copyToReturn IntSet2
	for _, word := range s.words {
		copyToReturn.words = append(copyToReturn.words, word)
	}
	return &copyToReturn
}

