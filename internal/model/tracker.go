package model

import "time"

type ExpenseTracker struct {
	ID int `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Description string `json:"description"`
	Amount int `json:"amount"`
}
