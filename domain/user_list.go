package domain

import (
	"core/config"
	"core/models"
)

type UserListDomain interface {
	GetUserList(userListParam models.UserList) ([]models.UserList, error)
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
