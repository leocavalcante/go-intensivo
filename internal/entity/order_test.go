package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfItGetsAnErrorWhenIDIsBlank(t *testing.T) {
	_, err := NewOrder("", 1.0, 1.0)
	assert.Error(t, err, "invalid id")
}

func TestWithValidParameters(t *testing.T) {
	o, err := NewOrder("123", 1.0, 1.0)
	assert.Nil(t, err)
	assert.Equal(t, "123", o.ID)
	assert.Equal(t, 1.0, o.Price)
	assert.Equal(t, 1.0, o.Tax)
	assert.Equal(t, 2.0, o.FinalPrice)
}
