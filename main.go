package main

import (
	"fmt"
	"log"

	"github.com/eltoncasacio/ddd-modelagem-tatica/entities"
)

func main() {
	address, err := entities.NewAddress("país", "cidade", "bairro", "rua", "12234567", 23)
	if err != nil {
		log.Fatalln(err)
	}

	customer, err := entities.NewCustomer("João", *address)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(customer.ToString())
}
