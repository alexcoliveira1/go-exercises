package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type product struct {
	ID    string  `json:"id"`
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type productInventory struct {
	Product  product `json:"product"`
	Quantity int     `json:"quantity"`
}

func main() {
	initInventory()
	r := mux.NewRouter()
	r.HandleFunc("/products", getAllProductsHandler).Methods("GET")
	r.HandleFunc("/products", addProductHandler).Methods("POST")
	r.HandleFunc("/products/{id}", getProductHandler).Methods("GET")
	r.HandleFunc("/products/{id}", updateProductHandler).Methods("PUT")
	r.HandleFunc("/products/{id}", removeProductHandler).Methods("DELETE")
	http.Handle("/", r)
	http.ListenAndServe(":8090", nil)
}
