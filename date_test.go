package harvest

import (
	"testing"
	"time"
)

func TestDateMarshallJSON(t *testing.T) {
	time, _ := time.Parse("2006-01-02", "2017-03-01")
	d := Date{time}
	jsonBytes, err := d.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(jsonBytes) != "\"2017-03-01\"" {
		t.Errorf("Expected '\"2017-03-01\"'. Got '%s'", string(jsonBytes))
	}
}
