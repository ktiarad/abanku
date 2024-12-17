package model

import "time"

type Transaction struct {
	ID              int
	AccountID       int
	TransactionType string
	Description     string
	Amount          float64
	EndingBalance   float64
	TransactionDate *time.Time
}
