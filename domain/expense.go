package domain

import (
	"core/config"
	"core/models"
)

type ExpenseDomain interface {
	CreateExpense(expense models.Expense) (models.Expense, error)
}

type ExpenseDomainCtx struct{}

func (e *ExpenseDomainCtx) CreateExpense(expense models.Expense) (models.Expense, error) {
	result := config.DbManager().Create(&expense)
	if result.Error != nil {
		return models.Expense{}, result.Error
	}
	return expense, nil
}
