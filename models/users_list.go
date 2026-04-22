package models

import "time"

// UserList represents a named collection created by a user (e.g. "Trip to Bali", "Groceries").
type UserList struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement"`
	UserID      int64     `gorm:"column:user_id;not null;index"`
	Name        string    `gorm:"column:name;not null"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoUpdateTime"`
	Expenses    []Expense `gorm:"many2many:user_list_expenses;foreignKey:ID;joinForeignKey:ListID;References:ID;joinReferences:ExpenseID"`
}

func (UserList) TableName() string {
	return "user_lists"
}

type GetUserList struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type CreateUserListReqs struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	UserId      int64  `json:"-"`
}

type CreateUserListResp struct {
	Name        string    `json:"name" validate:"required"`
	Description string    `json:"description"`
	UserId      int64     `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
}

type UserListExpenseSummary struct {
	ID           int64   `json:"id"`
	Name         string  `json:"name"`
	Description  string  `json:"description"`
	TotalExpense float64 `json:"total_expense"`
}

type UserListExpenseQuery struct {
	Page  int `query:"page"`
	Limit int `query:"limit"`
}
