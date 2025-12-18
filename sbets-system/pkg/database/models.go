package database

import "time"

type Expense struct {
	ID              int       `json:"id" db:"id"`
	Amount          float64   `json:"amount" db:"amount"`
	Currency        string    `json:"currency" db:"currency"`
	ConvertedAmount float64   `json:"convertedAmount" db:"converted_amount"`
	Description     string    `json:"description" db:"description"`
	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
}

type Budget struct {
	TotalExpenses   float64 `json:"totalExpenses"`
	BaseCurrency    string  `json:"baseCurrency"`
	ExpenseCount    int     `json:"expenseCount"`
}