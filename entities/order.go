package entities

import "errors"

type order struct {
	id         string
	customerID string
	items      []item
}

func NewOrder(customerID string, items []item) (*order, error) {
	o := &order{
		id:         "1",
		customerID: customerID,
		items:      items,
	}
	err := o.Validate()
	if err != nil {
		return nil, err
	}
	return o, nil
}

func (o *order) Validate() error {
	if o.customerID == "" {
		return errors.New("customer id is require")
	}
	if len(o.items) < 1 {
		return errors.New("item must be greather than zero")
	}
	return nil
}
