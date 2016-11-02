package harvest

import (
	"time"
)

type UserAssignment struct {
	ID               int64
	UserID           int64
	ProjectID        int64
	Deactivated      bool
	HourlyRate       *float64
	IsProjectManager bool
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
