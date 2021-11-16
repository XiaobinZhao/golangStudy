package ch6

// 练习 6.2： 定义一个变参方法(*IntSet).AddAll(...int)，这个方法可以添加一组IntSet，比如s.AddAll(1,2,3)。
func (s *IntSet) AddAll(xs ...int) {
	for _,x := range xs {
		s.Add(x)
	}
}

