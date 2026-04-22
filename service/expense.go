package service

import (
	"core/domain"
	"core/models"
	"time"
)

type ExpenseService struct {
	ExpenseDomain domain.ExpenseDomain
}

func (s *ExpenseService) CreateExpense(param models.CreateExpenseRequest) (models.ExpenseResponse, error) {
	date, err := time.Parse("2006-01-02", param.ExpenseDate)
	if err != nil {
		return models.ExpenseResponse{}, err
	}

	currency := param.Currency
	if currency == "" {
		currency = "INR"
	}

	expense := models.Expense{
		UserID:      param.UserID,
		UserListID:  param.UserListID,
		Amount:      param.Amount,
		CategoryID:  param.CategoryID,
		Currency:    currency,
		Description: param.Description,
		Date:        date,
	}

	data, err := s.ExpenseDomain.CreateExpense(expense)
	if err != nil {
		return models.ExpenseResponse{}, err
	}

	return models.ExpenseResponse{
		ID:          data.ID,
		Amount:      data.Amount,
		Currency:    data.Currency,
		CategoryID:  data.CategoryID,
		Description: data.Description,
		Date:        data.Date,
		CreatedAt:   data.CreatedAt,
	}, nil
}
