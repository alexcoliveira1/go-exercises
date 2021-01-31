package main

import (
	"fmt"
	"sync"
	"time"
)

var x = 0

func increment(m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
}

func main() {
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		go increment(&m)
	}
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("final value of x", x)
}
