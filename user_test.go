package harvest

import "testing"

func testUser(t *testing.T) *User {
	a := testAPI()
	userResponse := mockResponse("users", "user-example.json")
	a.BaseURL = userResponse.URL
	user, err := a.GetUser(508343, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if user.Email != "user@example.com" {
		t.Errorf("Incorrect Email '%s'", user.Email)
	}
	if user.FirstName != "Harvest" {
		t.Errorf("Incorrect First Name '%s'", user.FirstName)
	}
	return user
}

func TestGetUser(t *testing.T) {
	user := testUser(t)
	if user == nil {
		t.Fatal("testUser() returned nil instead of user")
	}
	if user.ID != 508343 {
		t.Errorf("Incorrect user ID '%v'", user.ID)
	}
}

func TestGetUsers(t *testing.T) {
	a := testAPI()
	userResponse := mockResponse("users", "users-example.json")
	a.BaseURL = userResponse.URL
	users, err := a.GetUsers(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(users) != 2 {
		t.Errorf("Incorrect number of users '%v'", len(users))
	}
	if users[1].LastName != "User" {
		t.Errorf("Incorrect Last Name '%s'", users[1].LastName)
	}
	if users[0].Timezone != "Eastern Time (US & Canada)" {
		t.Errorf("Incorrect Time Zone '%s'", users[0].Timezone)
	}
}
