package main

import (
	"fmt"
	"os"
	"strconv"
	"studygolang/weightconv"
)

const 是多少 = "汉字命名变量"

func main() {
	for _, arg := range os.Args[1:] {
		arg64, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr,"参数转化为float64异常： %v", err)
			os.Exit(1)
		}
		lb := weightconv.Lb(arg64)
		kg := weightconv.Kilogram(arg64)
		fmt.Printf("%s = %s, %s = %s \n", lb, weightconv.LbToKg(lb), kg, weightconv.KgToLb(kg))
	}
}

func init() {
	var a [3]int             // array of 3 integers
	fmt.Println(a[0])        // print the first element
	fmt.Println(a[len(a)-1]) // print the last element, a[2]

	// Print the indices and elements.
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}

	// Print the elements only.
	for _, v := range a {
		fmt.Printf("%d\n", v)
	}
}

func init() {
	var f float32 = 16777216 // 1 << 24
	fmt.Println(f == f+1)    // "true"!
}