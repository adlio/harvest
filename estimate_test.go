package harvest

import "testing"

func TestGetEstimate(t *testing.T) {
	estimate := testEstimate(t)

	if estimate.ID != 1439818 {
		t.Errorf("Incorrect Estimate ID '%v'", estimate.ID)
	}
	if estimate.Subject != "Online Store - Phase 2" {
		t.Errorf("Incorrect Subject '%s'", estimate.Subject)
	}
	if *estimate.Discount != 10.0 {
		t.Errorf("Expected Discount of 10.0. Got %0.1f", estimate.Discount)
	}
	if estimate.Amount != 9630.0 {
		t.Errorf("Expected Amount of 9630.00. Got %0.2f", estimate.Amount)
	}
}

func testEstimate(t *testing.T) *Estimate {
	a := testAPI()
	estimateResponse := mockResponse("estimates", "estimate-example.json")
	a.BaseURL = estimateResponse.URL
	estimate, err := a.GetEstimate(1439818, Defaults())
	if err != nil {
		t.Fatal(err)
	}
	if estimate == nil {
		t.Error("GetEstimate() failed.")
	}
	return estimate
}
