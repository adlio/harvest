package harvest

import (
	"fmt"
	"time"
)

type InvoiceMessageResponse struct {
	InvoiceMessage *InvoiceMessage `json:"message"`
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

func (a *API) GetInvoiceMessages(invoiceID int64, args Arguments) (invoicemessages []*InvoiceMessage, err error) {
	invoiceMessagesResponse := make([]*InvoiceMessageResponse, 0)
	path := fmt.Sprintf("/invoices/%v/messages", invoiceID)
	err = a.Get(path, args, &invoiceMessagesResponse)
	for _, m := range invoiceMessagesResponse {
		invoicemessages = append(invoicemessages, m.InvoiceMessage)
	}
	return invoicemessages, err
}

func (a *API) GetInvoiceMessage(invoiceID int64, messageID int64, args Arguments) (invoicemessage *InvoiceMessage, err error) {
	invoiceMessageResponse := InvoiceMessageResponse{}
	path := fmt.Sprintf("/invoices/%v/messages/%v", invoiceID, messageID)
	err = a.Get(path, args, &invoiceMessageResponse)
	return invoiceMessageResponse.InvoiceMessage, err
}
