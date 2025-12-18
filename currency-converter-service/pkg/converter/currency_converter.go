package converter

import (
	"fmt"
	"strings"
	"time"
)

// CurrencyConverter implements the ICurrencyConverter interface
type CurrencyConverter struct {
	exchangeRates map[string]map[string]float64
}

// NewCurrencyConverter creates a new currency converter instance
func NewCurrencyConverter() *CurrencyConverter {
	// Deep copy of default rates
	rates := make(map[string]map[string]float64)

	// loop through the nested map and extract K(from) and V(toRates which is  also a map)
	for from, toRates := range DefaultExchangeRates {
		rates[from] = make(map[string]float64)
		//perform an inner loop to fill all the rate for a given currency
		for to, rate := range toRates {
			rates[from][to] = rate
		}
	}

	return &CurrencyConverter{
		exchangeRates: rates,
	}
}

// Convert converts an amount from one currency to another
func (c *CurrencyConverter) Convert(amount float64, from, to string) (float64, error) {
	if amount <= 0 {
		return 0, ConversionError{
			Code:    "INVALID_AMOUNT",
			Message: "amount must be greater than zero",
			Amount:  amount,
		}
	}

	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	// handle same currency
	if from == to {
		return amount, nil
	}

	rate, err := c.GetRate(from, to)
	if err != nil {
		return 0, err
	}

	return amount * rate, nil
}

// GetRate retrieves the exchange rate between two currencies
func (c *CurrencyConverter) GetRate(from, to string) (float64, error) {
	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	if fromRates, exists := c.exchangeRates[from]; exists {
		if rate, exists := fromRates[to]; exists {
			return rate, nil
		}
	}

	return 0, ConversionError{
		Code:    "RATE_NOT_FOUND",
		Message: fmt.Sprintf("exchange rate not found for %s to %s", from, to),
		From:    from,
		To:      to,
	}
}

// SetExchangeRate sets a custom exchange rate
func (c *CurrencyConverter) SetExchangeRate(from, to string, rate float64) error {
	if rate <= 0 {
		return ConversionError{
			Code:    "INVALID_RATE",
			Message: "exchange rate must be greater than zero",
		}
	}

	from = strings.ToUpper(from)
	to = strings.ToUpper(to)

	if c.exchangeRates[from] == nil {
		c.exchangeRates[from] = make(map[string]float64)
	}
	c.exchangeRates[from][to] = rate

	return nil
}

// SetRate implements IExchangeRateProvider interface
func (c *CurrencyConverter) SetRate(from, to string, rate float64) error {
	return c.SetExchangeRate(from, to, rate)
}

// GetSupportedCurrencies returns the list of supported currencies
func (c *CurrencyConverter) GetSupportedCurrencies() []string {
	return SupportedCurrencies()
}

// ResetRates resets all exchange rates to default values
func (c *CurrencyConverter) ResetRates() {
	c.exchangeRates = make(map[string]map[string]float64)
	for from, toRates := range DefaultExchangeRates {
		c.exchangeRates[from] = make(map[string]float64)
		for to, rate := range toRates {
			c.exchangeRates[from][to] = rate
		}
	}
}

// ConvertWithResult returns a detailed conversion result
func (c *CurrencyConverter) ConvertWithResult(amount float64, from, to string) (*ConversionResult, error) {
	convertedAmount, err := c.Convert(amount, from, to)
	if err != nil {
		return nil, err
	}

	rate, _ := c.GetRate(from, to)

	return &ConversionResult{
		ConvertedAmount: convertedAmount,
		OriginalAmount:  amount,
		FromCurrency:    strings.ToUpper(from),
		ToCurrency:      strings.ToUpper(to),
		ExchangeRate:    rate,
		Timestamp:       time.Now(),
	}, nil
}
