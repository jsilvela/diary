package filters

import (
	"github.com/jsilvela/diary"
	"testing"
	"time"
)

func TestByWeek(t *testing.T) {
	var d diary.Diary
	const shortForm = "2006-01-02"

	t1, _ := time.Parse(shortForm, "2014-07-20")
	(&d).AddEntry(&diary.Record{EventTime: t1, Tags: []string{"testing"}})
	if len(*ByWeek(d)) > 0 {
		t.Errorf("Expected an empty diary when filtering by week on a diary with an old entry. Got %v", ByWeek(d))
	}

	t2, err := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(&diary.Record{EventTime: t2, Text: "My name is Coyote", Tags: []string{"testing"}})
	if len(*ByWeek(d)) != 1 {
		t.Errorf("Expected a diary with one entry when filtering by week on a diary with an fresh entry. Err %s, Got %v", err, d)
	}
}

func TestByTag(t *testing.T) {
	var d diary.Diary
	const shortForm = "2006-01-02"

	t1, _ := time.Parse(shortForm, "2014-07-20")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(&diary.Record{EventTime: t1, Tags: []string{"testing"}})
	(&d).AddEntry(&diary.Record{EventTime: t1, Tags: []string{"B"}})
	(&d).AddEntry(&diary.Record{EventTime: t2, Tags: []string{"B", "C"}})

	if len(*ByTag(d, "B")) != 2 {
		t.Errorf("Baaad")
	}
	if len(*ByTag(d, "C")) != 1 {
		t.Errorf("woorse")
	}
}
