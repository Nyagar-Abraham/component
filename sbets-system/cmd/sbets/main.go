package main

import (
	"log"
	"net/http"

	"sbets-system/pkg/client"
	"sbets-system/pkg/database"
	"sbets-system/pkg/expense"
	"sbets-system/pkg/ui"
)

func main() {
	// Initialize database
	repo, err := database.NewRepository("sbets.db")
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}
	defer repo.Close()

	// Initialize converter client
	converterClient := client.NewConverterClient()

	// Initialize expense service
	expenseService := expense.NewService(repo, converterClient)

	// Setup routes
	router := ui.SetupRoutes(expenseService)

	// Start server
	log.Println("SBETS starting on port 8081...")
	log.Println("Make sure Currency Converter Service is running on port 8080")
	log.Fatal(http.ListenAndServe(":8081", router))
}