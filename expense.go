package harvest

import (
	"fmt"
	"time"
)

type ExpenseResponse struct {
	Expense *Expense `json:"expense"`
}

type Expense struct {
	ID                int64     `json:"id"`
	Notes             string    `json:"notes"`
	TotalCost         float64   `json:"total_cost"`
	Units             float64   `json:"units"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ProjectID         int64     `json:"project_id"`
	ExpenseCategoryID int64     `json:"expense_category_id"`
	UserID            int64     `json:"user_id"`
	SpentAt           string    `json:"spent_at"`
	IsClosed          bool      `json:"is_closed"`
	InvoiceID         int64     `json:"invoice_id"`
	Billable          bool      `json:"billable"`
	CompanyID         int64     `json:"company_id"`
	HasReceipt        bool      `json:"has_receipt"`
	ReceiptURL        string    `json:"receipt_url"`
	IsLocked          bool      `json:"is_locked"`
	LockedReason      string    `json:"locked_reason"`
}

func (a *API) GetExpense(expenseID int64, args Arguments) (expense *Expense, err error) {
	expenseResponse := ExpenseResponse{}
	path := fmt.Sprintf("/expenses/%v", expenseID)
	err = a.Get(path, args, &expenseResponse)
	return expenseResponse.Expense, err
}

func (a *API) GetExpenses(args Arguments) (expenses []*Expense, err error) {
	expensesResponse := make([]*ExpenseResponse, 0)
	path := fmt.Sprintf("/expenses")
	err = a.Get(path, args, &expensesResponse)
	for _, er := range expensesResponse {
		expenses = append(expenses, er.Expense)
	}
	return expenses, err
}
