package converter

import "time"

// ICurrencyConverter defines the core currency conversion interface
type ICurrencyConverter interface {
	Convert(amount float64, from, to string) (float64, error)
	SetExchangeRate(from, to string, rate float64) error
	GetSupportedCurrencies() []string
	ResetRates()
}

// IExchangeRateProvider defines the exchange rate management interface
type IExchangeRateProvider interface {
	GetRate(from, to string) (float64, error)
	SetRate(from, to string, rate float64) error
}

// ConversionResult represents the result of a currency conversion
type ConversionResult struct {
	ConvertedAmount float64   `json:"convertedAmount"`
	OriginalAmount  float64   `json:"originalAmount"`
	FromCurrency    string    `json:"fromCurrency"`
	ToCurrency      string    `json:"toCurrency"`
	ExchangeRate    float64   `json:"exchangeRate"`
	Timestamp       time.Time `json:"timestamp"`
}

// ConversionError represents conversion-specific errors
type ConversionError struct {
	Code    string  `json:"code"`
	Message string  `json:"message"`
	Amount  float64 `json:"amount,omitempty"`
	From    string  `json:"from,omitempty"`
	To      string  `json:"to,omitempty"`
}

func (e ConversionError) Error() string {
	return e.Message
}