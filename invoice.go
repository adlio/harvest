package harvest

import "time"

type Invoice struct {
	ID                 int64     `json:"id"`
	ClientID           int64     `json:"client_id"`
	PeriodStart        Date      `json:"period_start"`
	PeriodEnd          Date      `json:"period_end"`
	Number             string    `json:"number"`
	IssuedAt           Date      `json:"issued_at"`
	DueAt              Date      `json:"due_at"`
	Amount             float64   `json:"amount"`
	Currency           string    `json:"currency"`
	State              string    `json:"state"`
	Notes              string    `json:"notes"`
	PurchaseOrder      string    `json:"purchase_order"`
	DueAmount          int64     `json:"due_amount"`
	DueAtHumanFormat   string    `json:"due_at_human_format"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Tax                float64   `json:"tax"`
	TaxAmount          float64   `json:"tax_amount"`
	Subject            string    `json:"subject"`
	RecurringInvoiceID int64     `json:"recurring_invoice_id"`
	Tax2               float64   `json:"tax2"`
	Tax2Amount         float64   `json:"tax2_amount"`
	ClientKey          string    `json:"client_key"`
	EstimateID         int64     `json:"estimate_id"`
	Discount           float64   `json:"discount"`
	DiscountAmount     float64   `json:"discount_amount"`
	RetainerID         int64     `json:"retainer_id"`
	CreatedByID        int64     `json:"created_by_id"`
	CSVLineItems       string    `json:"csv_line_items"`
}
