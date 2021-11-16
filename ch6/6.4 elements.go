package ch6

// 练习6.4: 实现一个Elems方法，返回集合中的所有元素，用于做一些range之类的遍历操作。

func (s *IntSet) Elems() []int {
	var result []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				result = append(result, 64*i+j)
			}
		}
	}
	// 方法使用值返回还是指针返回？
	// 1. 有状态的对象必须使用指针返回，如系统内置的 sync.WaitGroup、sync.Pool 之类的值，在 Go 中有些结构体中会显式存在 noCopy 字段提醒不能进行值拷贝
	// 2. 生命周期短的对象使用值返回，如果对象的生命周期存在比较久或者对象比较大，可以使用指针返回；
	// 3. 大对象推荐使用指针返回，对象大小临界值需要在具体平台进行基准测试得出数据；
	// 4. 参考一些大的开源项目中的使用方式，比如 kubernetes、docker 等；
	return result
}
