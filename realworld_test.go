package harvest

import (
	"os"
	"testing"
	"time"
)

func TestCreateUpdateDeleteProject(t *testing.T) {
	api := realWorldTestAPI(t)

	clients, err := api.GetClients(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	p := Project{
		Name:     "Test Harvest API Project from github.com/adlio/harvest",
		ClientID: clients[0].ID,
		Active:   true,
		BillBy:   "Project",
	}

	// Create Project
	err = api.CreateProject(&p, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if p.ID == 0 {
		t.Errorf("Project %s should not have Zero ID", p.Name)
	}

	// Now make a change
	p.StartsOn = Date{time.Now()}
	p.EndsOn = Date{time.Now().AddDate(0, 0, 1)}
	p.Notes = "These are the automatically-generated notes from the library."

	err = api.UpdateProject(&p, Defaults())
	if err != nil {
		t.Fatal(err)
	}

	// Re-fetch to verify the update
	p2, err := api.GetProject(p.ID, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if p2.Notes != "These are the automatically-generated notes from the library." {
		t.Errorf("UpdateProject() seemed to have failed to change the notes on Project %d", p2.ID)
	}

	// Delete all the UserAssignments.. this avoids a mass-email about the delete
	// to all attached staff.
	assignments, err := api.GetUserAssignments(p.ID, Defaults())
	for i, ua := range assignments {
		if i == 0 {
			ua.IsProjectManager = true
			api.UpdateUserAssignment(ua, Defaults())
		}
		api.DeleteUserAssignment(ua, Defaults())
	}

	// Clean Up Our Mess
	err = api.DeleteProject(&p, Defaults())
	if err != nil {
		t.Fatal(err)
	}
}

func realWorldTestAPI(t *testing.T) *API {
	realworld := os.Getenv("HARVEST_REALWORLD")

	account := os.Getenv("HARVEST_ID")
	username := os.Getenv("HARVEST_USER")
	password := os.Getenv("HARVEST_PASS")

	if realworld == "true" && account != "" && username != "" && password != "" {
		return NewBasicAuthAPI(account, username, password)
	} else {
		t.Skipf("Skipping realworld tests because HARVEST_REALWORLD != true or HARVEST_ID, HARVEST_USER, HARVEST_PASS not supplied as environment variables.")
		return nil
	}
}
