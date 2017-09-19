package harvest

import (
	"os"
	"testing"
	"time"
)

func TestRealWorldGetTasks(t *testing.T) {
	api := realWorldTestAPI(t)

	tasks, err := api.GetTasks(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(tasks) < 1 {
		t.Error("GetTasks() returned no tasks. Are you testing with an empty Harvest account?")
	}

	task, err := api.GetTask(tasks[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}
	if task.Name == "" {
		t.Error("Task name was blank")
	}
}

func TestRealWorldGetClients(t *testing.T) {
	api := realWorldTestAPI(t)

	clients, err := api.GetClients(Defaults())
	if err != nil {
		t.Error(err)
	}
	if len(clients) < 1 {
		t.Error("GetClients() returned no clients. Are you testing with an empty Harvest account?")
	}

	client, err := api.GetClient(clients[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}
	if client.Name == "" {
		t.Error("Client name was blank")
	}

	_, err = api.GetClientContacts(client.ID, Defaults())
	if err != nil {
		t.Error(err)
	}
}

func TestRealWorldGetProjects(t *testing.T) {
	api := realWorldTestAPI(t)

	projects, err := api.GetProjects(Defaults())
	if err != nil {
		t.Error(err)
	}
	if len(projects) < 1 {
		t.Error("GetProjects() returned no projects. Are you testing with an empty Harvest account?")
	}

	project, err := api.GetProject(projects[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}
	if project.Name == "" {
		t.Error("Project name was blank")
	}

	uas, err := api.GetUserAssignments(project.ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(uas) < 1 {
		t.Errorf("Project %d %s didn't contain any user assignments.", project.ID, project.Name)
	}

	tas, err := api.GetTaskAssignments(project.ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(tas) < 1 {
		t.Errorf("Project #%d %s didn't contain any task assignments.", project.ID, project.Name)
	}

	_, err = api.GetTimeEntriesForProjectBetween(project.ID, time.Now().AddDate(-10, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Error(err)
	}
}

func TestRealWorldGetTimeEntries(t *testing.T) {
	api := realWorldTestAPI(t)
	tes, err := api.GetTimeEntriesBetween(time.Now().AddDate(0, 0, -7), time.Now(), Defaults())
	if err != nil {
		t.Error(err)
	}
	if len(tes) < 1 {
		t.Error("GetTimeEntriesBetween() failed to find time entries in the last 7 days. Is this an empty Harvest account?")
	}
}

func TestRealWorldGetExpenses(t *testing.T) {
	api := realWorldTestAPI(t)

	expenses, err := api.GetExpenses(Defaults())
	if err != nil {
		t.Fatal(err)
	}

	if len(expenses) < 1 {
		t.Error("GetExpenses() returned no expenses. Are you testing with an empty Harvest account?")
	}

	expense, err := api.GetExpense(expenses[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}
	if expense.TotalCost == 0 {
		t.Error("Got an expense with no TotalCost.")
	}
}

func TestRealWorldGetExpenseCategories(t *testing.T) {
	api := realWorldTestAPI(t)

	categories, err := api.GetExpenseCategories(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(categories) < 1 {
		t.Error("GetExpenseCategories() returned no categories. Are you testing with an empty Harvest account?")
	}

	category, err := api.GetExpenseCategory(categories[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if category.Name == "" {
		t.Error("Category didn't have a name")
	}
}

func TestRealWorldGetInvoices(t *testing.T) {
	api := realWorldTestAPI(t)

	invoices, err := api.GetInvoices(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(invoices) < 1 {
		t.Error("GetInvoices() returned no invoices. Are you testing with an empty Harvest account?")
	}

	invoice, err := api.GetInvoice(invoices[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if invoice.Amount <= 0 {
		t.Error("Invoice should have an amount.")
	}

	messages, err := api.GetInvoiceMessages(invoice.ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(messages) < 1 {
		t.Skipf("Invoice %d had no messages, which might signal a problem, or that it doesn't have any messages.")
	}

}

func TestRealWorldGetUsers(t *testing.T) {
	api := realWorldTestAPI(t)

	users, err := api.GetUsers(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(users) < 1 {
		t.Error("GetUsers() returned no users. Are you testing with an empty Harvest account?")
	}

	user, err := api.GetUser(users[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if user.FirstName == "" {
		t.Error("User FirstName was blank")
	}
}

func TestRealWorldGetEstimates(t *testing.T) {
	api := realWorldTestAPI(t)

	estimates, err := api.GetEstimates(Defaults())
	if err != nil {
		t.Error(err)
	}

	if len(estimates) < 1 {
		t.Error("GetEstimates() returned no estimates. Are you testing with an empty Harvest account?")
	}

	estimate, err := api.GetEstimate(estimates[0].ID, Defaults())
	if err != nil {
		t.Error(err)
	}

	if estimate.Subject == "" {
		t.Error("Retrieved estimate was missing a subject.")
	}
}

/*
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
*/

func realWorldTestAPI(t *testing.T) *API {
	realworld := os.Getenv("HARVEST_REALWORLD")

	account := os.Getenv("HARVEST_ACCOUNT_ID")
	token := os.Getenv("HARVEST_TOKEN")

	if realworld == "true" && account != "" && token != "" {
		return NewTokenAPI(account, token)
	} else {
		t.Skipf("Skipping realworld tests because HARVEST_REALWORLD != true or HARVEST_ACCOUNT_ID / HARVEST_TOKEN not supplied as environment variables.")
		return nil
	}
}
