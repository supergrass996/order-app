package order

import (
	"errors"
	"testing"
)

type MockWriter struct {
	Saved []Order
	Err   error
}

func (m *MockWriter) SaveOrder(o Order) (Order, error) {
	if m.Err != nil {
		return Order{}, m.Err
	}
	o.ID = len(m.Saved) + 1
	m.Saved = append(m.Saved, o)
	return o, nil
}

type MockNotifier struct {
	Sent []string
	Err  error
}

func (m *MockNotifier) Send(customer, message string) error {
	if m.Err != nil {
		return m.Err
	}
	m.Sent = append(m.Sent, customer+": "+message)
	return nil
}

func TestCreateOrder_SavesAndNotifies(t *testing.T) {
	writer := &MockWriter{}
	notifier := &MockNotifier{}
	service := NewOrderService(writer, notifier)

	if err := service.CreateOrder("Иван", []string{"apple"}, 10.5); err != nil {
		t.Fatalf("ожидался nil, получено: %v", err)
	}
	if len(writer.Saved) != 1 {
		t.Fatalf("ожидался 1 заказ, сохранено: %d", len(writer.Saved))
	}
	if len(notifier.Sent) != 1 {
		t.Fatalf("ожидалось 1 уведомление, отправлено: %d", len(notifier.Sent))
	}
}

func TestCreateOrder_WriteError_NoNotification(t *testing.T) {
	writer := &MockWriter{Err: errors.New("db down")}
	notifier := &MockNotifier{}
	service := NewOrderService(writer, notifier)

	if err := service.CreateOrder("Иван", []string{"apple"}, 10.5); err == nil {
		t.Fatal("ожидалась ошибка от хранилища")
	}
	if len(notifier.Sent) != 0 {
		t.Fatal("уведомление не должно отправляться при ошибке записи")
	}
}
