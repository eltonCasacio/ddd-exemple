package order

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	o, err := NewOrder("1", "1", &[]item{
		{
			id:       "1",
			name:     "item 1",
			price:    1000,
			quantity: 2,
		},
		{
			id:       "2",
			name:     "item 2",
			price:    1000,
			quantity: 4,
		},
	})
	assert.Nil(t, err)
	assert.Equal(t, o.id, "1")
	assert.Equal(t, o.customerID, "1")
	assert.Equal(t, len(*o.items), 2)
	assert.Equal(t, float64(6000), o.GetPrice())
}
