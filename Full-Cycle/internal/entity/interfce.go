package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTOtal() (int, error)
}
