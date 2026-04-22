package validation

import (
	"core/models"
	"core/utils"
)

func CreateExpense(param *models.CreateExpenseRequest) error {
	if param.Amount <= 0 {
		return utils.ErrInvalidAmount
	}
	if param.CategoryID <= 0 {
		return utils.ErrEmptyCategoryID
	}
	if param.UserListID <= 0 {
		return utils.ErrEmptyUserListID
	}
	if param.ExpenseDate == "" {
		return utils.ErrEmptyExpenseDate
	}
	return nil
}
