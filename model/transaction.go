package model

import "time"

type Transaction struct {
	ID              int
	AccountID       int
	TransactionType string
	Description     string
	Amount          float32
	EndingBalance   float32
	TransactionDate *time.Time
}
