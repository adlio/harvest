package harvest

import "testing"

func TestGetTaskAssignments(t *testing.T) {
	a := testAPI()
	taskAssignmentResponse := mockResponse("taskassignments", "task-assignments-example.json")
	a.BaseURL = taskAssignmentResponse.URL
	taskassignments, err := a.GetTaskAssignments(9292184, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(taskassignments) != 9 {
		t.Errorf("Incorrect number of task assignments '%v'", len(taskassignments))
	}
	if taskassignments[0].ID != 101061850 {
		t.Errorf("Incorrect task assignment ID '%v'", taskassignments[0].ID)
	}
	if taskassignments[2].TaskID != 733904 {
		t.Errorf("Incorrect TaskID '%v'", taskassignments[2].TaskID)
	}
}

func TestGetTaskAssignment(t *testing.T) {
	a := testAPI()
	taskAssignmentResponse := mockResponse("taskassignments", "task-assignment-example.json")
	a.BaseURL = taskAssignmentResponse.URL
	taskassignment, err := a.GetTaskAssignment(3554414, 37453419, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if taskassignment.TaskID != 2086199 {
		t.Errorf("Incorrect task ID '%v'", taskassignment.TaskID)
	}
}

func TestCreateTaskAssignment(t *testing.T) {
	a := testAPI()
	res := mockRedirectResponse("taskassignments", "task-assignment-example.json")
	a.BaseURL = res.URL
	ta := TaskAssignment{ID: 123456, ProjectID: 12345}
	err := a.CreateTaskAssignment(&ta, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if ta.ID != 37453419 {
		t.Errorf("Expected ID=%d, got ID=%d", 37453419, ta.ID)
	}
}

func TestUpdateTaskAssignment(t *testing.T) {
	a := testAPI()
	res := mockRedirectResponse("taskassignments", "task-assignment-example.json")
	a.BaseURL = res.URL
	ta := TaskAssignment{ID: 12456, ProjectID: 12345}
	err := a.UpdateTaskAssignment(&ta, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if ta.ID != 37453419 {
		t.Errorf("Expected ID=%d, got ID=%d", 37453419, ta.ID)
	}
}

func TestDeleteTaskAssignment(t *testing.T) {
	a := testAPI()
	err := a.DeleteTaskAssignment(&TaskAssignment{ID: 123456, ProjectID: 12345}, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}

func TestCopyTaskAssignments(t *testing.T) {
	a := testAPI()
	res := mockDynamicPathResponse()
	a.BaseURL = res.URL
	err := a.CopyTaskAssignments(1, 2)
	if err != nil {
		t.Fatal(err)
	}
}

func TestContainsTaskIDPresent(t *testing.T) {
	ids := []*TaskAssignment{
		&TaskAssignment{TaskID: 1},
		&TaskAssignment{TaskID: 2},
	}
	if ContainsTaskID(1, ids) != true {
		t.Errorf("ContainsTaskID should be true for 1 when ids contains TaskID: 1")
	}
}

func TestContainsTaskIDMissing(t *testing.T) {
	ids := []*TaskAssignment{
		&TaskAssignment{TaskID: 1},
		&TaskAssignment{TaskID: 2},
	}
	if ContainsTaskID(10, ids) != false {
		t.Errorf("ContainsTaskID should be false for 10 when ids has no TaskID: 10")
	}
}
