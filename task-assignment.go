package harvest

import (
	"time"
)

type TaskAssignment struct {
	ID          int64
	ProjectID   int64
	TaskID      int64
	Billable    bool
	Deactivated bool
	Budget      *float64
	HourlyRate  *float64
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
