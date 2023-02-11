package service

import (
	"abanku/configs"
	"abanku/model"
	"abanku/repository"
	"fmt"
	"log"
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

func (t *TransactionServices) MonthlyTransaction() (err error) {
	var accounts *[]model.Account

	accounts, err = t.Bank.GetAllAccounts()
	if err != nil {
		log.Fatalf("error when GetAllAccounts: %s", err)
		panic(err)
	}

	accountChannel := make(chan model.Account)

	defer close(accountChannel)

	<-accountChannel

	// go t.AdminFee(accounts)
	// go t.Interest(accounts)
	// go t.TaxInterest(accounts)

	for _, account := range *accounts {
		err = t.AdminFee(&account)

		if err != nil {
			log.Fatalf("error when AdminFee userID: %d", account.ID)
		}
	}

	for _, account := range *accounts {
		err = t.Interest(&account)

		if err != nil {
			log.Fatalf("error when Interest userID: %d", account.ID)
		}
	}

	for _, account := range *accounts {
		err = t.TaxInterest(&account)

		if err != nil {
			log.Fatalf("error when TaxInterest userID: %d", account.ID)
		}
	}

	return nil
}

func (t *TransactionServices) AdminFee(account *model.Account) (err error) {
	if t.isBelowMinimumBalance(account) {
		return nil
	}

	account.Balance -= configs.ADMIN_FEE

	return nil
}

func (t *TransactionServices) Interest(account *model.Account) (err error) {
	if t.isBelowMinimumBalance(account) {
		return nil
	}

	account.Balance = account.Balance + (account.Balance * configs.INTEREST_PERCENTAGE)

	return nil
}

func (t *TransactionServices) TaxInterest(account *model.Account) (err error) {
	if t.isBelowMinimumBalance(account) {
		return nil
	}

	account.Balance = account.Balance - (account.Balance * configs.TAX_INTEREST_PERCENTAGE)

	return nil
}

func (t *TransactionServices) isBelowMinimumBalance(account *model.Account) bool {
	return account.Balance < configs.MINIMUM_BALANCE
}
