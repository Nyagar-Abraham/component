package models

// ConvertRequest represents a currency conversion request
type ConvertRequest struct {
	Amount float64 `json:"amount" validate:"required,gt=0"`
	From   string  `json:"from" validate:"required,len=3"`
	To     string  `json:"to" validate:"required,len=3"`
}

// SetRateRequest represents a request to set an exchange rate
type SetRateRequest struct {
	From string  `json:"from" validate:"required,len=3"`
	To   string  `json:"to" validate:"required,len=3"`
	Rate float64 `json:"rate" validate:"required,gt=0"`
}