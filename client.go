package harvest

import (
	"fmt"
	"time"
)

type ClientsResponse struct {
	PagedResponse
	Clients []*Client `json:"clients"`
}

type Client struct {
	ID                      int64     `json:"id"`
	Name                    string    `json:"name"`
	Active                  bool      `json:"active"`
	CreatedAt               time.Time `json:"created_at"`
	UpdatedAt               time.Time `json:"updated_at"`
	HighriseID              int64     `json:"highrise_id"`
	CacheVersion            int64     `json:"cache_version"`
	Currency                string    `json:"currency"`
	CurrencySymbol          string    `json:"currency_symbol"`
	Details                 string    `json:"details"`
	DefaultInvoiceTimeframe string    `json:"default_invoice_timeframe"`
	LastInvoiceKind         string    `json:"last_invoice_kind"`
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
