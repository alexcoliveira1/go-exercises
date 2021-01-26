package main

import (
	"fmt"
)

type Calculator struct {
	result int
}

func (c *Calculator) Add(n int) int {
	c.result += n
	return c.result
}

func (c *Calculator) Subtract(n int) int {
	c.result -= n
	return c.result
}

func (c *Calculator) Multiply(n int) int {
	c.result *= n
	return c.result
}

func (c *Calculator) Divide(n int) int {
	if n == 0 {
		panic("Can't divide by 0")
	}
	c.result /= n
	return c.result
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Sorry for trying to break the rule:", r)
		}
	}()
	calc := Calculator{2}
	fmt.Printf("%v\n", calc.Divide(0))
}
