package payment

import (
	"shareddomain/money"
)

type ChargeCreditCard struct {
	paymentService PaymentService
}

func NewChargeCreditCard(paymentService PaymentService) *ChargeCreditCard {
	return &ChargeCreditCard{paymentService: paymentService}
}

func (uc *ChargeCreditCard) Execute(inputCard Card, amount float64, currency string) *payment.Transaction {
	chargingAmount := money.New(currency, amount)
	transaction := CreateTransaction(chargingAmount)
	chargingCard, err := NewCard(inputCard)
	if err != nil {
		_ = transaction.MarkFailed(err.Error())

		return transaction
	}

	uc.paymentService.Charge(chargingCard, transaction)

	return transaction
}
