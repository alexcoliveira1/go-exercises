package main

import (
	"errors"
	"fmt"
)

func SumPositiveNum(arr []int) (int, error) {
	if arr == nil {
		return 0, errors.New("")
	}
	res := 0
	for _, n := range arr {
		res += n
	}
	return res, nil
}

func main() {
	sum, err := SumPositiveNum([]int{1, 2, 3})
	if err == nil {
		fmt.Printf("Sum: %v\n", sum)
	}
}
