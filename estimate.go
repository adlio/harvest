package harvest

import (
	"fmt"
	"time"
)

type EstimatesResponse struct {
	PagedResponse
	Estimates []*Estimate `json:"estimates"`
}

type Estimate struct {
	ID             int64               `json:"id"`
	Subject        string              `json:"subject"`
	Client         ClientStub          `json:"client,omitempty"`
	Creator        UserStub            `json:"creator,omitempty"`
	ClientKey      string              `json:"client_key"`
	Number         string              `json:"number"`
	PurchaseOrder  string              `json:"purchase_order"`
	Amount         float64             `json:"amount"`
	Tax            *float64            `json:"tax"`
	TaxAmount      float64             `json:"tax_amount"`
	Tax2           *float64            `json:"tax2"`
	Tax2Amount     float64             `json:"tax2_amount"`
	Discount       *float64            `json:"discount"`
	DiscountAmount float64             `json:"discount_amount"`
	Notes          string              `json:"notes"`
	Currency       string              `json:"currency"`
	IssueDate      *Date               `json:"issue_date,omitempty"`
	SentAt         *time.Time          `json:"sent_at,omitempty"`
	LineItems      []*EstimateLineItem `json:"line_items"`
	AcceptedAt     *time.Time          `json:"accepted_at,omitempty"`
	DeclinedAt     *time.Time          `json:"declined_at,omitempty"`
	CreatedAt      time.Time           `json:"created_at,omitempty"`
	UpdatedAt      time.Time           `json:"updated_at,omitempty"`
}

type EstimateLineItem struct {
	ID          int64   `json:"id"`
	Kind        string  `json:"kind"`
	Description string  `json:"description"`
	Quantity    float64 `json:"quantity"`
	UnitPrice   float64 `json:"unit_price"`
	Amount      float64 `json:"amount"`
	Taxed       bool    `json:"taxed"`
	Taxed2      bool    `json:"taxed2"`
}

func (a *API) GetEstimate(estimateID int64, args Arguments) (estimate *Estimate, err error) {
	estimate = &Estimate{}
	path := fmt.Sprintf("/estimates/%d", estimateID)
	err = a.Get(path, args, estimate)
	return estimate, err
}

func (a *API) GetEstimates(args Arguments) (estimates []*Estimate, err error) {
	estimates = make([]*Estimate, 0)
	estimatesResponse := EstimatesResponse{}
	err = a.GetPaginated("/estimates", args, &estimatesResponse, func() {
		for _, e := range estimatesResponse.Estimates {
			estimates = append(estimates, e)
		}
		estimatesResponse.Estimates = make([]*Estimate, 0)
	})
	return estimates, err
}
