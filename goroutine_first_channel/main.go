package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go func(c chan int) {
		c <- 1
	}(ch)
	fmt.Printf("%v\n", <-ch)
}
