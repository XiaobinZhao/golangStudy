// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	caseCounts := map[string]int {
		"letter": 0,
		"space": 0,
		"control": 0,
		"number": 0,
		"mark": 0,
		"punct": 0,
		"symbol": 0,
		"other": 0,
	}
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		if unicode.IsLetter(r) {
			caseCounts["letter"]++
		} else if unicode.IsSpace(r) {
			caseCounts["space"]++
		} else if unicode.IsNumber(r) {
			caseCounts["number"]++
		} else if unicode.IsSymbol(r) {
			caseCounts["symbol"]++
		} else if unicode.IsMark(r) {
			caseCounts["mark"]++
		} else if unicode.IsControl(r) {
			caseCounts["control"]++
		} else if unicode.IsPunct(r) {
			caseCounts["punct"]++
		} else {
			caseCounts["other"]++
		}
		counts[r]++
		utflen[n]++
	}
	fmt.Printf("rune\tcount\n")
	for r, n := range counts {
		fmt.Printf("%q\t%d\n", r, n)
	}
	fmt.Printf("case\tcount\n")
	for k, n := range caseCounts {
		fmt.Printf("%s\t%d\n", k, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}