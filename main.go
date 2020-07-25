package main

import (
	"fmt"
	"log"

	"./pkg/product"
	"./storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	m, err := serviceProduct.GetByID(2)
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	fmt.Println(m)

}
