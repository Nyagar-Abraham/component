package api

import (
	"github.com/gorilla/mux"
	"currency-converter-service/pkg/converter"
)

// SetupRoutes configures all API routes
func SetupRoutes(conv converter.ICurrencyConverter) *mux.Router {
	handler := NewHandler(conv)
	
	r := mux.NewRouter()
	
	// Apply middleware
	r.Use(LoggingMiddleware)
	r.Use(CORSMiddleware)
	
	// API routes
	r.HandleFunc("/convert", handler.ConvertHandler).Methods("POST")
	r.HandleFunc("/currencies", handler.CurrenciesHandler).Methods("GET")
	r.HandleFunc("/rates", handler.SetRateHandler).Methods("POST")
	
	return r
}