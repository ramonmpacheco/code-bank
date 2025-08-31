package repository

import (
	"database/sql"
	"github.com/ramonmpacheco/code-bank/codebank/domain"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (tr *TransactionRepository) Save(transaction domain.Transaction, creditCard domain.CreditCard) error {
	stmt, err := tr.db.Prepare(`INSERT INTO transactions(id, credit_card, amount, status, description, store, created_at) 
								values($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		transaction.ID,
		transaction.CreditCardId,
		transaction.Amount,
		transaction.Status,
		transaction.Description,
		transaction.Store,
		transaction.CreatedAt,
	)
	if err != nil {
		return err
	}
	if transaction.Status == "approved" {
		err = tr.updateBalance(creditCard)
		if err != nil {
			return err
		}
	}

	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) updateBalance(creditCard domain.CreditCard) error {
	_, err := tr.db.Exec(`UPDATE credit_cards set balance = $1 where id = $2`, creditCard.Balance, creditCard.ID)
	if err != nil {
		return err
	}
	return nil
}

func (tr *TransactionRepository) CreateCreditCard(creditCard domain.CreditCard) error {
	stmt, err := tr.db.Prepare(`insert into credit_cards(id, name, number, expiration_month,expiration_year, CVV,balance, balance_limit) 
								values($1,$2,$3,$4,$5,$6,$7,$8)`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(
		creditCard.ID,
		creditCard.Name,
		creditCard.Number,
		creditCard.ExpirationMonth,
		creditCard.ExpirationYear,
		creditCard.CVV,
		creditCard.Balance,
		creditCard.Limit,
	)
	if err != nil {
		return err
	}
	err = stmt.Close()
	if err != nil {
		return err
	}
	return nil
}
