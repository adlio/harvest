package harvest

import (
	"fmt"
	"time"
)

type ProjectResponse struct {
	Project *Project `json:"project"`
}

type Project struct {
	ID                               int64     `json:"id"`
	ClientID                         int64     `json:"client_id"`
	Name                             string    `json:"name"`
	Code                             string    `json:"code"`
	Active                           bool      `json:"active"`
	Billable                         bool      `json:"billable"`
	BillBy                           string    `json:"bill_by"`
	HourlyRate                       *float64  `json:"hourly_rate"`
	BudgetBy                         string    `json:"budget_by"`
	Budget                           *float64  `json:"budget"`
	NotifyWhenOverBudget             bool      `json:"notify_when_over_budget"`
	OverBudgetNotificationPercentage float64   `json:"over_budget_notification_percentage"`
	OverBudgetNotifiedAt             *Date     `json:"over_budget_notified_at"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all"`
	CreatedAt                        time.Time `json:"created_at"`
	UpdatedAt                        time.Time `json:"updated_at"`
	StartsOn                         Date      `json:"starts_on"`
	EndsOn                           Date      `json:"ends_on"`
	Estimate                         *float64  `json:"estimate"`
	EstimateBy                       string    `json:"estimate_by"`
	HintEarliestRecordAt             Date      `json:"hint_earliest_record_at"`
	HintLatestRecordAt               Date      `json:"hint_latest_record_at"`
	Notes                            string    `json:"notes"`
	CostBudget                       *float64  `json:"cost_budget"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses"`
}

func (a *API) GetProject(projectID int64, args Arguments) (project *Project, err error) {
	projectResponse := ProjectResponse{}
	path := fmt.Sprintf("/projects/%d", projectID)
	err = a.Get(path, args, &projectResponse)
	return projectResponse.Project, err
}

func (a *API) GetProjects(args Arguments) (projects []*Project, err error) {
	projects = make([]*Project, 0)
	projectsResponse := make([]*ProjectResponse, 0)
	path := fmt.Sprintf("/projects")
	err = a.Get(path, args, &projectsResponse)
	for _, pr := range projectsResponse {
		projects = append(projects, pr.Project)
	}
	return projects, err
}
