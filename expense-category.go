package harvest

import (
	"fmt"
	"time"
)

type ExpenseCategoriesResponse struct {
	ExpenseCategories []*ExpenseCategory `json:"expense_categories"`
	PerPage           int64              `json:"per_page"`
	TotalPages        int64              `json:"total_pages"`
	TotalEntries      int64              `json:"total_entries"`
	NextPage          *int64             `json:"next_page"`
	PreviousPage      *int64             `json:"previous_page"`
	Page              int64              `json:"page"`
}

type ExpenseCategory struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	UnitName    string    `json:"unit_name"`
	UnitPrice   float64   `json:"unit_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Deactivated bool      `json:"deactivated"`
}

func (a *API) GetExpenseCategory(expenseCategoryID int64, args Arguments) (expenseCategory *ExpenseCategory, err error) {
	expenseCategory = &ExpenseCategory{}
	path := fmt.Sprintf("/expense_categories/%v", expenseCategoryID)
	err = a.Get(path, args, expenseCategory)
	return expenseCategory, err
}

func (a *API) GetExpenseCategories(args Arguments) (expenseCategories []*ExpenseCategory, err error) {
	expenseCategoriesResponse := ExpenseCategoriesResponse{}
	path := fmt.Sprintf("/expense_categories")
	err = a.Get(path, args, &expenseCategoriesResponse)
	return expenseCategoriesResponse.ExpenseCategories, err
}
