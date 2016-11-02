package harvest

import (
	"time"
)

type Client struct {
	ID                      int64
	Name                    string
	CreatedAt               time.Time
	UpdatedAt               time.Time
	HighriseID              int64
	CacheVersion            int64
	Currency                string
	Details                 string
	DefaultInvoiceTimeframe string
	LastInvoiceKind         string
}
