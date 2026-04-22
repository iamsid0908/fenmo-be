package domain

import (
	"core/config"
	"core/models"
)

type CategoryDomain interface {
	CreateCategory(category models.Category) (models.Category, error)
	GetCategories(userID int64) ([]models.Category, error)
}

type CategoryDomainCtx struct{}

func (c *CategoryDomainCtx) CreateCategory(category models.Category) (models.Category, error) {
	result := config.DbManager().Create(&category)
	if result.Error != nil {
		return models.Category{}, result.Error
	}
	return category, nil
}

func (c *CategoryDomainCtx) GetCategories(userID int64) ([]models.Category, error) {
	var categories []models.Category
	result := config.DbManager().
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
