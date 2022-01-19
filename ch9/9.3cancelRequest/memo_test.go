// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	memo "studygolang/ch9/9.3cancelRequest"
	"testing"
)


func Test(t *testing.T) {
	done := make(chan struct{})
	m := memo.New(memo.HttpGetBodyWithCancel)
	memo.Sequential(t, m, done)
}

func TestConcurrent(t *testing.T) {
	done := make(chan struct{})
	m := memo.New(memo.HttpGetBodyWithCancel)
	memo.Concurrent(t, m, done)
}
