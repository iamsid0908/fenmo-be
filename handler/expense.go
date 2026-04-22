package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type ExpenseHandler struct {
	ExpenseService service.ExpenseService
}

func (h *ExpenseHandler) CreateExpense(c echo.Context) error {
	userId := c.Get("id").(int64)
	param := models.CreateExpenseRequest{}
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	param.UserID = userId

	data, err := h.ExpenseService.CreateExpense(param)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, models.BasicResp{
		Message: utils.Success,
		Data:    data,
	})
}
