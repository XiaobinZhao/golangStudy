package main

import "fmt"

func myRotate(s []string) {
	for i, _ := range s {
		if i > len(s)/2 {
			break
		}
		s[i], s[len(s)-i-1] = s[len(s)-i-1], s[i]
		fmt.Println(s)
	}
}

func main() {
	ori := []string{"a", "b", "c"}
	fmt.Printf("origin string: %v \n", ori)
	myRotate(ori)
	fmt.Printf("after string: %v \n", ori)
}