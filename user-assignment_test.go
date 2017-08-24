package harvest

import "testing"

func TestGetUserAssignments(t *testing.T) {
	a := testAPI()
	userAssignmentResponse := mockResponse("userassignments", "user-assignments-example.json")
	a.BaseURL = userAssignmentResponse.URL
	userassignments, err := a.GetUserAssignments(9876543, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(userassignments) != 2 {
		t.Errorf("Incorrect number of user assignments '%v'", len(userassignments))
	}
	if userassignments[1].UserID != 611911 {
		t.Errorf("Incorrect user ID '%v'", userassignments[1].UserID)
	}
	if userassignments[0].IsProjectManager != false {
		t.Errorf("User Is Project Manager incorrectly set to '%v'", userassignments[0].IsProjectManager)
	}
}

func TestGetUserAssignment(t *testing.T) {
	a := testAPI()
	userAssignmentResponse := mockResponse("userassignments", "user-assignment-example.json")
	a.BaseURL = userAssignmentResponse.URL
	userassignment, err := a.GetUserAssignment(9876543, 88888888, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if userassignment.UserID != 1000000 {
		t.Errorf("Incorrect user ID '%v'", userassignment.UserID)
	}
}

func TestCreateUserAssignment(t *testing.T) {
	a := testAPI()
	res := mockRedirectResponse("userassignments", "user-assignment-example.json")
	a.BaseURL = res.URL
	ta := UserAssignment{UserID: 123456, ProjectID: 12345}
	err := a.CreateUserAssignment(&ta, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if ta.ID != 88888888 {
		t.Errorf("Expected ID=%d, got ID=%d", 88888888, ta.ID)
	}
}

func TestUpdateUserAssignment(t *testing.T) {
	a := testAPI()
	res := mockRedirectResponse("userassignments", "user-assignment-example.json")
	a.BaseURL = res.URL
	ta := UserAssignment{UserID: 12456, ProjectID: 12345}
	err := a.UpdateUserAssignment(&ta, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if ta.ID != 88888888 {
		t.Errorf("Expected ID=%d, got ID=%d", 88888888, ta.ID)
	}
}

func TestDeleteUserAssignment(t *testing.T) {
	a := testAPI()
	err := a.DeleteUserAssignment(&UserAssignment{ID: 123456, ProjectID: 12345}, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyUserAssignments(t *testing.T) {
	a := testAPI()
	res := mockDynamicPathResponse()
	a.BaseURL = res.URL
	err := a.CopyUserAssignments(2, 1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContainsUserIDPresent(t *testing.T) {
	ids := []*UserAssignment{
		&UserAssignment{UserID: 1},
		&UserAssignment{UserID: 2},
	}
	if ContainsUserID(1, ids) != true {
		t.Errorf("ContainsUserID should be true for 1 when ids contains UserID: 1")
	}
}

func TestContainsUserIDMissing(t *testing.T) {
	ids := []*UserAssignment{
		&UserAssignment{UserID: 1},
		&UserAssignment{UserID: 2},
	}
	if ContainsUserID(10, ids) != false {
		t.Errorf("ContainsUserID should be false for 10 when ids has no UserID: 10")
	}
}
