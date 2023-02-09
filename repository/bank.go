package repository

import (
	"abanku/model"
	"database/sql"
)

type BankRepo interface {
	GetAllAccounts() (response *[]model.Account, err error)
	InsertTransaction(transaction *model.Transaction) (err error)
}

type bankRepo struct {
	db *sql.DB
}

func NewBankRepo(db *sql.DB) BankRepo {
	return &bankRepo{
		db: db,
	}
}

func (b *bankRepo) GetAllAccounts() (response *[]model.Account, err error) {
	var accounts []model.Account

	rows, err := b.db.Query(`SELECT id, fullname, balance FROM accounts WHERE balance >= 100000;`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var id int
		var fullname string
		var balance float32

		err = rows.Scan(&id, &fullname, &balance)
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, model.Account{
			ID:       id,
			Fullname: fullname,
			Balance:  int(balance),
		})
	}

	return &accounts, nil
}

func (b *bankRepo) InsertTransaction(transaction *model.Transaction) (err error) {
	query := `INSERT INTO transactions (account_id, transaction_type, description, amount, ending_balance, transaction_date)
	VALUES ($1, $2, $3, $4, $5, $6);`

	_, err = b.db.Exec(query, transaction.AccountID, transaction.TransactionType, transaction.Description, transaction.Amount, transaction.EndingBalance, transaction.TransactionDate)
	if err != nil {
		panic(err)
	}

	return nil
}
