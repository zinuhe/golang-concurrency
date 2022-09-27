package example03

import (
	"fmt"
	"math/rand"
	"time"
)

func asyncReceiver(c chan int) {
	for {
		n := <-c
		fmt.Println("asyncReceiver:", n)
	}
}

func asyncSender(c chan int) {
	for {
		n := rand.Intn(10)
		c <- n
		fmt.Println("asyncSender:", n)
		time.Sleep(time.Millisecond * 100)
	}
}

func MainConcurrency() {
	c := make(chan int, 10) // we've created a channel with ten integers capacity.
	go asyncReceiver(c)
	go asyncSender(c)
	time.Sleep(time.Second * 2)
}
