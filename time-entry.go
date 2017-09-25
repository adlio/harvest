package harvest

import (
	"strconv"
	"time"
)

type TimeEntriesResponse struct {
	PagedResponse
	TimeEntries []*TimeEntry `json:"time_entries"`
}

type TimeEntry struct {
	ID             int64        `json:"id"`
	SpentDate      Date         `json:"spent_date"`
	User           *UserStub    `json:"user,omitempty"`
	Client         *ClientStub  `json:"client,omitempty"`
	Project        *ProjectStub `json:"project,omitempty"`
	Task           *TaskStub    `json:"task,omitempty"`
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

func (a *API) GetTimeEntries(args Arguments) ([]*TimeEntry, error) {
	entries := make([]*TimeEntry, 0)
	timeEntriesResponse := TimeEntriesResponse{}
	err := a.GetPaginated("/time_entries", args, &timeEntriesResponse, func() {
		for _, te := range timeEntriesResponse.TimeEntries {
			entries = append(entries, te)
		}
		timeEntriesResponse.TimeEntries = make([]*TimeEntry, 0)
	})
	return entries, err
}

func (a *API) GetTimeEntriesUpdatedSince(sinceDate time.Time, args Arguments) ([]*TimeEntry, error) {
	since := sinceDate.Format(time.RFC3339)
	args["updated_since"] = since
	return a.GetTimeEntries(args)
}

func (a *API) GetTimeEntriesBetween(fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	args["from"] = from
	args["to"] = to
	return a.GetTimeEntries(args)
}

func (a *API) GetTimeEntriesForProjectBetween(projectID int64, fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	args["project_id"] = strconv.FormatInt(projectID, 10)
	args["from"] = from
	args["to"] = to
	return a.GetTimeEntries(args)
}

func (a *API) GetTimeEntriesForUserBetween(userID int64, fromDate time.Time, toDate time.Time, args Arguments) ([]*TimeEntry, error) {
	from := fromDate.Format("20060102")
	to := toDate.Format("20060102")
	args["user_id"] = strconv.FormatInt(userID, 10)
	args["from"] = from
	args["to"] = to
	return a.GetTimeEntries(args)
}
