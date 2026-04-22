package handler

import (
	"core/models"
	"core/service"
	"core/utils"
	"net/http"

	"github.com/labstack/echo"
)

type CategoryHandler struct {
	CategoryService service.CategoryService
}

type createCategoryReq struct {
	Name string `json:"name" validate:"required"`
}

func (h *CategoryHandler) CreateCategory(c echo.Context) error {
	userID := c.Get("id").(int64)
	var req createCategoryReq
	if err := c.Bind(&req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	data, err := h.CategoryService.CreateCategory(userID, req.Name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, models.BasicResp{
		Message: utils.Success,
		Data:    data,
	})
}

func (h *CategoryHandler) GetCategories(c echo.Context) error {
	userID := c.Get("id").(int64)
	data, err := h.CategoryService.GetCategories(userID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, models.BasicResp{
		Message: utils.Success,
		Data:    data,
	})
}
