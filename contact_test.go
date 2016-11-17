package harvest

import "testing"

func testContact(t *testing.T) *Contact {
	a := testAPI()
	contactResponse := mockResponse("contacts", "contact-example.json")
	a.BaseURL = contactResponse.URL
	contact, err := a.GetContact(2937808, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	return contact
}

func TestGetContact(t *testing.T) {
	contact := testContact(t)
	if contact == nil {
		t.Fatal("testContact() returned nil instead of contact")
	}
	if contact.ID != 2937808 {
		t.Errorf("Incorrect contact ID '%v'", contact.ID)
	}
}

func TestGetClientContacts(t *testing.T) {
	a := testAPI()
	contactResponse := mockResponse("contacts", "contacts-example.json")
	a.BaseURL = contactResponse.URL
	contacts, err := a.GetClientContacts(1661738, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(contacts) != 2 {
		t.Errorf("Incorrect number of contacts '%v'", len(contacts))
	}
	if contacts[0].LastName != "Contact" {
		t.Errorf("Incorrect Last Name '%s'", contacts[0].LastName)
	}
	if contacts[1].Email != "person@example.com" {
		t.Errorf("Incorrect Email '%s'", contacts[1].Email)
	}
}

func TestGetContacts(t *testing.T) {
	a := testAPI()
	contactResponse := mockResponse("contacts", "contacts-example.json")
	a.BaseURL = contactResponse.URL
	contacts, err := a.GetContacts(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(contacts) != 2 {
		t.Errorf("Incorrect number of contacts '%v'", len(contacts))
	}
	if contacts[0].LastName != "Contact" {
		t.Errorf("Incorrect Last Name '%s'", contacts[0].LastName)
	}
	if contacts[1].Email != "person@example.com" {
		t.Errorf("Incorrect Email '%s'", contacts[1].Email)
	}
}
