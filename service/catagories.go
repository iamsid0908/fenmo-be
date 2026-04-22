package service

import (
	"core/domain"
	"core/models"
)

type CategoryService struct {
	CategoryDomain domain.CategoryDomain
}

func (s *CategoryService) CreateCategory(userID int64, name string) (models.Category, error) {
	return s.CategoryDomain.CreateCategory(models.Category{
		UserId: userID,
		Name:   name,
	})
}

func (s *CategoryService) GetCategories(userID int64) ([]models.Category, error) {
	return s.CategoryDomain.GetCategories(userID)
}
