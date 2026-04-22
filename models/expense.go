package models

import "time"

type Expense struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int64     `gorm:"column:user_id;not null;index"`
	UserListID  int64     `gorm:"column:user_list_id;not null;index"`
	Amount      float64   `gorm:"column:amount;not null"`
	CategoryID  int64     `gorm:"column:category_id;not null;index"`
	Category    Category  `gorm:"foreignKey:CategoryID;references:ID"`
	Currency    string    `gorm:"column:currency;not null;default:INR"`
	Description string    `gorm:"column:description"`
	Date        time.Time `gorm:"column:date;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

func (Expense) TableName() string {
	return "expenses"
}

type CreateExpenseRequest struct {
	UserID      int64   `json:"-"`
	UserListID  int64   `json:"user_list_id"    validate:"required,gt=0"`
	Amount      float64 `json:"amount"          validate:"required,gt=0"`
	CategoryID  int64   `json:"category_id"     validate:"required,gt=0"`
	Currency    string  `json:"currency"`
	Description string  `json:"description"`
	ExpenseDate string  `json:"expense_date"    validate:"required"` // YYYY-MM-DD
}

type ListExpenseQuery struct {
	// Filter by category_id (optional). Pass ?category_id=3 in the query string.
	CategoryID int64 `query:"category_id"`
}

type ExpenseResponse struct {
	ID           int64     `json:"id"`
	Amount       float64   `json:"amount"`
	Currency     string    `json:"currency"`
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
