package harvest

import "testing"

func testExpense(t *testing.T) *Expense {
	a := testAPI()
	expenseResponse := mockResponse("expenses", "expense-example.json")
	a.BaseURL = expenseResponse.URL
	expense, err := a.GetExpense(7631396, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if expense.ExpenseCategoryID != 1338061 {
		t.Errorf("Incorrect expense category id '%v'", expense.ExpenseCategoryID)
	}
	if expense.Notes != "Your Updated Expense" {
		t.Errorf("Incorrect Expense Notes '%s'", expense.Notes)
	}
	return expense
}

func TestGetExpense(t *testing.T) {
	expense := testExpense(t)
	if expense == nil {
		t.Fatal("testExpense() returned nil instead of expense")
	}
	if expense.ID != 7631396 {
		t.Errorf("Incorrect expense ID '%v'", expense.ID)
	}
}

func TestGetExpenses(t *testing.T) {
	a := testAPI()
	expenseResponse := mockResponse("expenses", "expenses-example.json")
	a.BaseURL = expenseResponse.URL
	expenses, err := a.GetExpenses(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(expenses) != 2 {
		t.Errorf("Incorrect number of expenses '%v'", len(expenses))
	}
	if expenses[0].ID != 7631396 {
		t.Errorf("Incorrect expense ID '%v'", expenses[0].ID)
	}
	if expenses[1].UserID != 508343 {
		t.Errorf("Incorrect User ID '%v'", expenses[1].UserID)
	}
}
