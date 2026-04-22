package models

import "time"

type Expense struct {
	ID             int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID         int64     `gorm:"column:user_id;not null;index"`
	Amount         float64   `gorm:"column:amount;not null"`
	CategoryID     int64     `gorm:"column:category_id;not null;index"`
	Category       Category  `gorm:"foreignKey:CategoryID;references:ID"`
	Description    string    `gorm:"column:description"`
	Date           time.Time `gorm:"column:date;not null"`
	IdempotencyKey string    `gorm:"column:idempotency_key;uniqueIndex;not null"`
	CreatedAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt      time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Expense) TableName() string {
	return "expenses"
}

type CreateExpenseRequest struct {
	// IdempotencyKey is a client-generated UUID so that retries on unreliable
	// networks are safe – duplicate submissions return the original record.
	IdempotencyKey string  `json:"idempotency_key" validate:"required"`
	Amount         float64 `json:"amount"          validate:"required,gt=0"`
	CategoryID     int64   `json:"category_id"     validate:"required,gt=0"`
	Description    string  `json:"description"`
	Date           string  `json:"date"            validate:"required"` // RFC3339 or YYYY-MM-DD
}

type ListExpenseQuery struct {
	// Filter by category_id (optional). Pass ?category_id=3 in the query string.
	CategoryID int64 `query:"category_id"`
}

type ExpenseResponse struct {
	ID           int64     `json:"id"`
	Amount       float64   `json:"amount"`
	CategoryID   int64     `json:"category_id"`
	CategoryName string    `json:"category"`
	Description  string    `json:"description"`
	Date         time.Time `json:"date"`
	CreatedAt    time.Time `json:"created_at"`
}

type ExpenseListResponse struct {
	Expenses []ExpenseResponse `json:"expenses"`
	Total    float64           `json:"total"`
}
