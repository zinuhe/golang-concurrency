//go:build 02_tests
// +build 02_tests

// go test
// go test -run TestMainConcurrency02
// go test -tags=02_tests -v
// https://medium.com/inside-sumup/golang-goroutines-101-8e32e70a7410

package example02_test

import (
	"fmt"
	"testing"

	c02 "github.com/zinuhe/golang-concurrency/example02"
)

func TestMainConcurrency02(t *testing.T) {
	fmt.Println("test example concurrency02")

	c02.MainConcurrency()
}