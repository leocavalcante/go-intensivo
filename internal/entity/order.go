package entity

import "errors"

type Order struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

func NewOrder(id string, price, tax float64) (*Order, error) {
	o := &Order{
		ID:    id,
		Price: price,
		Tax:   tax,
	}

	err := o.Validate()
	if err != nil {
		return nil, err
	}

	o.FinalPrice = o.CalculateFinalPrice()

	return o, nil
}

func (o *Order) Validate() error {
	if o.ID == "" {
		return errors.New("id is required")
	}

	if o.Price <= 0 {
		return errors.New("price must be greater than zero")
	}

	if o.Tax <= 0 {
		return errors.New("tax must be greater than zero")
	}

	return nil // nulo
}

func (o *Order) CalculateFinalPrice() float64 {
	return o.Price + o.Tax
}
