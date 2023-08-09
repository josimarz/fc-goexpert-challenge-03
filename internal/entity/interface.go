package entity

type OrderRepository interface {
	Save(order *Order) error
	FindAll() ([]Order, error)
}
