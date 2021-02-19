package main

import (
	"encoding/json"
	"net/http"
)

var inventory []productInventory

type notFoundError struct {
	Name string
}

func (e *notFoundError) Error() string { return e.Name + ": not found" }

func initInventory() {
	inventory = make([]productInventory, 0)
	inventory = append(inventory, productInventory{product{"1", "code1", "name1", 1.99}, 1})
}

func getAllProducts() []productInventory {
	return inventory
}

func addProduct(pI productInventory) {
	inventory = append(inventory, pI)
}

func updateProductByID(updatedProductInventory productInventory, id string) (bool, *notFoundError) {
	productFoundIndex, _, err := getProductInventoryByID(id)

	if err != nil {
		return false, err
	}

	inventory[productFoundIndex] = updatedProductInventory

	return true, nil
}

func removeProductByID(id string) (bool, *notFoundError) {
	productFoundIndex, _, err := getProductInventoryByID(id)

	if err != nil {
		return false, err
	}

	inventory = append(inventory[:productFoundIndex], inventory[productFoundIndex+1:]...)

	return true, nil
}

func getProductInventoryByID(id string) (int, *productInventory, *notFoundError) {
	productFoundIndex := -1

	for i, product := range inventory {
		if product.Product.ID == id {
			productFoundIndex = i
			break
		}
	}

	if productFoundIndex == -1 {
		return -1, nil, &notFoundError{"Product not found!"}
	}

	return productFoundIndex, &inventory[productFoundIndex], nil
}

func getProductInventoryFromRequest(r *http.Request) (*productInventory, error) {
	var pI productInventory
	if err := json.NewDecoder(r.Body).Decode(&pI); err != nil {
		return nil, err
	}
	return &pI, nil
}
