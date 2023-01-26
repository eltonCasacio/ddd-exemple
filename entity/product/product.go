package product

import "errors"

type product struct {
	id    string
	name  string
	price float64
}

func NewProduct(id, name string, price float64) (*product, error) {
	p := &product{
		id:    id,
		name:  name,
		price: price,
	}
	if err := p.Validate(); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *product) Validate() error {
	if p.id == "" {
		return errors.New("id is required")
	}
	if p.id == "0" {
		return errors.New("id cant be zero")
	}
	if p.name == "" {
		return errors.New("name is required")
	}
	if p.price < 0 {
		return errors.New("price must be greather than zero")
	}
	return nil
}

func (p *product) ChangePrice(newPrice float64) error {
	if newPrice < 0 {
		return errors.New("price must be greather than zero")
	}

	p.price = newPrice
	return nil
}

func (p *product) GetID() string {
	return p.id
}

func (p *product) GetName() string {
	return p.name
}

func (p *product) GetPrice() float64 {
	return p.price
}
