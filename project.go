package harvest

import (
	"time"
)

type Project struct {
	ID                               int64
	ClientID                         int64
	Name                             string
	Code                             string
	Billable                         bool
	BillBy                           string
	HourlyRate                       *float64
	BudgetBy                         string
	Budget                           *float64
	NotifyWhenOverBudget             bool
	OverBudgetNotificationPercentage float64
	OverBudgetNotifiedAt             time.Time
	ShowBudgetToAll                  bool
	CreatedAt                        time.Time
	UpdatedAt                        time.Time
	StartsOn                         *Date
	EndsOn                           *Date
	Estimate                         *float64
	EstimateBy                       string
	Notes                            string
	CostBudget                       *float64
	CostBudgetIncludeExpenses        bool
}
