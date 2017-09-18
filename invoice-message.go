package harvest

import (
	"fmt"
	"time"
)

type InvoiceMessagesResponse struct {
	InvoiceMessages []*InvoiceMessage `json:"invoice_messages"`
	PerPage         int64             `json:"per_page"`
	TotalPages      int64             `json:"total_pages"`
	TotalEntries    int64             `json:"total_entries"`
	NextPage        *int64            `json:"next_page"`
	PreviousPage    *int64            `json:"previous_page"`
	Page            int64             `json:"page"`
}

type InvoiceMessage struct {
	ID                int64     `json:"id"`
	InvoiceID         int64     `json:"invoice_id"`
	SendMeACopy       bool      `json:"send_me_a_copy"`
	Body              string    `json:"body"`
	CreatedAt         time.Time `json:"created_at"`
	SentBy            string    `json:"sent_by"`
	SentByEmail       string    `json:"sent_by_email"`
	ThankYou          bool      `json:"thank_you"`
	Subject           string    `json:"subject"`
	IncludePayPalLink bool      `json:"include_pay_pal_link"`
	UpdatedAt         time.Time `json:"updated_at"`
	SentFromEmail     string    `json:"sent_from_email"`
	SentFrom          string    `json:"sent_from"`
	SendReminderOn    Date      `json:"send_reminder_on"`
	FullRecipientList string    `json:"full_recipient_list"`
}

func (a *API) GetInvoiceMessages(invoiceID int64, args Arguments) (invoiceMessages []*InvoiceMessage, err error) {
	invoiceMessagesResponse := InvoiceMessagesResponse{}
	path := fmt.Sprintf("/invoices/%v/messages", invoiceID)
	err = a.Get(path, args, &invoiceMessagesResponse)
	return invoiceMessagesResponse.InvoiceMessages, err
}

func (a *API) GetInvoiceMessage(invoiceID int64, messageID int64, args Arguments) (invoiceMessage *InvoiceMessage, err error) {
	invoiceMessage = &InvoiceMessage{}
	path := fmt.Sprintf("/invoices/%v/messages/%v", invoiceID, messageID)
	err = a.Get(path, args, invoiceMessage)
	return invoiceMessage, err
}
