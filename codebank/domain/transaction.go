package domain

import "time"
import uuid "github.com/satori/go.uuid"

type Transaction struct {
	ID           string
	Amount       float64
	Status       string
	Description  string
	Store        string
	CreditCardId string
	CreatedAt    time.Time
}

func NewTransaction() *Transaction {
	return &Transaction{
		ID:        uuid.NewV4().String(),
		CreatedAt: time.Now(),
	}
}

func (transaction *Transaction) ProcessAndValidate(creditCard *CreditCard) {
	if transaction.Amount+creditCard.Balance > creditCard.Limit {
		transaction.Status = "rejected"
		return
	}
	transaction.Status = "approved"
	creditCard.Balance = creditCard.Balance + transaction.Amount
}
