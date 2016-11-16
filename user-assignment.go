package harvest

import (
	"fmt"
	"time"
)

type UserAssignmentResponse struct {
	UserAssignment *UserAssignment `json:"user_assignment"`
}

type UserAssignment struct {
	ID               int64     `json:"id"`
	UserID           int64     `json:"user_id"`
	ProjectID        int64     `json:"project_id"`
	Deactivated      bool      `json:"deactivated"`
	HourlyRate       *float64  `json:"hourly_rate"`
	IsProjectManager bool      `json:"is_project_manager"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Estimate         int64     `json:"estimate"`
}

func (a *API) GetUserAssignments(projectID int64, args Arguments) (userassignments []*UserAssignment, err error) {
	userAssignmentsResponse := make([]*UserAssignmentResponse, 0)
	path := fmt.Sprintf("/projects/%v/user_assignments", projectID)
	err = a.Get(path, args, &userAssignmentsResponse)
	for _, ua := range userAssignmentsResponse {
		userassignments = append(userassignments, ua.UserAssignment)
	}
	return userassignments, err
}

func (a *API) GetUserAssignment(projectID int64, userAssignmentID int64, args Arguments) (userassignment *UserAssignment, err error) {
	userAssignmentResponse := UserAssignmentResponse{}
	path := fmt.Sprintf("/projects/%v/user_assignments/%v", projectID, userAssignmentID)
	err = a.Get(path, args, &userAssignmentResponse)
	return userAssignmentResponse.UserAssignment, err
}
