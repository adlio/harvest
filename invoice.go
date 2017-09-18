package harvest

import (
	"fmt"
	"strconv"
	"time"
)

type InvoicesResponse struct {
	Invoices     []*Invoice `json:"invoices"`
	PerPage      int64      `json:"per_page"`
	TotalPages   int64      `json:"total_pages"`
	TotalEntries int64      `json:"total_entries"`
	NextPage     *int64     `json:"next_page"`
	PreviousPage *int64     `json:"previous_page"`
	Page         int64      `json:"page"`
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
	invoice = &Invoice{}
	path := fmt.Sprintf("/invoices/%d", invoiceID)
	err = a.Get(path, args, &invoice)
	return invoice, err
}

func (a *API) GetInvoices(args Arguments) (invoices []*Invoice, err error) {
	invoices = make([]*Invoice, 0)
	invoicesResponse := InvoicesResponse{}

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
	for _, i := range invoicesResponse.Invoices {
		invoices = append(invoices, i)
	}

	if !singlePage && invoicesResponse.TotalPages > 1 {
		page = 2
		moreInvoices := true

		// Loop over additional pages
		for moreInvoices == true {

			// Get the next page
			args["page"] = fmt.Sprintf("%d", page)
			invoicesResponse.Invoices = make([]*Invoice, 0)
			err = a.Get(path, args, &invoicesResponse)
			if err != nil {
				return
			}

			if len(invoicesResponse.Invoices) > 0 {
				for _, i := range invoicesResponse.Invoices {
					invoices = append(invoices, i)
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
