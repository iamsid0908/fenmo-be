package models

import "time"

type Category struct {
	ID        int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserId    int64     `gorm:"column:user_id;not null;uniqueIndex:idx_user_category"`
	Name      string    `gorm:"column:name;not null;uniqueIndex:idx_user_category"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Category) TableName() string {
	return "categories"
}
