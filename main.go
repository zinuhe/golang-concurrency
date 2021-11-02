package main

import (
	"fmt"
	"time"
)

func main (){
    fmt.Println("Concurrency")
    // https://www.youtube.com/watch?v=yNOe3STbtGE

    //Simple one
    // out1 :=  make(chan string)
    // go process("order", out1)
    // for msg := range out1 {
    // 	fmt.Println(msg)
    // }

    // func process(item string, out chan string) {
    //     defer close(out)
    //     for i:=0; i< 5; i++ {
    //         time.Sleep(time.Second / 2)
    //         out <- item
    //     }
    // }

    //Not that simple
    out1 :=  make(chan string)
    out2 :=  make(chan string)

    go func() {
        for {
            time.Sleep(time.Second / 2)
            out1 <- "oder processed, half second"  //You will see this twice in the output for each of out2
        }
    }()

    go func() {
        for {
            time.Sleep(time.Second)
            out2 <- "refund processed, full second"
        }
    }()

    for {
        select {
        case msg := <- out1:  //If channel 1 has data ready it will go out
            fmt.Println(msg)
        case msg := <- out2:  //If channel 2 has data ready it will go out
            fmt.Println(msg)
        }
    }
}
