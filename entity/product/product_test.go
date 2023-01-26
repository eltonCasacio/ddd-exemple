package product

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("1", "macboook air m1", 1.2)
	assert.Nil(t, err)
	assert.Equal(t, product.GetID(), "1")
	assert.Equal(t, product.GetName(), "macboook air m1")
	assert.EqualValues(t, product.GetPrice(), 1.20)
}

func TestNewProduct_WhenIDIsEmpty(t *testing.T) {
	product, err := NewProduct("", "macboook air m1", 1.2)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.EqualError(t, err, "id is required")
}

func TestNewProduct_WhenIDIsZero(t *testing.T) {
	product, err := NewProduct("0", "macboook air m1", 1.2)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.EqualError(t, err, "id cant be zero")
}

func TestNewProduct_WhenNameIsEmpty(t *testing.T) {
	product, err := NewProduct("1", "", 1.2)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.EqualError(t, err, "name is required")
}

func TestNewProduct_WhenPriceLessThanZero(t *testing.T) {
	product, err := NewProduct("1", "macboook air m1", -1)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.EqualError(t, err, "price must be greather than zero")
}

func TestNewProduct_WhenPriceZeroOrGreather(t *testing.T) {
	product, err := NewProduct("1", "macboook air m1", 0)
	assert.Nil(t, err)
	assert.NotNil(t, product)
}

func TestChangePrice(t *testing.T) {
	product, _ := NewProduct("1", "macboook air m1", 0)
	err := product.ChangePrice(10.55)
	assert.Nil(t, err)
	assert.EqualValues(t, product.GetPrice(), 10.55)
}

func TestChangePrice_LessThanZero(t *testing.T) {
	product, _ := NewProduct("1", "macboook air m1", 0)
	err := product.ChangePrice(-10.55)
	assert.NotNil(t, err)
	assert.EqualError(t, err, "price must be greather than zero")
}
