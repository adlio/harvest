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
}
