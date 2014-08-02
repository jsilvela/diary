package reports

import (
	"github.com/jsilvela/diary"
	"time"
	"testing"
)

func TestLatest(t * testing.T) {
	const shortForm = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(shortForm, "2014-07-17")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(&diary.Record{
		Tags: []string{"A", "B"},
		EventTime: t1,
		Text: "Blah"})
	
	(&d).AddEntry(&diary.Record{
		Tags: []string{"C", "B"},
		EventTime: t2,
		Text: "Bleh"})

	lt := Latest(d)
	if *lt["B"] != t2 {
		t.Errorf("Latest entry of B should be the latest, was %s", lt["B"])
	}
	
	if len(lt) != 3 {
		t.Error("Wrong number of tags counted. Should be 3. %q", lt)
	}
}