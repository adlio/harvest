package harvest

import (
	"fmt"
	"time"
)

type ClientsResponse struct {
	Clients      []*Client `json:"clients"`
	PerPage      int64     `json:"per_page"`
	TotalPages   int64     `json:"total_pages"`
	TotalEntries int64     `json:"total_entries"`
	NextPage     *int64    `json:"next_page"`
	PreviousPage *int64    `json:"previous_page"`
	Page         int64     `json:"page"`
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
	clientsResponse := ClientsResponse{}
	path := fmt.Sprintf("/clients")
	err = a.Get(path, args, &clientsResponse)
	return clientsResponse.Clients, err
}
