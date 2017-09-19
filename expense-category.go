package harvest

import (
	"fmt"
	"time"
)

type ExpenseCategoriesResponse struct {
	PagedResponse
	ExpenseCategories []*ExpenseCategory `json:"expense_categories"`
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
	expenseCategories = make([]*ExpenseCategory, 0)
	expenseCategoriesResponse := ExpenseCategoriesResponse{}
	err = a.GetPaginated("/expense_categories", args, &expenseCategoriesResponse, func() {
		for _, ec := range expenseCategoriesResponse.ExpenseCategories {
			expenseCategories = append(expenseCategories, ec)
		}
		expenseCategoriesResponse.ExpenseCategories = make([]*ExpenseCategory, 0)
	})
	return expenseCategories, err
}
