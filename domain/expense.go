package domain

import (
	"core/config"
	"core/models"
)

type ExpenseDomain interface {
	CreateExpense(expense models.Expense) (models.Expense, error)
	RecentExpenses(userID int64, query models.ListExpenseQuery) ([]models.Expense, int64, error)
}

type ExpenseDomainCtx struct{}

func (e *ExpenseDomainCtx) CreateExpense(expense models.Expense) (models.Expense, error) {
	result := config.DbManager().Create(&expense)
	if result.Error != nil {
		return models.Expense{}, result.Error
	}
	return expense, nil
}

func (e *ExpenseDomainCtx) RecentExpenses(userID int64, query models.ListExpenseQuery) ([]models.Expense, int64, error) {
	var expenses []models.Expense
	var total int64

	db := config.DbManager().Model(&models.Expense{}).
		Preload("Category").
		Preload("UserList").
		Where("user_id = ?", userID)

	if query.CategoryID > 0 {
		db = db.Where("category_id = ?", query.CategoryID)
	}
	if query.UserListID > 0 {
		db = db.Where("user_list_id = ?", query.UserListID)
	}

	db.Count(&total)

	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}
	page := query.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	result := db.Order("date DESC, created_at DESC").Limit(limit).Offset(offset).Find(&expenses)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return expenses, total, nil
}
