package harvest

import (
	"encoding/json"
	"time"
)

type Date struct {
	time.Time
}

func (d *Date) Matches(t time.Time) bool {
	return d.Format("2006-01-02") == t.Format("2006-01-02")
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if d == nil {
		return []byte("null"), nil
	} else if d.IsZero() {
		return []byte("null"), nil
	} else {
		return json.Marshal(d.Format("2006-01-02"))
	}
}

func (d *Date) UnmarshalJSON(b []byte) (err error) {
	var src string
	if err = json.Unmarshal(b, &src); err == nil {
		d.ScanString(src)
	}
	return err
}

func (d *Date) ScanString(s string) {
	t, _ := time.Parse("2006-01-02", s)
	*d = Date{t}
}
