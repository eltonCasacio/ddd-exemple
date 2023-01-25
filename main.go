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

	customer, err := entities.NewCustomer("1", "João", *address)
	if err != nil {
		log.Fatalln(err)
	}

	itens := entities.NewArrayOfItems()

	item := entities.NewItem("1", "macbook air m1", 8.000)
	itens = append(itens, *item)

	order, err := entities.NewOrder("1", customer.GetID(), itens)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ordem\n", order.ToString())
	fmt.Println("Cliente\n", customer.ToString())
}
