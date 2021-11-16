package ch6

// 练习 6.3： (*IntSet).UnionWith会用|操作符计算两个集合的并集，
// IntersectWith（交集：元素在A集合B集合均出现），
// DifferenceWith（差集：元素出现在A集合，未出现在B集合），
// SymmetricDifference（并差集：元素出现在A但没有出现在B，或者出现在B没有出现在A）。

func (s *IntSet) IntersectWith(t *IntSet) IntSet {
	var result IntSet
	for i, word := range t.words {
		if i < len(s.words) {
			result.words = append(result.words, s.words[i] & word)
		}
	}
	return result
}

func (s *IntSet) SymmetricDifference(t *IntSet) IntSet {
	var result IntSet
	for i, word := range s.words {
		if i < len(s.words){
			result.words = append(result.words, t.words[i] ^ word)
		} else {
			result.words = append(result.words, s.words[i])
		}
	}
	return result
}

func (s *IntSet) DifferenceWith(t *IntSet) IntSet {
	var result IntSet
	for i, word := range s.words {
		if i < len(t.words){
			result.words = append(result.words, word & (t.words[i] ^ word))
		} else {
			result.words = append(result.words, word)
		}
	}
	return result
}



