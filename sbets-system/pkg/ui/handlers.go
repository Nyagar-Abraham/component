package ui

import (
	"encoding/json"
	"html/template"
	"net/http"
	"path/filepath"

	"sbets-system/pkg/expense"

	"github.com/gorilla/mux"
)

type Handler struct {
	expenseService *expense.Service
	templates      *template.Template
}

func NewHandler(expenseService *expense.Service) *Handler {
	templates := template.Must(template.ParseGlob(filepath.Join("web", "templates", "*.html")))
	return &Handler{
		expenseService: expenseService,
		templates:      templates,
	}
}

type AddExpenseRequest struct {
	Amount      float64 `json:"amount"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
}

func (h *Handler) AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var req AddExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.expenseService.AddExpense(req.Amount, req.Currency, req.Description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"status": "expense added"})
}

func (h *Handler) GetExpensesHandler(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.expenseService.GetExpenses()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}

func (h *Handler) GetBudgetHandler(w http.ResponseWriter, r *http.Request) {
	budget, err := h.expenseService.GetBudget()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(budget)
}

func (h *Handler) HomeHandler(w http.ResponseWriter, r *http.Request) {
	if err := h.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SetupRoutes(expenseService *expense.Service) *mux.Router {
	handler := NewHandler(expenseService)

	r := mux.NewRouter()

	// Static files
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))

	// Web UI
	r.HandleFunc("/", handler.HomeHandler).Methods("GET")

	// API routes
	r.HandleFunc("/api/expenses", handler.AddExpenseHandler).Methods("POST")
	r.HandleFunc("/api/expenses", handler.GetExpensesHandler).Methods("GET")
	r.HandleFunc("/api/budget", handler.GetBudgetHandler).Methods("GET")

	return r
}
