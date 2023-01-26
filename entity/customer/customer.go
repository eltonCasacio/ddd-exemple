package customer

import (
	"errors"
	"fmt"
)

/*
	ENTIDADE
	Customer é uma entidade e ela deve ser RICA, ou seja
	deve ter atributos, metodos, validacao, regra de negocio e PRINCIPALMENTE uma IDENTIDADE

	REGRA DE NEGÓCIO
	Podemos dizer que regra de negócio é um motivo, motivo para mudanca na entidade.

	Um ID em uma entidade representa uma IDENTIDADE, não tem relacao com ID do BD

*/

type customer struct {
	id      string
	name    string
	address address
	active  bool
}

func NewCustomer(id, name string, address address) (*customer, error) {
	c := customer{id: id, name: name, address: address}
	err := c.Validate()
	if err != nil {
		return nil, err
	}

	c.Active()
	return &c, nil
}

func (c *customer) Validate() error {
	if c.id == "" {
		return errors.New("id is required")
	}
	if c.name == "" {
		return errors.New("name is required")
	}

	if c.address.city == "" {
		return errors.New("address is required")
	}
	return nil
}

func (c *customer) ChangeName(name string) {
	c.name = name
}

func (c *customer) Active() {
	c.active = true
}

func (c *customer) Deactive() {
	c.active = false
}

func (c *customer) GetID() string {
	return c.id
}

func (c *customer) GetName() string {
	return c.name
}

func (c *customer) GetAddress() address {
	return c.address
}

func (c *customer) IsActive() bool {
	return c.active
}

func (c *customer) ToString() string {
	return fmt.Sprintf(`
	ID: %s
	Name: %s
	Is active: %t
	Address: %s
	`, c.GetID(), c.GetName(), c.IsActive(), c.address.ToString())
}
