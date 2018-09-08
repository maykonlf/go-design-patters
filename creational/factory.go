package creational

import (
	"github.com/pkg/errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	CASH       = 1
	DEBIT_CARD = 2
)

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case CASH:
		return &CashPayment{}, nil
	case DEBIT_CARD:
		return &DebitCardPayment{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized\n", method))
	}
}


type CashPayment struct{}

func (c *CashPayment) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using cash\n", amount)
}


type DebitCardPayment struct{}

func (c *DebitCardPayment) Pay(amount float32) string {
	return fmt.Sprintf("%0.2f paid using debit card\n", amount)
}
