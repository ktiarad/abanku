package main

import (
	"abanku/configs"
	"abanku/model"
	"abanku/repository"
	"abanku/service"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Start ABanku")

	db := configs.ConnectDB()

	defer db.Close()

	bankRepo := repository.NewBankRepo(db)

	// run this first before run the transactionServices
	accountServices := service.NewAccountService(bankRepo)
	_ = accountServices.CreateAccount()
	// finish run this

	transactionServices := service.NewTransactionService(bankRepo)

	err := transactionServices.GetAllAccounts()
	if err != nil {
		fmt.Println("error transactionServices.GetAllAccount()")
	}

	now := time.Now()
	payload := model.Transaction{
		AccountID:       1,
		TransactionType: "Debit",
		Description:     "Admin Fee",
		Amount:          -15000,
		EndingBalance:   1000000,
		TransactionDate: &now,
	}

	err = bankRepo.InsertTransaction(&payload)
	if err != nil {
		fmt.Println("error bankRepo.InsertTransaction()")
	}

	fmt.Println("Finish ABanku")
	// go run main.go

	// get all users (1mio) with money > 50k

	// add admin fee, interest, and interest tax, then insert to database (total 3 transaction per users)

	// check that all users already get three transactions respectively
}
