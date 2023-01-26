package customer

/*
	Pense no objeto de valor como um tipo personalizado.
	Objeto de valor não tem regras de negócio e nem IDENTIFICADOR.
	Se válida
	Não tem ID
	Imutável
		Exemplo:
		Em um ENDEREÇO, se um atributo/estado for alterado o endereço muda. Não tem como alterar qualquer atributo sem mudar de endereço. Assim como o CEP ou CPF, RG, CNPJ, Número de Telefone…


	***********************************************************
	Depois que um endereço é criado não é possivel altera-lo,
	para alterar terá que criar outro e substituir o antigo.
	***********************************************************
*/

import (
	"errors"
	"fmt"
)

type address struct {
	country      string
	city         string
	neighborhood string
	street       string
	number       int
	zipcode      string
}

func NewAddress(country, city, neighborhood, street, zipcode string, number int) (*address, error) {
	a := address{
		country:      country,
		city:         city,
		neighborhood: neighborhood,
		street:       street,
		number:       number,
		zipcode:      zipcode,
	}

	err := a.Validate()
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func (a *address) Validate() error {
	if a.country == "" {
		return errors.New("country is required")
	}
	if a.city == "" {
		return errors.New("city is required")
	}
	if a.neighborhood == "" {
		return errors.New("neighborhood is required")
	}
	if a.street == "" {
		return errors.New("street is required")
	}
	if a.number == 0 {
		return errors.New("number is required")
	}
	if a.zipcode == "" {
		return errors.New("zipcode is required")
	}
	return nil
}

func (a *address) ToString() string {
	return fmt.Sprintf(`country: %s
	city: %s
	neighborhood: %s
	street: %s
	number: %d
	zipcode: %s
	neighborhood: %s
	`, a.country, a.city, a.neighborhood, a.street, a.number, a.zipcode, a.neighborhood)
}
