package service

import (
	"core/domain"
	"core/models"
)

type UserListService struct {
	UserListDomain domain.UserListDomain
}

func (u *UserListService) GetUserList(userId int64) ([]models.GetUserList, error) {
	userListParam := models.UserList{
		UserID: userId,
	}
	data, err := u.UserListDomain.GetUserList(userListParam)
	if err != nil {
		return nil, err
	}

	var result []models.GetUserList
	for _, item := range data {
		result = append(result, models.GetUserList{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return result, nil
}
