package models

import "time"

// ConvertResponse represents a currency conversion response
type ConvertResponse struct {
	ConvertedAmount float64   `json:"convertedAmount"`
	OriginalAmount  float64   `json:"originalAmount"`
	FromCurrency    string    `json:"fromCurrency"`
	ToCurrency      string    `json:"toCurrency"`
	ExchangeRate    float64   `json:"exchangeRate"`
	Timestamp       time.Time `json:"timestamp"`
}

// CurrenciesResponse represents the supported currencies response
type CurrenciesResponse struct {
	Currencies []string `json:"currencies"`
}

// ErrorResponse represents an error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Code    string `json:"code,omitempty"`
	Details string `json:"details,omitempty"`
}