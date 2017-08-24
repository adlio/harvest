package harvest

import (
	"fmt"
	"time"
)

type TaskAssignmentRequest struct {
	TaskAssignment *TaskAssignment `json:"task_assignment"`
}

type TaskAssignmentResponse struct {
	TaskAssignment *TaskAssignment `json:"task_assignment"`
}

type TaskAssignment struct {
	ID          int64     `json:"id"`
	ProjectID   int64     `json:"project_id"`
	TaskID      int64     `json:"task_id"`
	Billable    bool      `json:"billable"`
	Deactivated bool      `json:"deactivated"`
	Budget      *float64  `json:"budget"`
	HourlyRate  *float64  `json:"hourly_rate"`
	Estimate    *float64  `json:"estimate"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`
}

func (a *API) GetTaskAssignments(projectID int64, args Arguments) (taskassignments []*TaskAssignment, err error) {
	taskAssignmentsResponse := make([]*TaskAssignmentResponse, 0)
	path := fmt.Sprintf("/projects/%v/task_assignments", projectID)
	err = a.Get(path, args, &taskAssignmentsResponse)
	for _, ta := range taskAssignmentsResponse {
		taskassignments = append(taskassignments, ta.TaskAssignment)
	}
	return taskassignments, err
}

func (a *API) GetTaskAssignment(projectID int64, taskAssignmentID int64, args Arguments) (taskassignment *TaskAssignment, err error) {
	taskAssignmentResponse := TaskAssignmentResponse{}
	path := fmt.Sprintf("/projects/%v/task_assignments/%v", projectID, taskAssignmentID)
	err = a.Get(path, args, &taskAssignmentResponse)
	return taskAssignmentResponse.TaskAssignment, err
}

func (a *API) CreateTaskAssignment(ta *TaskAssignment, args Arguments) error {
	req := TaskAssignmentRequest{TaskAssignment: ta}
	resp := TaskAssignmentResponse{TaskAssignment: ta}
	path := fmt.Sprintf("/projects/%v/task_assignments", ta.ProjectID)
	return a.Post(path, args, &req, &resp)
}

func (a *API) UpdateTaskAssignment(ta *TaskAssignment, args Arguments) error {
	req := TaskAssignmentRequest{TaskAssignment: ta}
	resp := TaskAssignmentResponse{TaskAssignment: ta}
	path := fmt.Sprintf("/projects/%v/task_assignments/%v", ta.ProjectID, ta.ID)
	return a.Put(path, args, &req, &resp)
}

func (a *API) DeleteTaskAssignment(ta *TaskAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/task_assignments/%v", ta.ProjectID, ta.ID)
	return a.Delete(path, args)
}

func ContainsTaskID(taskID int64, tas []*TaskAssignment) bool {
	for _, ta := range tas {
		if ta.TaskID == taskID {
			return true
		}
	}
	return false
}
