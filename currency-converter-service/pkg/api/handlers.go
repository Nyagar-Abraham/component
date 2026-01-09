package api

import (
	"encoding/json"
	"net/http"

	"currency-converter-service/pkg/converter"
	"currency-converter-service/pkg/models"
)

type Handler struct {
	converter converter.ICurrencyConverter
}

// constructor
func NewHandler(conv converter.ICurrencyConverter) *Handler {
	return &Handler{converter: conv}
}

// ConvertHandler handles POST /convert
func (h *Handler) ConvertHandler(w http.ResponseWriter, r *http.Request) {
	var req models.ConvertRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON format")
		return
	}

	if req.Amount <= 0 {
		h.writeError(w, http.StatusBadRequest, "INVALID_AMOUNT", "Amount must be greater than zero")
		return
	}

	result, err := h.converter.(*converter.CurrencyConverter).ConvertWithResult(req.Amount, req.From, req.To)
	if err != nil {
		if convErr, ok := err.(converter.ConversionError); ok {
			h.writeError(w, http.StatusBadRequest, convErr.Code, convErr.Message)
		} else {
			h.writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		}
		return
	}

	response := models.ConvertResponse{
		ConvertedAmount: result.ConvertedAmount,
		OriginalAmount:  result.OriginalAmount,
		FromCurrency:    result.FromCurrency,
		ToCurrency:      result.ToCurrency,
		ExchangeRate:    result.ExchangeRate,
		Timestamp:       result.Timestamp,
	}

	h.writeJSON(w, http.StatusOK, response)
}

// CurrenciesHandler handles GET /currencies
func (h *Handler) CurrenciesHandler(w http.ResponseWriter, r *http.Request) {
	currencies := h.converter.GetSupportedCurrencies()
	response := models.CurrenciesResponse{Currencies: currencies}
	h.writeJSON(w, http.StatusOK, response)
}

// SetRateHandler handles POST /rates
func (h *Handler) SetRateHandler(w http.ResponseWriter, r *http.Request) {
	var req models.SetRateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.writeError(w, http.StatusBadRequest, "INVALID_JSON", "Invalid JSON format")
		return
	}

	if err := h.converter.SetExchangeRate(req.From, req.To, req.Rate); err != nil {
		if convErr, ok := err.(converter.ConversionError); ok {
			h.writeError(w, http.StatusBadRequest, convErr.Code, convErr.Message)
		} else {
			h.writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", err.Error())
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "success"})
}

// ResetRatesHandler handles DELETE /rates
func (h *Handler) ResetRatesHandler(w http.ResponseWriter, r *http.Request) {
	h.converter.ResetRates()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "rates reset"})
}

func (h *Handler) writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func (h *Handler) writeError(w http.ResponseWriter, status int, code, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(models.ErrorResponse{
		Error: message,
		Code:  code,
	})
}
