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
	ms, err := serviceProduct.GetAll()
	if err != nil {
		log.Fatalf("product.GetAll: %v", err)
	}

	for row := range ms {
		fmt.Println(ms[row])
	}

}
