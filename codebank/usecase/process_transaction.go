package usecase

import (
	"github.com/ramonmpacheco/code-bank/codebank/domain"
	"github.com/ramonmpacheco/code-bank/codebank/dto"
	"time"
)

type TransactionUseCase struct {
	domain.TransactionRepository
}

func NewTransactionUseCase(transactionRepository domain.TransactionRepository) TransactionUseCase {
	return TransactionUseCase{transactionRepository}
}

func (u TransactionUseCase) ProcessTransaction(dto dto.Transaction) (domain.Transaction, error) {
	creditCard := u.hydrateCreditCard(dto)
	ccBalanceAndLimit, err := u.TransactionRepository.GetCreditCard(*creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	creditCard.ID = ccBalanceAndLimit.ID
	creditCard.Limit = ccBalanceAndLimit.Limit
	creditCard.Balance = ccBalanceAndLimit.Balance

	t := u.NewTransaction(dto, ccBalanceAndLimit)
	t.ProcessAndValidate(creditCard)

	err = u.TransactionRepository.Save(*t, *creditCard)
	if err != nil {
		return domain.Transaction{}, err
	}
	return *t, nil
}

func (u TransactionUseCase) hydrateCreditCard(dto dto.Transaction) *domain.CreditCard {
	creditCard := domain.NewCreditCard()

	creditCard.Name = dto.Name
	creditCard.Number = dto.Number
	creditCard.ExpirationMonth = dto.ExpirationMonth
	creditCard.ExpirationYear = dto.ExpirationYear
	creditCard.CVV = dto.CVV

	return creditCard
}

func (u TransactionUseCase) NewTransaction(dto dto.Transaction, cc domain.CreditCard) *domain.Transaction {
	t := domain.NewTransaction()
	t.CreditCardId = cc.ID
	t.Amount = dto.Amount
	t.Store = dto.Store
	t.Description = dto.Description
	t.CreatedAt = time.Now()
	return t
}
