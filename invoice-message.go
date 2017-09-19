package harvest

import (
	"fmt"
	"time"
)

type InvoiceMessagesResponse struct {
	PagedResponse
	InvoiceMessages []*InvoiceMessage `json:"invoice_messages"`
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
	invoiceMessages = make([]*InvoiceMessage, 0)
	invoiceMessagesResponse := InvoiceMessagesResponse{}
	path := fmt.Sprintf("/invoices/%v/messages", invoiceID)
	err = a.GetPaginated(path, args, &invoiceMessagesResponse, func() {
		for _, im := range invoiceMessagesResponse.InvoiceMessages {
			invoiceMessages = append(invoiceMessages, im)
		}
		invoiceMessagesResponse.InvoiceMessages = make([]*InvoiceMessage, 0)
	})
	return invoiceMessages, err
}

func (a *API) GetInvoiceMessage(invoiceID int64, messageID int64, args Arguments) (invoiceMessage *InvoiceMessage, err error) {
	invoiceMessage = &InvoiceMessage{}
	path := fmt.Sprintf("/invoices/%v/messages/%v", invoiceID, messageID)
	err = a.Get(path, args, invoiceMessage)
	return invoiceMessage, err
}
