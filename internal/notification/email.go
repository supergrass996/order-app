package notification

import "fmt"

type EmailSender struct{}

func (e *EmailSender) Send(customer, message string) error {
	fmt.Printf("Email для %s: %s\n", customer, message)
	return nil
}
