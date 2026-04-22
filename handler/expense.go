package handler

import (
	"core/handler/validation"
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

	if err := validation.CreateExpense(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, err := h.ExpenseService.CreateExpense(param)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, models.BasicResp{
		Message: utils.Success,
		Data:    data,
	})
}

func (h *ExpenseHandler) RecentExpenses(c echo.Context) error {
	userId := c.Get("id").(int64)

	var query models.ListExpenseQuery
	if err := c.Bind(&query); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	data, total, err := h.ExpenseService.RecentExpenses(userId, query)
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
