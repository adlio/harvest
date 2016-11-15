package harvest

import (
	"fmt"
	"time"
)

type DayEntryResponse struct {
	DayEntries []*DayEntry `json:"day_entries"`
}

type DayEntry struct {
	ID                int64     `json:"id"`
	UserID            int64     `json:"user_id"`
	SpentAt           string    `json:"spent_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ProjectID         string    `json:"project_id"`
	TaskID            string    `json:"task_id"`
	Project           string    `json:"project"`
	Task              string    `json:"task"`
	Client            string    `json:"client"`
	Notes             string    `json:"notes"`
	HoursWithoutTimer int64     `json:"hours_without_timer"`
	Hours             int64     `json:"hours"`
}

func (a *API) GetTodayEntry(args Arguments) (dayentries []*DayEntry, err error) {
	dayEntriesResponse := DayEntryResponse{}
	path := fmt.Sprintf("/daily?slim=1")
	err = a.Get(path, args, &dayEntriesResponse)
	return dayEntriesResponse.DayEntries, err
}
