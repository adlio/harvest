package harvest

import (
	"testing"
	"time"
)

func TestGetTimeEntriesUpdatedSince(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("time-entries", "project-example.json")
	a.BaseURL = projectResponse.URL
	sinceDate, err := time.Parse("2006-01-02", "2017-01-01")
	timeEntries, err := a.GetTimeEntriesUpdatedSince(sinceDate, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	validateTimeEntriesFromProjectExample(timeEntries, t)
}

func TestGetTimeEntriesBetween(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("time-entries", "project-example.json")
	a.BaseURL = projectResponse.URL
	timeEntries, err := a.GetTimeEntriesBetween(time.Now().AddDate(-1, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Fatal(err)
	}
	validateTimeEntriesFromProjectExample(timeEntries, t)
}

func TestGetTimeEntriesForProjectBetween(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("time-entries", "project-example.json")
	a.BaseURL = projectResponse.URL
	timeEntries, err := a.GetTimeEntriesForProjectBetween(3, time.Now().AddDate(-1, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Fatal(err)
	}
	validateTimeEntriesFromProjectExample(timeEntries, t)
}

func validateTimeEntriesFromProjectExample(timeEntries []*TimeEntry, t *testing.T) {
	if len(timeEntries) != 3 {
		t.Errorf("Incorrect number of entries %d", len(timeEntries))
	}
	if timeEntries[0].Project.ID != 3 {
		t.Errorf("Expected 3 for ProjectID, got %d", timeEntries[0].Project.ID)
	}
	if timeEntries[0].User.ID != 1 {
		t.Errorf("Expected 1 for UserID, got %d", timeEntries[0].User.ID)
	}
	if timeEntries[0].User.Name != "The First User" {
		t.Errorf("Expected 'The First User' for User name, got '%s'", timeEntries[0].User.Name)
	}
}

func TestGetTimeEntriesForUserBetween(t *testing.T) {
	a := testAPI()
	usersResponse := mockResponse("time-entries", "user-example.json")
	a.BaseURL = usersResponse.URL
	timeEntries, err := a.GetTimeEntriesForUserBetween(1, time.Now().AddDate(-1, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(timeEntries) != 2 {
		t.Errorf("Incorrect number of entries %d", len(timeEntries))
	}
	if timeEntries[0].Project.ID != 2 {
		t.Errorf("Expected ProjectID 2, got %d", timeEntries[0].Project.ID)
	}
	if timeEntries[0].Notes != "First task" {
		t.Errorf("Expected 'First task', got '%s'.", timeEntries[0].Notes)
	}
	if timeEntries[1].Notes != "" {
		t.Errorf("Expected blank Notes, got %v", timeEntries[1].Notes)
	}
}
