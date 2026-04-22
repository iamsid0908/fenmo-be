package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type UserListHandler struct {
	UserListService service.UserListService
}

func (userListHandler *UserListHandler) GetUserList(c echo.Context) error {
	userId := c.Get("id").(int64)
	data, err := userListHandler.UserListService.GetUserList(userId)
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (userListHandler *UserListHandler) CreateUserList(c echo.Context) error {
	userId := c.Get("id").(int64)
	param := models.CreateUserListReqs{}
	err := c.Bind(&param)
	if err != nil {
		return err
	}
	param.UserId = userId
	data, err := userListHandler.UserListService.CreateUserList(param)
	if err != nil {
		return err
	}
	resp := models.BasicResp{
		Message: utils.Success,
		Data:    data,
	}
	return c.JSON(http.StatusOK, resp)
}

func (userListHandler *UserListHandler) GetUserListExpenses(c echo.Context) error {
	userId := c.Get("id").(int64)

	var query models.UserListExpenseQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, total, err := userListHandler.UserListService.GetUserListExpenses(userId, query)
	if err != nil {
		return err
	}

	limit := query.Limit
	if limit <= 0 {
		limit = 10
	}
	page := query.Page
	if page <= 0 {
		page = 1
	}
	totalPages := (total + int64(limit) - 1) / int64(limit)

	return c.JSON(http.StatusOK, models.BasicRespWithMeta{
		Message: utils.Success,
		Data:    data,
		Meta: models.MetaPagination{
			PageNumber:   int64(page),
			PageSize:     int64(limit),
			TotalPages:   totalPages,
			TotalRecords: total,
		},
	})
}
