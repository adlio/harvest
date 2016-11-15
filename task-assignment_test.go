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
