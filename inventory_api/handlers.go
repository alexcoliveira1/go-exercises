package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(getAllProducts()); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func addProductHandler(w http.ResponseWriter, r *http.Request) {
	pI, err := getProductInventoryFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	addProduct(*pI)

	w.Write([]byte("true"))
}

func updateProductHandler(w http.ResponseWriter, r *http.Request) {
	updatedProductInventory, err := getProductInventoryFromRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := mux.Vars(r)["id"]

	_, err = updateProductByID(*updatedProductInventory, id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("true"))
}

func removeProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, err := removeProductByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("true"))
}

func getProductHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	_, pI, err := getProductInventoryByID(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := json.NewEncoder(w).Encode(pI); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
