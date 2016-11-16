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
