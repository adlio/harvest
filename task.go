package harvest

import (
	"time"
)

type Task struct {
	ID                int64
	Name              string
	BillableByDefault bool
	Deactivated       bool
	DefaultHourlyRate float64
	IsDefault         bool
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
