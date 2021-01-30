package main

import (
	"testing"
)

type TestCase struct {
	input    [2]int
	expected int
}

func TestIntMin(t *testing.T) {
	cases := []TestCase{
		TestCase{[2]int{0, 1}, 0},
		TestCase{[2]int{1, 0}, 0},
		TestCase{[2]int{1, 1}, 1},
	}
	for _, c := range cases {
		actual := IntMin(c.input[0], c.input[1])
		if actual != c.expected {
			t.Fatalf("Expected %d, got %d", c.expected, actual)
		}
	}
}
