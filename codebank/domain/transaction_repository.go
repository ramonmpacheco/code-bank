package domain

type TransactionRepository interface {
	Save(Transaction, CreditCard) error
	GetCreditCard(CreditCard) (CreditCard, error)
	CreateCreditCard(CreditCard) error
}
