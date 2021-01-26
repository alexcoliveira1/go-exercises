package main

import (
	"fmt"
)

func filterAge(lower, upper int, ageList []int) []int {
	response := make([]int, 0)
	for _, age := range ageList {
		if age >= lower && age <= upper {
			response = append(response, age)
		}
	}
	return response
}

func main() {
	s := filterAge(2, 4, []int{1, 2, 3, 4, 5})
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
