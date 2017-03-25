package harvest

import "testing"

func TestGetInvoice(t *testing.T) {
	a := testAPI()
	invoiceResponse := mockResponse("invoices", "invoice-example.json")
	a.BaseURL = invoiceResponse.URL
	invoice, err := a.GetInvoice(12286767, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if invoice == nil {
		t.Fatal("GetInvoice returned nil")
	}
	if invoice.ID != 12286767 {
		t.Errorf("Incorrect expense ID '%v'", invoice.ID)
	}
	if invoice.ClientID != 3781881 {
		t.Errorf("Incorrect client ID '%v'", invoice.ClientID)
	}
	if invoice.Notes != "Thank you in advance for your prompt payment, which is essential to our ability to best serve you." {
		t.Errorf("Incorrect Invoice Notes '%s'", invoice.Notes)
	}
}

func TestGetInvoices(t *testing.T) {
	a := testAPI()
	a.BaseURL = mockDynamicPathResponse().URL
	invoices, err := a.GetInvoices(Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if len(invoices) != 8 {
		t.Errorf("Incorrect number of invoices. Expected 8, got %d", len(invoices))
	}
	if invoices[0].ID != 7941658 {
		t.Errorf("Incorrect invoice ID '%v'", invoices[0].ID)
	}
}
