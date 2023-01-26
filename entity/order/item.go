package order

import (
	"errors"
	"fmt"
)

var ItemType item

type item struct {
	id    string
	name  string
	price float64
}

func NewItem(id, name string, price float64) *item {
	return &item{
		id:    id,
		name:  name,
		price: price,
	}
}

func NewArrayOfItems() []item {
	return []item{}
}

func (i *item) Validate() error {
	if i.name == "" {
		return errors.New("name is require")
	}
	if i.price <= 0 {
		return errors.New("price must be greather then zero")
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
