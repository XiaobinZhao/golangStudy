package main

import "fmt"

// 练习5.15： 编写类似sum的可变参数函数max和min。考虑不传参时，max和min该如何处理，再编写至少接收1个参数的版本。

func max(vals ...int) (int, error){
	if len(vals) == 0 {
		return 0, nil
	} else {
		max := 0
		for _,v := range vals {
			if v > max {
				max = v
			}
		}
		return max, nil
	}
}

func min(vals ...int) (int, error){
	if len(vals) == 0 {
		return 0, nil
	} else {
		min := vals[0]
		for _,v := range vals {
			if v < min {
				min = v
			}
		}
		return min, nil
	}
}

func maxLeast1Param(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("return at least one argument")
	} else {
		max := vals[0]
		for _,v := range vals {
			if v > max {
				max = v
			}
		}
		return max, nil
	}
}
func minLeast1Param(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, fmt.Errorf("return at least one argument")
	} else {
		min := vals[0]
		for _,v := range vals {
			if v < min {
				min = v
			}
		}
		return min, nil
	}
}

func main() {
	fmt.Println(max(1,2,3))
	fmt.Println(max(4,9,3))
	fmt.Println(max())
	fmt.Println(maxLeast1Param(1,2,3))
	fmt.Println(maxLeast1Param(4,9,3))
	fmt.Println(maxLeast1Param())
	fmt.Println("==================")
	fmt.Println(min(1,2,3))
	fmt.Println(min(4,6,9))
	fmt.Println(min())
	fmt.Println(minLeast1Param(1,2,3))
	fmt.Println(minLeast1Param(4,2,9))
	fmt.Println(minLeast1Param())

}
