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

	var dNull *Date
	jsonBytes, err = dNull.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	if string(jsonBytes) != "null" {
		t.Errorf("Expected nil Date{} to marshal to JSON as 'null. Got '%s' instead.", string(jsonBytes))
	}
}

func TestDateMatchesTrue(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2017-03-01")
	t2 := t1.Add(time.Minute * time.Duration(30))
	d := Date{t1}
	if d.Matches(t2) != true {
		t.Errorf("Date.Matches() should be true as long as the date parts match")
	}
}

func TestDateMatchesFalse(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2017-03-01")
	t2, _ := time.Parse("2006-01-02", "2017-03-02")
	d := Date{t1}
	if d.Matches(t2) != false {
		t.Errorf("Date.Matches() should be fale when date parts are different")
	}
}
