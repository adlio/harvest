package harvest

import (
	"fmt"
	"strconv"
	"time"
)

type InvoiceResponse struct {
	Invoice *Invoice `json:"invoices"`
}

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
	DueAmount          float64   `json:"due_amount"`
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

func (a *API) GetInvoice(invoiceID int64, args Arguments) (invoice *Invoice, err error) {
	invoiceResponse := InvoiceResponse{}
	path := fmt.Sprintf("/invoices/%d", invoiceID)
	err = a.Get(path, args, &invoiceResponse)
	return invoiceResponse.Invoice, err
}

func (a *API) GetInvoices(args Arguments) (invoices []*Invoice, err error) {
	invoices = make([]*Invoice, 0)
	invoicesResponse := make([]*InvoiceResponse, 0)

	path := fmt.Sprintf("/invoices")
	singlePage := false
	page := 1

	// If a "page" argument is provided, just get that single
	// page. Otherwise, we're going to iterate over all pages.
	if strPage := args["page"]; strPage != "" {
		if pageArg, _ := strconv.Atoi(strPage); pageArg > 0 {
			page = pageArg
			singlePage = true
		}
	}

	args["page"] = fmt.Sprintf("%d", page)
	err = a.Get(path, args, &invoicesResponse)
	for _, ir := range invoicesResponse {
		invoices = append(invoices, ir.Invoice)
	}

	if !singlePage && len(invoicesResponse) > 0 {
		page = 2
		moreInvoices := true

		// Loop over additional pages
		for moreInvoices == true {

			// Get the next page
			args["page"] = fmt.Sprintf("%d", page)
			err = a.Get(path, args, &invoicesResponse)
			if err != nil {
				return
			}

			if len(invoicesResponse) > 0 {
				for _, ir := range invoicesResponse {
					invoices = append(invoices, ir.Invoice)
				}
			} else {
				// Stop the loop
				moreInvoices = false
			}
			page += 1
		}
	}

	return
}
