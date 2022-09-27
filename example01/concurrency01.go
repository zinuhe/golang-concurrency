package example01

import (
	"fmt"
	"time"
)

func asyncFunc(n int) {
	fmt.Println("asyncFunc:", n)
}

func MainConcurrency() {
	go asyncFunc(0)
	time.Sleep(time.Second)
}
