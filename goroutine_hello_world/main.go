package main

import (
	"fmt"
)

func main() {
	// Time will not be printed because the program will exit first
	go func() {
		fmt.Println("Hello World")
	}()
	fmt.Println("main function")
}
