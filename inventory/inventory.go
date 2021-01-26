package main

import (
	"fmt"
)

type DuplicateProductError struct {
	ID string
}

func (e *DuplicateProductError) Error() string { return e.ID + ": already exist" }

type Product struct {
	ID   string
	Name string
}

type Inventory struct {
	Products map[string]Product
}

func (iv *Inventory) Add(product Product) (bool, *DuplicateProductError) {
	_, found := iv.Products[product.ID]
	if found {
		return false, &DuplicateProductError{product.ID}
	}
	iv.Products[product.ID] = product
	return true, nil
}

func main() {
	inventory := Inventory{make(map[string]Product)}
	inventory.Add(Product{"1", "product1"})
	_, err := inventory.Add(Product{"1", "product1"})
	fmt.Printf("%v\n", err)
	fmt.Printf("%v\n", inventory)
}
