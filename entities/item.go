package entities

import "errors"

type item struct {
	id    string
	name  string
	price int64
}

func NewItem(name string, price int64) *item {
	return &item{
		id:    "1",
		name:  name,
		price: price,
	}
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
