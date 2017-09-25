package harvest

import (
	"fmt"
	"time"
)

type ClientsResponse struct {
	PagedResponse
	Clients []*Client `json:"clients"`
}
type ClientStub struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Client struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	IsActive  bool      `json:"is_active"`
	Address   string    `json:"address"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Currency  string    `json:"currency"`
}

func (a *API) GetClient(clientID int64, args Arguments) (client *Client, err error) {
	c := Client{}
	path := fmt.Sprintf("/clients/%v", clientID)
	err = a.Get(path, args, &c)
	return &c, err
}

func (a *API) GetClients(args Arguments) (clients []*Client, err error) {
	clients = make([]*Client, 0)
	clientsResponse := ClientsResponse{}
	path := fmt.Sprintf("/clients")
	err = a.GetPaginated(path, args, &clientsResponse, func() {
		for _, c := range clientsResponse.Clients {
			clients = append(clients, c)
		}
		clientsResponse.Clients = make([]*Client, 0)
	})
	return clients, err
}
