package main

import (
	"fmt"
	"log"

	"github.com/eltoncasacio/ddd-modelagem-tatica/entity/customer"
	"github.com/eltoncasacio/ddd-modelagem-tatica/entity/order"
)

func main() {
	address, err := customer.NewAddress("país", "cidade", "bairro", "rua", "12234567", 23)
	if err != nil {
		log.Fatalln(err)
	}

	customer, err := customer.NewCustomer("1", "João", *address)
	if err != nil {
		log.Fatalln(err)
	}

	itens := order.NewArrayOfItems()

	item, _ := order.NewItem("1", "macbook air m1", 8.000, 2)
	itens = append(itens, *item)

	order, err := order.NewOrder("1", customer.GetID(), &itens)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Ordem\n", order.ToString())
	fmt.Println("Cliente\n", customer.ToString())
}
