package main

import (
	"log"
	"net/http"

	"currency-converter-service/pkg/api"
	"currency-converter-service/pkg/converter"
)

func main() {
	// Create currency converter instance
	conv := converter.NewCurrencyConverter()

	// Setup routes
	router := api.SetupRoutes(conv)

	// Start server
	log.Println("Currency Converter Service starting on port 8085...")
	log.Fatal(http.ListenAndServe(":8085", router))
}
