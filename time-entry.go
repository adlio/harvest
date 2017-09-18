package harvest

import (
	"fmt"
	"time"
)

type TimeEntriesResponse struct {
	TimeEntries  []*TimeEntry `json:"time_entries"`
	PerPage      int64        `json:"per_page"`
	TotalPages   int64        `json:"total_pages"`
	TotalEntries int64        `json:"total_entries"`
	NextPage     *int64       `json:"next_page"`
	PreviousPage *int64       `json:"previous_page"`
	Page         int64        `json:"page"`
}

type TimeEntry struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	SpentDateRaw   string `json:"spent_at"`
	SpentDate      time.Time
	User           *UserStub    `json:"user,omitempty"`
	Client         *ClientStub  `json:"client,omitempty"`
	Project        *ProjectStub `json:"project,omitempty"`
	Task           *TaskStub    `json:"task:omitempty"`
	HoursWithTimer float64      `json:"hours_with_timer"`
	Hours          float64      `json:"hours"`
	Notes          string       `json:"notes"`
	CreatedAt      time.Time    `json:"created_at"`
	UpdatedAt      time.Time    `json:"updated_at"`
	IsLocked       bool         `json:"is_locked"`
	LockedReason   string       `json:"locked_reason"`
	IsClosed       bool         `json:"is_closed"`
	IsBilled       bool         `json:"is_billed"`
	TimerStartedAt *time.Time   `json:"timer_started_at"`
	StartedTime    string       `json:"started_time"`
	EndedTime      string       `json:"ended_time"`
	IsRunning      bool         `json:"is_running"`
	IsBillable     bool         `json:"billable"`
	IsBudgeted     bool         `json:"budgeted"`
	BillableRate   float64      `json:"billable_rate"`
	CostRate       float64      `json:"cost_rate"`
}

type UserStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ClientStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ProjectStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type TaskStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

func (a *API) GetTimeEntriesUpdatedSince(sinceDate time.Time, args Arguments) ([]*TimeEntry, error) {
	timeEntriesResponse := TimeEntriesResponse{}
	since := sinceDate.Format(time.RFC3339)
	path := fmt.Sprintf("/time_entries?updated_since=%s", since)
	err := a.Get(path, args, &timeEntriesResponse)
	return timeEntriesResponse.TimeEntries, err
}

func (a *API) GetTimeEntriesBetween(fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	timeEntriesResponse := TimeEntriesResponse{}
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	path := fmt.Sprintf("/time_entries?from=%s&to=%s", from, to)
	err := a.Get(path, args, &timeEntriesResponse)
	return timeEntriesResponse.TimeEntries, err
}

func (a *API) GetTimeEntriesForProjectBetween(projectID int64, fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	timeEntriesResponse := TimeEntriesResponse{}
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	path := fmt.Sprintf("/time_entries?project_id=%d&from=%s&to=%s", projectID, from, to)
	err := a.Get(path, args, &timeEntriesResponse)
	return timeEntriesResponse.TimeEntries, err
}

func (a *API) GetTimeEntriesForUserBetween(userID int64, fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	timeEntriesResponse := TimeEntriesResponse{}
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	path := fmt.Sprintf("/time_entries?user_id=%d&from=%s&to=%s", userID, from, to)
	err := a.Get(path, args, &timeEntriesResponse)
	return timeEntriesResponse.TimeEntries, err
}
