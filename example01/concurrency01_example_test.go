// go test
// go test -test.run=TestMainConcurrency01
// https://medium.com/inside-sumup/golang-goroutines-101-8e32e70a7410

package example01_test

import (
	"fmt"
	"testing"

	c01 "github.com/zinuhe/golang-concurrency/example01"
)

func TestMainConcurrency01(t *testing.T) {
	fmt.Println("test example concurrency01")

	c01.MainConcurrency()
}