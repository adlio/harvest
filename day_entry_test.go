package harvest

import "testing"

func TestGetTodayEntry(t *testing.T) {
	a := testAPI()
	dayEntryResponse := mockResponse("dayentries", "today-example.json")
	a.BaseURL = dayEntryResponse.URL
	dayentries, err := a.GetTodayEntry(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(dayentries) != 2 {
		t.Errorf("Incorrect number of entries '%v'", len(dayentries))
	}
	if dayentries[0].ID != 538242480 {
		t.Errorf("Incorrect day entry ID '%v'", dayentries[0].ID)
	}
	if dayentries[1].UserID != 1420761 {
		t.Errorf("Incorrect UserID '%v'", dayentries[1].ID)
	}
}
