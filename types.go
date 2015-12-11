package goharvest

import (
	"time"
)

type Date time.Time
type Timestamp time.Time

type Client struct {
	ID                      int
	Name                    string
	CreatedAt               Timestamp
	UpdatedAt               Timestamp
	HighriseID              int
	CacheVersion            int
	Currency                string
	Details                 string
	DefaultInvoiceTimeframe string
	LastInvoiceKind         string
}

type Project struct {
	ID                               int
	ClientID                         int
	Name                             string
	Code                             string
	Billable                         bool
	BillBy                           string
	HourlyRate                       *float
	BudgetBy                         string
	Budget                           *float
	NotifyWhenOverBudget             bool
	OverBudgetNotificationPercentage float
	OverBudgetNotifiedAt             Date
	ShowBudgetToAll                  bool
	CreatedAt                        Timestamp
	UpdatedAt                        Timestamp
	StartsOn                         *Date
	EndsOn                           *Date
	Estimate                         *float
	EstimateBy                       string
	HintEarliestRecordAt             *Date
	HintLatestRecordAt               *Date
	Notes                            string
	CostBudget                       *float
	CostBudgetIncludeExpenses        bool
}

type Task struct {
	ID                int
	Name              string
	BillableByDefault bool
	Deactivated       bool
	DefaultHourlyRate float
	IsDefault         bool
	CreatedAt         Timestamp
	UpdatedAt         Timestamp
}

type TaskAssignment struct {
	ID          int
	ProjectID   int
	TaskID      int
	Billable    bool
	Deactivated bool
	Budget      *float
	HourlyRate  *float
	UpdatedAt   Timestamp
	CreatedAt   Timestamp
}

type DayEntry struct {
	ID             int
	SpentAt        Date
	UserID         int
	Client         string
	ProjectID      int
	Project        string
	TaskID         int
	Task           string
	Hours          float
	Notes          string
	TimerStartedAt Timestamp
	CreatedAt      Timestamp
	UpdatedAt      Timestamp
	StartedAt      string
	EndedAt        string
}

type Company struct {
	Name               string  `xml:"name"`
	BaseURI            string  `xml:"base-uri"`
	FullDomain         string  `xml:"full-domain"`
	Active             string  `xml:"active"`
	WeekStartDay       string  `xml:"week-start-day"`
	TimeFormat         string  `xml:"time-format"`
	Clock              string  `xml:"clock"`
	DecimalSymbol      string  `xml:"decimal-symbol"`
	ThousandsSeparator string  `xml:"thousands-separator"`
	ColorScheme        string  `xml:"color-scheme"`
	Modules            Modules `xml:"modules"`
}

type Modules struct {
	Expenses  bool `xml:"expenses"`
	Invoices  bool `xml:"invoices"`
	Estimates bool `xml:"estimates"`
	Approval  bool `xml:"approval"`
}

type User struct {
	ID                        int                `xml:"id"`
	FirstName                 string             `xml:"first-name"`
	LastName                  string             `xml:"last-name"`
	Email                     string             `xml:"email"`
	Admin                     bool               `xml:"admin"`
	AvatarURL                 string             `xml:"avatar-url"`
	Timezone                  string             `xml:"timezone"`
	TimezoneUTCOffset         int                `xml:"timezone-utc-offset"`
	TimestampTimers           bool               `xml:"timestamp-timers"`
	ProjectManagerPermissions ManagerPermissions `xml:"project-manager"`
}

type ProjectManagerPermissions struct {
	IsProjectManager  bool `xml:"is-project-manager"`
	CanSeeRates       bool `xml:"can-see-rates"`
	CanCreateProjects bool `xml:"can-create-projects"`
	CanCreateInvoices bool `xml:"can-create-invoices"`
}

type UserAssignment struct {
	ID               int
	UserID           int
	ProjectID        int
	Deactivated      bool
	HourlyRate       *float
	IsProjectManager bool
	CreatedAt        Timestamp
	UpdatedAt        Timestamp
}
