package database

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Repository struct {
	db *sql.DB
}

// Constructor to initialize database connection and create tables thereafter
func NewRepository(dbPath string) (*Repository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	repo := &Repository{db: db}
	if err := repo.createTables(); err != nil {
		return nil, err
	}

	return repo, nil
}

func (r *Repository) createTables() error {
	query := `
	CREATE TABLE IF NOT EXISTS expenses (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		amount DECIMAL(10,2) NOT NULL,
		currency VARCHAR(3) NOT NULL,
		converted_amount DECIMAL(10,2) NOT NULL,
		description TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err := r.db.Exec(query)
	return err
}

func (r *Repository) AddExpense(expense *Expense) error {
	query := `INSERT INTO expenses (amount, currency, converted_amount, description, created_at) 
			  VALUES (?, ?, ?, ?, ?)`

	_, err := r.db.Exec(query, expense.Amount, expense.Currency, expense.ConvertedAmount,
		expense.Description, time.Now())
	return err
}

func (r *Repository) GetExpenses() ([]Expense, error) {
	query := `SELECT id, amount, currency, converted_amount, description, created_at 
			  FROM expenses ORDER BY created_at DESC`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var expenses []Expense
	for rows.Next() {
		var expense Expense

		//What Scan does:
		//
		//Reads the current row
		//
		//Takes column values left → right
		//
		//Copies them into the provided variables
		err := rows.Scan(&expense.ID, &expense.Amount, &expense.Currency,
			&expense.ConvertedAmount, &expense.Description, &expense.CreatedAt)
		if err != nil {
			return nil, err
		}
		expenses = append(expenses, expense)
	}

	return expenses, nil
}

func (r *Repository) GetTotalExpenses() (float64, error) {
	// COALESCE(value, fallback)
	query := `SELECT COALESCE(SUM(converted_amount), 0) FROM expenses`

	var total float64
	// QueryRow is used when you expect exactly ONE row of result from the database.
	// Even if the table is empty, this query still returns one row because of SUM
	err := r.db.QueryRow(query).Scan(&total)
	return total, err
}

func (r *Repository) DeleteExpense(id int) error {
	query := `DELETE FROM expenses WHERE id = ?`
	_, err := r.db.Exec(query, id)
	return err
}

func (r *Repository) Close() error {
	return r.db.Close()
}
