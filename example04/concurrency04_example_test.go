// go test
// go test -test.run=TestMainConcurrency04
// go test ./...

// https://medium.com/inside-sumup/golang-goroutines-101-8e32e70a7410

package example04_test

import (
	"fmt"
	"testing"

	c04 "github.com/zinuhe/golang-concurrency/example04"
)

func TestMainConcurrency04(t *testing.T) {
	fmt.Println("test example concurrency04")

	c04.MainConcurrency()
}