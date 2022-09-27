// go test
// go test -test.run=TestMainConcurrency03
// https://medium.com/inside-sumup/golang-goroutines-101-8e32e70a7410

package example03_test

import (
	"fmt"
	"testing"

	c03 "github.com/zinuhe/golang-concurrency/example03"
)

func TestMainConcurrency03(t *testing.T) {
	fmt.Println("test example concurrency03")

	c03.MainConcurrency()
}