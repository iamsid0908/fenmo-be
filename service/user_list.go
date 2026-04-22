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

func (u *UserListService) CreateUserList(param models.CreateUserListReqs) (models.CreateUserListResp, error) {
	userList := models.UserList{
		UserID:      param.UserId,
		Name:        param.Name,
		Description: param.Description,
	}
	data, err := u.UserListDomain.Create(userList)
	if err != nil {
		return models.CreateUserListResp{}, err
	}
	return models.CreateUserListResp{
		Name:        data.Name,
		Description: data.Description,
		UserId:      data.UserID,
		CreatedAt:   data.CreatedAt,
	}, nil
}

func (u *UserListService) GetUserListExpenses(userID int64, query models.UserListExpenseQuery) ([]models.UserListExpenseSummary, int64, error) {
	return u.UserListDomain.GetUserListExpenses(userID, query)
}
