package order

import (
	"fmt"
	"strings"
)

// Order — модель заказа.
type Order struct {
	ID       int
	Customer string
	Products string
	Total    float64
	Status   string
}

// RepositoryInitializer отвечает за создание таблиц.
type RepositoryInitializer interface {
	Init() error
}

// RepositoryWriter отвечает за запись данных.
type RepositoryWriter interface {
	SaveOrder(order Order) (Order, error)
}

// Notifier — общий контракт для всех отправителей сообщений.
type Notifier interface {
	Send(customer, message string) error
}

// OrderService зависит только от абстракций (DIP).
type OrderService struct {
	repo RepositoryWriter
	nf   Notifier
}

func NewOrderService(repo RepositoryWriter, nf Notifier) *OrderService {
	return &OrderService{repo: repo, nf: nf}
}

func (s *OrderService) CreateOrder(customer string, products []string, total float64) error {
	order := Order{
		Customer: customer,
		Products: strings.Join(products, ", "),
		Total:    total,
		Status:   "pending",
	}

	saved, err := s.repo.SaveOrder(order)
	if err != nil {
		return err
	}

	return s.nf.Send(customer, fmt.Sprintf("Ваш заказ #%d успешно создан!", saved.ID))
}
