package test

import (
	"encoding/json"
	"testing"
	"time"

	calendar "github.com/elastic/go-json-schema-generate/test/calendar_gen"
)

func TestCalendar(t *testing.T) {
	data := []byte(`{
	"date": "2022-01-02T12:00:00Z"
    }`)
	cal := &calendar.Calendar{}
	if err := json.Unmarshal(data, &cal); err != nil {
		t.Fatal(err)
	}
	ts := time.Date(2022, 1, 2, 12, 0, 0, 0, time.UTC)
	if !cal.Date.Equal(ts) {
		t.Errorf("expected date to be %v got %v", ts, cal.Date)
	}
}
