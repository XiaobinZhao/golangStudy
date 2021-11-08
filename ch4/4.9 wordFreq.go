package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		world := scanner.Text()
		wordCounts[world] ++
	}
	fmt.Printf("world\tcount\n")
	for w, n := range wordCounts {
		fmt.Printf("%q\t%d\n", w, n)
	}
}