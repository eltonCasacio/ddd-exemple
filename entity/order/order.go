package order

import (
	"errors"
	"fmt"
)

type order struct {
	id         string
	customerID string
	items      *[]item
	price      float64
}

func NewOrder(id, customerID string, items *[]item) (*order, error) {
	o := &order{
		id:         id,
		customerID: customerID,
		items:      items,
	}
	err := o.Validate()
	o.CalculateTotalPrice()
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *order) CalculateTotalPrice() {
	for _, v := range *o.items {
		o.price = o.price + (v.price * float64(v.quantity))
	}
}

func (o *order) GetPrice() float64 {
	return o.price
}

func (o *order) Validate() error {
	if o.customerID == "" {
		return errors.New("customer id is require")
	}
	if len(*o.items) < 1 {
		return errors.New("item must be greather than zero")
	}
	return nil
}

func (o *order) ToString() string {
	return fmt.Sprintf(`
	ID: %s
	customerID: %s
	items: %v
	`, o.id, o.customerID, o.items)
}
