package harvest

import (
	"fmt"
	"time"
)

type ExpenseCategoryResponse struct {
	ExpenseCategory *ExpenseCategory `json:"expense_category"`
}

type ExpenseCategory struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	UnitName    string    `json:"unit_name"`
	UnitPrice   float64   `json:"unit_price"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Deactivated bool      `json:"deactivated"`
}

func (a *API) GetExpenseCategory(expensecategoryID int64, args Arguments) (expensecategory *ExpenseCategory, err error) {
	expenseCategoryResponse := ExpenseCategoryResponse{}
	path := fmt.Sprintf("/expense_categories/%v", expensecategoryID)
	err = a.Get(path, args, &expenseCategoryResponse)
	return expenseCategoryResponse.ExpenseCategory, err
}

func (a *API) GetExpenseCategories(args Arguments) (expensecategories []*ExpenseCategory, err error) {
	expenseCategoriesResponse := make([]*ExpenseCategoryResponse, 0)
	path := fmt.Sprintf("/expense_categories")
	err = a.Get(path, args, &expenseCategoriesResponse)
	for _, er := range expenseCategoriesResponse {
		expensecategories = append(expensecategories, er.ExpenseCategory)
	}
	return expensecategories, err
}
