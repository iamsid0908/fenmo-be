package domain

import (
	"core/config"
	"core/models"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// ExpenseDomain defines the data-access contract for expenses.
type ExpenseDomain interface {
	Create(expense models.Expense) (models.Expense, error)
	List(userID int64, category string) ([]models.Expense, error)
}

type ExpenseDomainCtx struct{}

// Create inserts a new expense or, when the idempotency_key already exists for
// this user, returns the existing record without error (safe retry behaviour).
func (d *ExpenseDomainCtx) Create(expense models.Expense) (models.Expense, error) {
	db := config.DbManager()

	// ON CONFLICT on idempotency_key → do nothing, then return the row.
	result := db.
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "idempotency_key"}},
			DoNothing: true,
		}).
		Create(&expense)

	if result.Error != nil {
		return models.Expense{}, result.Error
	}

	// If no rows were inserted (conflict), fetch the existing record.
	if result.RowsAffected == 0 {
		existing := models.Expense{}
		if err := db.
			Where("idempotency_key = ? AND user_id = ?", expense.IdempotencyKey, expense.UserID).
			First(&existing).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return models.Expense{}, errors.New("expense not found after conflict resolution")
			}
			return models.Expense{}, err
		}
		return existing, nil
	}

	return expense, nil
}

// List returns expenses for a user, optionally filtered by category, always
// sorted newest-first.
func (d *ExpenseDomainCtx) List(userID int64, category string) ([]models.Expense, error) {
	db := config.DbManager()
	var expenses []models.Expense

	query := db.Where("user_id = ?", userID)
	if category != "" {
		query = query.Where("category = ?", category)
	}

	if err := query.Order("date DESC, created_at DESC").Find(&expenses).Error; err != nil {
		return nil, err
	}
	return expenses, nil
}
