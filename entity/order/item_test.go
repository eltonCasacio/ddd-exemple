package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewItem(t *testing.T) {
	item, err := NewItem("1", "macbook air m1", 8.000, 3)
	assert.Nil(t, err)
	assert.Equal(t, item.GetID(), "1")
	assert.Equal(t, item.GetName(), "macbook air m1")
	assert.Equal(t, item.GetPrice(), 24.0)
	assert.Equal(t, item.id, "1")
}

func TestNewItem_WhenIDIsEmpty(t *testing.T) {
	item, err := NewItem("", "macbook air m1", 8.000, 3)
	assert.NotNil(t, err)
	assert.Nil(t, item)
	assert.EqualError(t, err, "id is require")
}

func TestNewItem_WhenNameIsEmpty(t *testing.T) {
	item, err := NewItem("1", "", 8.000, 3)
	assert.NotNil(t, err)
	assert.Nil(t, item)
	assert.EqualError(t, err, "name is require")
}

func TestNewItem_WhenPriceIsLessThanZero(t *testing.T) {
	item, err := NewItem("1", "macbook air m1", -1, 3)
	assert.NotNil(t, err)
	assert.Nil(t, item)
	assert.EqualError(t, err, "price cant be less than zero")
}
