package harvest

import (
	"fmt"
	"time"
)

type TaskAssignmentAddRequest struct {
	Task *Task `json:"task"`
}

type TaskAssignmentRequest struct {
	TaskAssignment *TaskAssignment `json:"task_assignment"`
}

type TaskAssignmentResponse struct {
	TaskAssignment *TaskAssignment `json:"task_assignment"`
}

type TaskAssignment struct {
	ID          int64     `json:"id,omitempty"`
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
	req := TaskAssignmentAddRequest{Task: &Task{ID: ta.TaskID}}
	resp := TaskAssignmentResponse{TaskAssignment: ta}
	path := fmt.Sprintf("/projects/%v/task_assignments", ta.ProjectID)
	err := a.Post(path, args, &req, &resp)
	if err != nil {
		return err
	}
	return a.UpdateTaskAssignment(ta, args)
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

func (a *API) CopyTaskAssignments(destProjectID int64, sourceProjectID int64) error {

	sourceTAs, err := a.GetTaskAssignments(sourceProjectID, Defaults())
	if err != nil {
		return err
	}

	destTAs, err := a.GetTaskAssignments(destProjectID, Defaults())
	if err != nil {
		return err
	}

	// Remove incorrect TaskAssignments
	for _, destTA := range destTAs {
		if !ContainsTaskID(destTA.TaskID, sourceTAs) {
			err = a.DeleteTaskAssignment(destTA, Defaults())
			if err != nil {
				return err
			}
		}
	}

	// Add missing TaskAssignments
	for _, sourceTA := range sourceTAs {
		if !ContainsTaskID(sourceTA.TaskID, destTAs) {
			ta := TaskAssignment{
				ProjectID:   destProjectID,
				TaskID:      sourceTA.TaskID,
				Billable:    sourceTA.Billable,
				Deactivated: sourceTA.Deactivated,
				Budget:      sourceTA.Budget,
				HourlyRate:  sourceTA.HourlyRate,
				Estimate:    sourceTA.Estimate,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}
			err = a.CreateTaskAssignment(&ta, Defaults())
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func ContainsTaskID(taskID int64, tas []*TaskAssignment) bool {
	for _, ta := range tas {
		if ta.TaskID == taskID {
			return true
		}
	}
	return false
}
