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
