package entity

type OrderRepository interface {
	Save(o *Order) error
	GetTotal() (int, error)
}
