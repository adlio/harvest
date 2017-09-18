package harvest

import (
	"fmt"
	"time"
)

type TaskAssignmentsResponse struct {
	TaskAssignments []*TaskAssignment `json:"task_assignments"`

	PerPage      int64  `json:"per_page"`
	TotalPages   int64  `json:"total_pages"`
	TotalEntries int64  `json:"total_entries"`
	NextPage     *int64 `json:"next_page"`
	PreviousPage *int64 `json:"previous_page"`
	Page         int64  `json:"page"`
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

func (a *API) GetTaskAssignments(projectID int64, args Arguments) (taskAssignments []*TaskAssignment, err error) {
	taskAssignmentsResponse := TaskAssignmentsResponse{}
	path := fmt.Sprintf("/projects/%v/task_assignments", projectID)
	err = a.Get(path, args, &taskAssignmentsResponse)
	return taskAssignmentsResponse.TaskAssignments, err
}

func (a *API) GetTaskAssignment(projectID int64, taskAssignmentID int64, args Arguments) (taskAssignment *TaskAssignment, err error) {
	taskAssignment = &TaskAssignment{}
	path := fmt.Sprintf("/projects/%v/task_assignments/%v", projectID, taskAssignmentID)
	err = a.Get(path, args, taskAssignment)
	return taskAssignment, err
}

func (a *API) CreateTaskAssignment(ta *TaskAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/task_assignments", ta.ProjectID)
	return a.Post(path, args, ta, ta)
}

func (a *API) UpdateTaskAssignment(ta *TaskAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/task_assignments/%v", ta.ProjectID, ta.ID)
	return a.Put(path, args, ta, ta)
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

	// Add missing TaskAssignments, update existing ones
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
		} else {
			for _, newTA := range destTAs {
				if newTA.TaskID == sourceTA.TaskID && TaskAssignmentAttributesDiffer(newTA, sourceTA) {
					newTA.Billable = sourceTA.Billable
					newTA.Deactivated = sourceTA.Deactivated
					newTA.Budget = sourceTA.Budget
					newTA.HourlyRate = sourceTA.HourlyRate
					newTA.Estimate = sourceTA.Estimate
					err = a.UpdateTaskAssignment(newTA, Defaults())
					if err != nil {
						return err
					}
				}
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

func TaskAssignmentAttributesDiffer(ta1, ta2 *TaskAssignment) bool {
	if ta1.Billable != ta2.Billable {
		return true
	}
	if ta1.Deactivated != ta2.Deactivated {
		return true
	}
	if !HaveSameFloat64Value(ta1.Budget, ta2.Budget) {
		return true
	}
	if !HaveSameFloat64Value(ta1.HourlyRate, ta2.HourlyRate) {
		return true
	}
	if !HaveSameFloat64Value(ta1.Estimate, ta2.Estimate) {
		return true
	}
	return false
}
