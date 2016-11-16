package harvest

import (
	"time"
)

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
