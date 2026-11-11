package notification

import "fmt"

type SMSSender struct{}

func (s *SMSSender) Send(customer, message string) error {
	fmt.Printf("SMS для %s: %s\n", customer, message)
	return nil
}
