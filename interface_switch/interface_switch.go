package main

import (
	"fmt"
)

func getCustomMessage(i interface{}) string {
	switch i.(type) {
	case int:
		return "Natural numbers"
	case string:
		return "It stores a poetry"
	default:
		return "Unknown"
	}
}

func main() {
	fmt.Printf("%v\n", getCustomMessage(1))
	fmt.Printf("%v\n", getCustomMessage("1"))
	fmt.Printf("%v\n", getCustomMessage(1.1))
}
