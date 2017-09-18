package harvest

import (
	"fmt"
	"time"
)

type UserAssignmentsResponse struct {
	UserAssignments []*UserAssignment `json:"user_assignments"`
	PerPage         int64             `json:"per_page"`
	TotalPages      int64             `json:"total_pages"`
	TotalEntries    int64             `json:"total_entries"`
	NextPage        *int64            `json:"next_page"`
	PreviousPage    *int64            `json:"previous_page"`
	Page            int64             `json:"page"`
}

type UserAssignment struct {
	ID               int64     `json:"id,omitempty"`
	UserID           int64     `json:"user_id"`
	ProjectID        int64     `json:"project_id"`
	Deactivated      bool      `json:"deactivated"`
	HourlyRate       *float64  `json:"hourly_rate"`
	IsProjectManager bool      `json:"is_project_manager"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	Estimate         float64   `json:"estimate"`
}

func (a *API) GetUserAssignments(projectID int64, args Arguments) (userAssignments []*UserAssignment, err error) {
	userAssignmentsResponse := UserAssignmentsResponse{}
	path := fmt.Sprintf("/projects/%v/user_assignments", projectID)
	err = a.Get(path, args, &userAssignmentsResponse)
	return userAssignmentsResponse.UserAssignments, err
}

func (a *API) GetUserAssignment(projectID int64, userAssignmentID int64, args Arguments) (userAssignment *UserAssignment, err error) {
	userAssignment = &UserAssignment{}
	path := fmt.Sprintf("/projects/%v/user_assignments/%v", projectID, userAssignmentID)
	err = a.Get(path, args, userAssignment)
	return userAssignment, err
}

func (a *API) CreateUserAssignment(ua *UserAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/user_assignments", ua.ProjectID)
	return a.Post(path, args, ua, ua)
}

func (a *API) UpdateUserAssignment(ua *UserAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/user_assignments/%v", ua.ProjectID, ua.ID)
	return a.Put(path, args, &ua, &ua)
}

func (a *API) DeleteUserAssignment(ua *UserAssignment, args Arguments) error {
	path := fmt.Sprintf("/projects/%v/user_assignments/%v", ua.ProjectID, ua.ID)
	return a.Delete(path, args)
}

func (a *API) CopyUserAssignments(destProjectID int64, sourceProjectID int64) error {

	originalUAs, err := a.GetUserAssignments(sourceProjectID, Defaults())
	if err != nil {
		return err
	}

	newUAs, err := a.GetUserAssignments(destProjectID, Defaults())
	if err != nil {
		return err
	}

	// Remove incorrect UserAssignments
	for _, newUA := range newUAs {
		if !ContainsUserID(newUA.UserID, originalUAs) {
			err = a.DeleteUserAssignment(newUA, Defaults())
			if err != nil {
				return err
			}
		}
	}

	// Add missing UserAssignments, update existing ones
	for _, originalUA := range originalUAs {
		if !ContainsUserID(originalUA.UserID, newUAs) {
			err = a.CreateUserAssignment(&UserAssignment{
				ProjectID:        destProjectID,
				UserID:           originalUA.UserID,
				Deactivated:      originalUA.Deactivated,
				HourlyRate:       originalUA.HourlyRate,
				IsProjectManager: originalUA.IsProjectManager,
				Estimate:         originalUA.Estimate,
				UpdatedAt:        time.Now(),
				CreatedAt:        time.Now(),
			}, Defaults())
			if err != nil {
				return err
			}
		} else {
			for _, newUA := range newUAs {
				if newUA.UserID == originalUA.UserID && UserAssignmentAttributesDiffer(newUA, originalUA) {
					newUA.Deactivated = originalUA.Deactivated
					newUA.HourlyRate = originalUA.HourlyRate
					newUA.IsProjectManager = originalUA.IsProjectManager
					newUA.Estimate = originalUA.Estimate
					err = a.UpdateUserAssignment(newUA, Defaults())
					if err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func ContainsUserID(userID int64, uas []*UserAssignment) bool {
	for _, ua := range uas {
		if ua.UserID == userID {
			return true
		}
	}
	return false
}

func UserAssignmentAttributesDiffer(ua1, ua2 *UserAssignment) bool {
	if ua1.Deactivated != ua2.Deactivated {
		return true
	}
	if !HaveSameFloat64Value(ua1.HourlyRate, ua2.HourlyRate) {
		return true
	}
	if ua1.IsProjectManager != ua2.IsProjectManager {
		return true
	}
	if ua1.Estimate != ua2.Estimate {
		return true
	}
	return false
}
