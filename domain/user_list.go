package domain

import (
	"core/config"
	"core/models"
)

type UserListDomain interface {
	GetUserList(userListParam models.UserList) ([]models.UserList, error)
	Create(userList models.UserList) (models.UserList, error)
	GetUserListExpenses(userID int64, query models.UserListExpenseQuery) ([]models.UserListExpenseSummary, int64, error)
}

type UserListDomainCtx struct {
}

func (u *UserListDomainCtx) GetUserList(userListParam models.UserList) ([]models.UserList, error) {
	var userLists []models.UserList
	result := config.DbManager().
		Where("user_id = ?", userListParam.UserID).
		Order("created_at DESC").
		Limit(10).
		Find(&userLists)
	if result.Error != nil {
		return nil, result.Error
	}
	return userLists, nil
}

func (u *UserListDomainCtx) Create(userList models.UserList) (models.UserList, error) {
	result := config.DbManager().Create(&userList)
	if result.Error != nil {
		return models.UserList{}, result.Error
	}
	return userList, nil
}

func (u *UserListDomainCtx) GetUserListExpenses(userID int64, query models.UserListExpenseQuery) ([]models.UserListExpenseSummary, int64, error) {
	var total int64
	config.DbManager().Model(&models.UserList{}).Where("user_id = ?", userID).Count(&total)

	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}
	page := query.Page
	if page <= 0 {
		page = 1
	}
	offset := (page - 1) * limit

	var results []models.UserListExpenseSummary
	err := config.DbManager().
		Model(&models.UserList{}).
		Select("user_lists.id, user_lists.name, user_lists.description, COALESCE(SUM(expenses.amount), 0) AS total_expense").
		Joins("LEFT JOIN expenses ON expenses.user_list_id = user_lists.id").
		Where("user_lists.user_id = ?", userID).
		Group("user_lists.id, user_lists.name, user_lists.description").
		Order("user_lists.created_at DESC").
		Limit(limit).Offset(offset).
		Scan(&results).Error
	if err != nil {
		return nil, 0, err
	}
	return results, total, nil
}
