package harvest

import "testing"

func TestGetInvoiceMessages(t *testing.T) {
	a := testAPI()
	invoiceMessageResponse := mockResponse("invoicemessages", "invoicemessages-example.json")
	a.BaseURL = invoiceMessageResponse.URL
	invoicemessages, err := a.GetInvoiceMessages(6770075, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(invoicemessages) != 2 {
		t.Errorf("Incorrect number of invoice messages '%v'", len(invoicemessages))
	}
	if invoicemessages[0].ID != 12465017 {
		t.Errorf("Incorrect task assignment ID '%v'", invoicemessages[0].ID)
	}
	if invoicemessages[1].InvoiceID != 6770075 {
		t.Errorf("Incorrect InvoiceID '%v'", invoicemessages[1].InvoiceID)
	}
}

func TestGetInvoiceMessage(t *testing.T) {
	a := testAPI()
	invoiceMessageResponse := mockResponse("invoicemessages", "invoicemessage-example.json")
	a.BaseURL = invoiceMessageResponse.URL
	invoicemessage, err := a.GetInvoiceMessage(6770075, 12465017, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if invoicemessage.InvoiceID != 6770075 {
		t.Errorf("Incorrect invoice ID '%v'", invoicemessage.InvoiceID)
	}
}
