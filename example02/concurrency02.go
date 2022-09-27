package example02

import (
	"fmt"
	"sync"
	"time"
)

func asyncFunc(n int) {
	fmt.Println("asyncFunc:", n)
	time.Sleep(time.Second * 5)
}

func MainConcurrency() {
	var wg sync.WaitGroup // WaitGroup which wait for one or more goroutines to finish
	wg.Add(1) // how many goroutines it will wait ti complete
	go func () {
		asyncFunc(2)
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("syncFunc")
}
