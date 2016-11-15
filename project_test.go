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
