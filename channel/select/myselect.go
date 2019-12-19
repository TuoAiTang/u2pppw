package main

import (
	"fmt"
	"time"
)

func main() {
	var c1 = make(chan int)
	var c2 = make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Millisecond * 500)
			c2 <- i
			i++
			c1 <- i
			i++
		}
	}()
	for {
		select {
		case n := <-c1:
			fmt.Println(n)
		case n := <-c2:
			fmt.Println(n)
		}
	}

}
