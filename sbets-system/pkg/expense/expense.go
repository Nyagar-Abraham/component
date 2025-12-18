package expense

import (
	"sbets-system/pkg/client"
	"sbets-system/pkg/database"
)

type Service struct {
	repo      *database.Repository
	converter *client.ConverterClient
}

func NewService(repo *database.Repository, converter *client.ConverterClient) *Service {
	return &Service{
		repo:      repo,
		converter: converter,
	}
}

func (s *Service) AddExpense(amount float64, currency, description string) error {
	// Convert to USD as base currency
	convertedAmount := amount
	if currency != "USD" {
		result, err := s.converter.Convert(amount, currency, "USD")
		if err != nil {
			return err
		}
		convertedAmount = result.ConvertedAmount
	}

	expense := &database.Expense{
		Amount:          amount,
		Currency:        currency,
		ConvertedAmount: convertedAmount,
		Description:     description,
	}

	return s.repo.AddExpense(expense)
}

func (s *Service) GetExpenses() ([]database.Expense, error) {
	return s.repo.GetExpenses()
}

func (s *Service) GetBudget() (*database.Budget, error) {
	total, err := s.repo.GetTotalExpenses()
	if err != nil {
		return nil, err
	}

	expenses, err := s.repo.GetExpenses()
	if err != nil {
		return nil, err
	}

	return &database.Budget{
		TotalExpenses: total,
		BaseCurrency:  "USD",
		ExpenseCount:  len(expenses),
	}, nil
}