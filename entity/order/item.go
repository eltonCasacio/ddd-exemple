package order

import (
	"errors"
	"fmt"
)

var ItemType item

type item struct {
	id       string
	name     string
	price    float64
	quantity int
}

func NewItem(id, name string, price float64, quantity int) (*item, error) {
	i := &item{
		id:       id,
		name:     name,
		price:    price * float64(quantity),
		quantity: quantity,
	}
	if err := i.Validate(); err != nil {
		return nil, err
	}
	return i, nil
}

func NewArrayOfItems() []item {
	return []item{}
}

func (i *item) GetID() string {
	return i.id
}

func (i *item) GetName() string {
	return i.name
}

func (i *item) GetPrice() float64 {
	return i.price
}

func (i *item) GetQuantity() int {
	return i.quantity
}

func (i *item) Validate() error {
	if i.id == "" {
		return errors.New("id is require")
	}
	if i.name == "" {
		return errors.New("name is require")
	}
	if i.price < 0 {
		return errors.New("price cant be less than zero")
	}
	if i.quantity <= 0 {
		return errors.New("quantity must be greather then zero")
	}
	return nil
}

func (i *item) ToString() string {
	return fmt.Sprintf(`
	ID: %s
	Name: %s
	items: %f
	`, i.id, i.name, i.price)
}
