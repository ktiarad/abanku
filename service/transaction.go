package service

import (
	"abanku/repository"
	"fmt"
)

type TransactionServices struct {
	Bank repository.BankRepo
}

func NewTransactionService(bank repository.BankRepo) *TransactionServices {
	return &TransactionServices{
		Bank: bank,
	}
}

func (t *TransactionServices) GetAllAccounts() (err error) {
	// var accounts *[]model.Account
	accounts, err := t.Bank.GetAllAccounts()
	if err != nil {
		panic(err)
	}

	for _, account := range *accounts {
		fmt.Println(account.Fullname)
	}

	return err
}
