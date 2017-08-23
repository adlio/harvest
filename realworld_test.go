package harvest

import (
	"os"
	"testing"
	"time"
)

func TestCreateUpdateDeleteProject(t *testing.T) {

	realworld := os.Getenv("HARVEST_REALWORLD")

	account := os.Getenv("HARVEST_ID")
	username := os.Getenv("HARVEST_USER")
	password := os.Getenv("HARVEST_PASS")

	if realworld == "true" && account != "" && username != "" && password != "" {
		api := NewBasicAuthAPI(account, username, password)

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

		err = api.CreateProject(&p, Defaults())
		if err != nil {
			t.Fatal(err)
		}

		if p.ID == 0 {
			t.Errorf("Project %s should not have Zero ID", p.Name)
		}

		p.StartsOn = Date{time.Now()}
		p.EndsOn = Date{time.Now().AddDate(0, 0, 1)}
		p.Notes = "These are the automatically-generated notes from the library."

		err = api.UpdateProject(&p, Defaults())
		if err != nil {
			t.Fatal(err)
		}

		err = api.DeleteProject(&p, Defaults())
		if err != nil {
			t.Fatal(err)
		}

	} else {
		t.Skipf("Skipping realworld tests because HARVEST_REALWORLD != true or HARVEST_ID, HARVEST_USER, HARVEST_PASS not supplied as environment variables.")
	}
}
