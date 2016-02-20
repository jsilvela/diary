package reports_test

import (
	"testing"
	"time"

	"github.com/jsilvela/diary"
	"github.com/jsilvela/diary/reports"
)

func Test_Latest(t *testing.T) {
	const shortForm = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(shortForm, "2014-07-17")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(diary.Record{
		Tags:      []string{"A", "B"},
		EventTime: t1,
		Text:      "Blah"})

	(&d).AddEntry(diary.Record{
		Tags:      []string{"C", "B"},
		EventTime: t2,
		Text:      "Bleh"})

	lt := reports.Latest(d)
	if lt["B"] != t2 {
		t.Errorf("reports.Latest entry of B incorrect, was %v", lt["B"])
	}

	if len(lt) != 3 {
		t.Errorf("Wrong number of tags counted. Should be 3. %v", lt)
	}
}

func TestTags(t *testing.T) {
	const shortForm = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(shortForm, "2014-07-17")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(diary.Record{
		Tags:      []string{"A", "B"},
		EventTime: t1,
		Text:      "Blah"})

	(&d).AddEntry(diary.Record{
		Tags:      []string{"C", "B"},
		EventTime: t2,
		Text:      "Bleh"})

	tags := reports.Tags(d)

	if len(tags) != 3 {
		t.Errorf("Should see three tags. Got: %s", tags)
	}
}

func TestTime_series(t *testing.T) {
	const shortForm = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(shortForm, "2014-07-17")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(diary.Record{
		Tags:      []string{"A", "B"},
		EventTime: t1,
		Text:      "Blah"})

	(&d).AddEntry(diary.Record{
		Tags:      []string{"C", "B"},
		EventTime: t2,
		Text:      "Bleh"})

	ts := reports.TimeSeries(d)
	if len(ts) != 2 {
		t.Errorf("Wrong number of tags counted. Should be 2. %v", ts)
	}
}
