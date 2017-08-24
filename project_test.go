package harvest

import "testing"

func TestGetProject(t *testing.T) {
	project := testProject(t)

	if project == nil {
		t.Fatal("testProject() returned nil instead of project")
	}
	if project.ID != 9292184 {
		t.Errorf("Incorrect project ID '%v'", project.ID)
	}
}

func testProject(t *testing.T) *Project {
	a := testAPI()
	projectResponse := mockResponse("projects", "project-example.json")
	a.BaseURL = projectResponse.URL
	project, err := a.GetProject(9292184, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if project.Name != "NYRA Web Site Builds x 5" {
		t.Errorf("Incorrect Project Name '%s'", project.Name)
	}
	if project.ID != 9292184 {
		t.Errorf("Incorrect Project ID '%v'", project.ID)
	}
	if project.BudgetBy != "project" {
		t.Errorf("Incorrect Project budget by '%s'", project.BudgetBy)
	}
	return project
}

func TestGetProjects(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("projects", "projects-example.json")
	a.BaseURL = projectResponse.URL
	projects, err := a.GetProjects(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(projects) != 4 {
		t.Errorf("Incorrect number of projects '%v'", len(projects))
	}
	if projects[1].Name != "KBL - Vine Website Build" {
		t.Errorf("Incorrect Project Name '%s'", projects[1].Name)
	}
	if projects[0].ID != 9292184 {
		t.Errorf("Incorrect Project ID '%v'", projects[2].ID)
	}
	if projects[3].BudgetBy != "project" {
		t.Errorf("Incorrect Project budget by '%s'", projects[3].BudgetBy)
	}
}

func TestGetProjectWithStartEndDates(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("projects", "12670372.json")
	a.BaseURL = projectResponse.URL
	project, err := a.GetProject(12670372, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if project.Name != "TEST" {
		t.Errorf("Expected 'TEST', got '%s'.", project.Name)
	}

	if !project.Active {
		t.Error("Project should have been active")
	}

	if !project.NotifyWhenOverBudget {
		t.Error("Should notify when over budget.")
	}

	if project.OverBudgetNotificationPercentage != 80.0 {
		t.Errorf("Expected '80.00', got '%0.2f'", project.OverBudgetNotificationPercentage)
	}

	if project.StartsOn.Format("2006-01-02") != "2017-01-01" {
		t.Errorf("Expected '2017-01-01', got '%s'", project.StartsOn.Format("2006-01-02"))
	}

	if project.EndsOn.Format("2006-01-02") != "2017-01-31" {
		t.Errorf("Expected '2017-01-31', got '%s'", project.EndsOn.Format("2006-01-02"))
	}

	if project.ShowBudgetToAll {
		t.Error("Project budget should be hidden")
	}

	if *project.Estimate != 24.0 {
		t.Errorf("Expected 24.0 budget. got %0.2f", *project.Estimate)
	}

	if project.EstimateBy != "project" {
		t.Errorf("Expected 'project' EstimateBy, got '%s'", project.EstimateBy)
	}

	if project.HintEarliestRecordAt.Format("2006-01-02") != "2017-01-03" {
		t.Errorf("Expected '2017-01-03', got '%s'", project.HintEarliestRecordAt.Format("2006-01-02"))
	}

	if project.Notes != "This is the notes." {
		t.Errorf("Expected 'This is the notes.', Got '%s'", project.Notes)
	}

	if *project.HourlyRate != 120.0 {
		t.Errorf("Expected '120.00', got '%0.2f'", *project.HourlyRate)
	}
}

func TestCreateProject(t *testing.T) {
	a := testAPI()
	projectResponse := mockRedirectResponse("projects", "12670372.json")
	a.BaseURL = projectResponse.URL

	p := Project{
		Name:     "New Name",
		Active:   true,
		ClientID: 12345,
	}

	err := a.SaveProject(&p, Defaults()) // Will call a.CreateProject()
	if err != nil {
		t.Fatal(err)
	}

	if *p.HourlyRate != 120.0 {
		t.Errorf("Hourly rate should have been picked up in the CreateProject response")
	}
}

func TestUpdateProject(t *testing.T) {
	a := testAPI()
	projectResponse := mockRedirectResponse("projects", "12670372.json")
	a.BaseURL = projectResponse.URL

	p := Project{
		ID:       1234,
		Name:     "New Name",
		Active:   true,
		ClientID: 12345,
	}

	err := a.SaveProject(&p, Defaults()) // Will call a.UpdateProject()
	if err != nil {
		t.Fatal(err)
	}

	if *p.HourlyRate != 120.0 {
		t.Errorf("Hourly rate should have been picked up in the CreateProject response")
	}
}

func TestDuplicateProject(t *testing.T) {
	a := testAPI()
	res := mockDynamicPathResponse()
	a.BaseURL = res.URL

	project, err := a.DuplicateProject(1, "Experimental Name")
	if err != nil {
		t.Fatal(err)
	}

	if project.ID != 2 {
		t.Errorf("Expected new project with ID=%d, but got ID=%d", 2, project.ID)
	}
}

func TestDeleteProject(t *testing.T) {
	a := testAPI()
	err := a.DeleteProject(&Project{ID: 1}, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}
