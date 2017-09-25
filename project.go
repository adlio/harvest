package harvest

import (
	"fmt"
	"time"
)

type ProjectsResponse struct {
	PagedResponse
	Projects []*Project `json:"projects"`
}

type ProjectStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Project struct {
	ID                               int64     `json:"id,omitempty"`
	ClientID                         int64     `json:"client_id,omitempty"`
	Name                             string    `json:"name,omitempty"`
	Code                             string    `json:"code,omitempty"`
	Active                           bool      `json:"active"`
	Billable                         bool      `json:"billable"`
	BillBy                           string    `json:"bill_by,omitempty"`
	HourlyRate                       *float64  `json:"hourly_rate,omitempty"`
	BudgetBy                         string    `json:"budget_by,omitempty"`
	Budget                           *float64  `json:"budget,omitempty"`
	NotifyWhenOverBudget             bool      `json:"notify_when_over_budget"`
	OverBudgetNotificationPercentage float64   `json:"over_budget_notification_percentage,omitempty"`
	OverBudgetNotifiedAt             *Date     `json:"over_budget_notified_at,omitempty"`
	ShowBudgetToAll                  bool      `json:"show_budget_to_all"`
	CreatedAt                        time.Time `json:"created_at,omitempty"`
	UpdatedAt                        time.Time `json:"updated_at,omitempty"`
	StartsOn                         Date      `json:"starts_on,omitempty"`
	EndsOn                           Date      `json:"ends_on,omitempty"`
	Estimate                         *float64  `json:"estimate,omitempty"`
	EstimateBy                       string    `json:"estimate_by,omitempty"`
	HintEarliestRecordAt             Date      `json:"hint_earliest_record_at,omitempty"`
	HintLatestRecordAt               Date      `json:"hint_latest_record_at,omitempty"`
	Notes                            string    `json:"notes,omitempty"`
	CostBudget                       *float64  `json:"cost_budget,omitempty"`
	CostBudgetIncludeExpenses        bool      `json:"cost_budget_include_expenses"`
}

func (a *API) GetProject(projectID int64, args Arguments) (project *Project, err error) {
	project = &Project{}
	path := fmt.Sprintf("/projects/%d", projectID)
	err = a.Get(path, args, project)
	return project, err
}

func (a *API) GetProjects(args Arguments) (projects []*Project, err error) {
	projects = make([]*Project, 0)
	projectsResponse := ProjectsResponse{}
	err = a.GetPaginated("/projects", args, &projectsResponse, func() {
		for _, p := range projectsResponse.Projects {
			projects = append(projects, p)
		}
		projectsResponse.Projects = make([]*Project, 0)
	})
	return projects, err
}

func (a *API) SaveProject(p *Project, args Arguments) error {
	if p.ID != 0 {
		return a.UpdateProject(p, args)
	} else {
		return a.CreateProject(p, args)
	}
}

func (a *API) UpdateProject(p *Project, args Arguments) error {
	path := fmt.Sprintf("/projects/%d", p.ID)
	return a.Put(path, args, p, p)
}

func (a *API) CreateProject(p *Project, args Arguments) error {
	return a.Post("/projects", args, p, p)
}

func (a *API) DeleteProject(p *Project, args Arguments) error {
	path := fmt.Sprintf("/projects/%d", p.ID)
	return a.Delete(path, args)
}

func (a *API) DuplicateProject(sourceProjectID int64, newName string) (*Project, error) {

	var project *Project

	project, err := a.GetProject(sourceProjectID, Defaults())
	if err != nil {
		return nil, err
	}

	project.ID = 0
	project.Name = newName

	err = a.CreateProject(project, Defaults())
	if err != nil {
		return nil, err
	}

	err = a.CopyTaskAssignments(project.ID, sourceProjectID)
	if err != nil {
		return nil, err
	}

	err = a.CopyUserAssignments(project.ID, sourceProjectID)
	if err != nil {
		return nil, err
	}

	return project, nil
}
