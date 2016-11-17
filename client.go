package harvest

import (
	"fmt"
	"time"
)

type ClientResponse struct {
	Client *Client `json:"client"`
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
	clientResponse := ClientResponse{}
	path := fmt.Sprintf("/clients/%v", clientID)
	err = a.Get(path, args, &clientResponse)
	return clientResponse.Client, err
}

func (a *API) GetClients(args Arguments) (clients []*Client, err error) {
	clientsResponse := make([]*ClientResponse, 0)
	path := fmt.Sprintf("/clients")
	err = a.Get(path, args, &clientsResponse)
	for _, cr := range clientsResponse {
		clients = append(clients, cr.Client)
	}
	return clients, err
}
