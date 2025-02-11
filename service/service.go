package service

import (
	"abanku/repository"
	"abanku/service/accounts"
	"database/sql"
)

type Services struct {
	accounts.Account
}

func NewService(
	db *sql.DB,
) Services {
	repo := repository.NewBankRepo(db)

	return Services{
		accounts.NewAccountService(repo),
	}
}
