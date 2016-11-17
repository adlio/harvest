package harvest

import "testing"

func testClient(t *testing.T) *Client {
	a := testAPI()
	clientResponse := mockResponse("clients", "client-example.json")
	a.BaseURL = clientResponse.URL
	client, err := a.GetClient(3398386, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if client.Name != "Your Account" {
		t.Errorf("Incorrect Client Name '%s'", client.Name)
	}
	return client
}

func TestGetClient(t *testing.T) {
	client := testClient(t)
	if client == nil {
		t.Fatal("testClient() returned nil instead of client")
	}
	if client.ID != 3398386 {
		t.Errorf("Incorrect client ID '%v'", client.ID)
	}
}

func TestGetClients(t *testing.T) {
	a := testAPI()
	clientResponse := mockResponse("clients", "clients-example.json")
	a.BaseURL = clientResponse.URL
	clients, err := a.GetClients(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(clients) != 2 {
		t.Errorf("Incorrect number of clients '%v'", len(clients))
	}
	if clients[0].Name != "Your Account" {
		t.Errorf("Incorrect clients name '%s'", clients[0].Name)
	}
	if clients[1].Name != "Another Account" {
		t.Errorf("Incorrect client ID '%v'", clients[1].Name)
	}
}
