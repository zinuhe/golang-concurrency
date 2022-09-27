package example04

import (
	"fmt"
	"math/rand"
	"time"
)

func receiver(c chan int) {
	for {
		if shouldScale(c) < 0 {
			fmt.Println("Scale down")
			return // ends goroutine to avoid underload
		}
		<-c                                             // get order
		r := rand.Intn(100) + 300                       // returns an integer between 300 and 400
		time.Sleep(time.Duration(r) * time.Millisecond) // process order
	}
}

func sender(c chan int) {
	for {
		c <- 1                                          // adds order to queue
		s := rand.Intn(100) + 100                       // returns an integer between 100 and 200
		time.Sleep(time.Duration(s) * time.Millisecond) // sleeps between 100-200
	}
}

func shouldScale(c chan int) int {
	capacity := cap(c)        // gets channel capacity
	orders := len(c)          // gets how many orders we have to process
	if orders >= capacity-5 { // if it's almost full
		return 1
	} else if orders <= 5 { // if it's almost empty
		return -1
	}
	return 0 // it's just fine
}

func MainConcurrency() {
	c := make(chan int, 20)
	go sender(c)

	for {
		l := len(c) // gets how many items the channel has
		fmt.Println(l)
		if shouldScale(c) > 0 {
			fmt.Println("Scale up")
			go receiver(c) // starts goroutine to avoid overload
		}
		time.Sleep(time.Second)
	}
}