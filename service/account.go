package service

import (
	"abanku/model"
	"abanku/repository"
	"log"
	"math/rand"
)

type AccountServices struct {
	Bank repository.BankRepo
}

func NewAccountService(bank repository.BankRepo) *AccountServices {
	return &AccountServices{
		Bank: bank,
	}
}

func (a *AccountServices) CreateAccount() error {
	var balance int
	var err error

	for i := 0; i < 5; i++ {
		go func() {
			balance = rand.Intn(9999) * 10_000

			err = a.Bank.CreateAccount(&model.Account{Balance: float32(balance)})
			if err != nil {
				log.Fatalf("Error when insert account: %v", err)
			}
		}()
	}

	return nil
}
