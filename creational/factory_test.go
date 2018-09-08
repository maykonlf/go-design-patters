package creational

import (
	"testing"
	"strings"
)

func TestPaymentMethodCash(t *testing.T) {
	payment, err := GetPaymentMethod(CASH)
	if err != nil {
		t.Fatal("A payment method of type 'CASH' must exists")
	}

	msg := payment.Pay(10.30)
	if !strings.Contains(msg, "paid using cash") {
		t.Error("The cash payment method message wasn't correct")
	}
}

func TestPaymentMethodDebitCard(t *testing.T) {
	payment, err := GetPaymentMethod(DEBIT_CARD)

	if err != nil {
		t.Error("A payment method of type 'DEBIT_CARD' must exists")
	}

	msg := payment.Pay(22.30)

	if !strings.Contains(msg, "paid using debit card") {
		t.Error("The debit card payment method message wasn't correct")
	}
}

func TestNonExistentPaymentMethod(t *testing.T)  {
	_, err := GetPaymentMethod(90)

	if err == nil {
		t.Error("A payment method with ID 90 must return an error")
	}
}
