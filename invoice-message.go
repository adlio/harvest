package harvest

import "time"

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
