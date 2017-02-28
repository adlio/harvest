package harvest

import (
	"testing"
	"time"
)

func TestGetTodayEntries(t *testing.T) {
	a := testAPI()
	dayEntryResponse := mockResponse("day_entries", "today-example.json")
	a.BaseURL = dayEntryResponse.URL
	dayEntries, err := a.GetTodayEntries(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(dayEntries) != 2 {
		t.Errorf("Incorrect number of entries %d", len(dayEntries))
	}
	if dayEntries[0].ID != 538242480 {
		t.Errorf("Incorrect day entry ID '%v'", dayEntries[0].ID)
	}
	if dayEntries[1].UserID != 1420761 {
		t.Errorf("Incorrect UserID '%v'", dayEntries[1].ID)
	}
	if dayEntries[1].SpentAtRaw != "2016-11-15" {
		t.Errorf("Incorrect SpentAtRaw '%s'", dayEntries[1].SpentAtRaw)
	}
	if dayEntries[1].SpentAt.IsZero() {
		t.Error("dayEntries[1].SpentAt should be a populated time.Time")
	}
	if dayEntries[1].SpentAt.Format("2006-01-02") != "2016-11-15" {
		t.Errorf("Expected '2016-11-15'. Got '%s'.", dayEntries[1].SpentAt.Format("2006-01-01"))
	}
}

func TestGetEntriesForProjectBetween(t *testing.T) {
	a := testAPI()
	projectResponse := mockResponse("day_entries", "project-example.json")
	a.BaseURL = projectResponse.URL
	dayEntries, err := a.GetEntriesForProjectBetween(3, time.Now().AddDate(-1, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(dayEntries) != 3 {
		t.Errorf("Incorrect number of entries %d", len(dayEntries))
	}
	if dayEntries[0].ProjectID != 3 {
		t.Errorf("Expected 3 for ProjectID, got %d", dayEntries[0].ProjectID)
	}
}

func TestGetEntriesForUserBetween(t *testing.T) {
	a := testAPI()
	usersResponse := mockResponse("day_entries", "user-example.json")
	a.BaseURL = usersResponse.URL
	dayEntries, err := a.GetEntriesForUserBetween(1, time.Now().AddDate(-1, 0, 0), time.Now(), Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(dayEntries) != 2 {
		t.Errorf("Incorrect number of entries %d", len(dayEntries))
	}
	if dayEntries[0].ProjectID != 2 {
		t.Errorf("Expected ProjectID 2, got %d", dayEntries[0].ProjectID)
	}
	if dayEntries[0].Notes != "First task" {
		t.Errorf("Expected 'First task', got '%s'.", dayEntries[0].Notes)
	}
	if dayEntries[1].Notes != "" {
		t.Errorf("Expected blank Notes, got %v", dayEntries[1].Notes)
	}
}
