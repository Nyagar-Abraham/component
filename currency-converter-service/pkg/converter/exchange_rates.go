package converter

// DefaultExchangeRates contains hardcoded exchange rates
var DefaultExchangeRates = map[string]map[string]float64{
	"USD": {
		"EUR": 0.85,
		"GBP": 0.73,
		"JPY": 110.0,
		"CAD": 1.25,
		"AUD": 1.35,
	},
	"EUR": {
		"USD": 1.18,
		"GBP": 0.86,
		"JPY": 129.5,
		"CAD": 1.47,
		"AUD": 1.59,
	},
	"GBP": {
		"USD": 1.37,
		"EUR": 1.16,
		"JPY": 150.5,
		"CAD": 1.71,
		"AUD": 1.85,
	},
	"JPY": {
		"USD": 0.0091,
		"EUR": 0.0077,
		"GBP": 0.0066,
		"CAD": 0.0114,
		"AUD": 0.0123,
	},
	"CAD": {
		"USD": 0.80,
		"EUR": 0.68,
		"GBP": 0.58,
		"JPY": 88.0,
		"AUD": 1.08,
	},
	"AUD": {
		"USD": 0.74,
		"EUR": 0.63,
		"GBP": 0.54,
		"JPY": 81.3,
		"CAD": 0.93,
	},
}

// SupportedCurrencies returns the list of supported currency codes
func SupportedCurrencies() []string {
	return []string{"USD", "EUR", "GBP", "JPY", "CAD", "AUD"}
}