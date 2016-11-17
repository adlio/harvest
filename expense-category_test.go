package harvest

import "testing"

func testExpenseCategory(t *testing.T) *ExpenseCategory {
	a := testAPI()
	expenseCategoryResponse := mockResponse("expensecategories", "expense-category-example.json")
	a.BaseURL = expenseCategoryResponse.URL
	expensecategory, err := a.GetExpenseCategory(1338056, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if expensecategory.Name != "Entertainment" {
		t.Errorf("Incorrect Name '%s'", expensecategory.Name)
	}
	return expensecategory
}

func TestGetExpenseCategory(t *testing.T) {
	expensecategory := testExpenseCategory(t)
	if expensecategory == nil {
		t.Fatal("testExpenseCategory() returned nil instead of expensecategory")
	}
	if expensecategory.ID != 1338056 {
		t.Errorf("Incorrect ID '%v'", expensecategory.ID)
	}
}

func TestGetExpenseCategories(t *testing.T) {
	a := testAPI()
	expenseCategoryResponse := mockResponse("expensecategories", "expense-categories-example.json")
	a.BaseURL = expenseCategoryResponse.URL
	expensecategories, err := a.GetExpenseCategories(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(expensecategories) != 2 {
		t.Errorf("Incorrect number of expense categories '%v'", len(expensecategories))
	}
	if expensecategories[1].Name != "Work" {
		t.Errorf("Incorrect Name '%s'", expensecategories[1].Name)
	}
	if expensecategories[0].ID != 1338056 {
		t.Errorf("Incorrect ID '%v'", expensecategories[0].ID)
	}
}
