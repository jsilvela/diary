package filters

import (
	"github.com/jsilvela/diary"
	"time"
	"testing"
)

func TestWeek(t *testing.T) {
	var d diary.Diary
	const shortForm = "2006-01-02"

	t1, _ := time.Parse(shortForm, "2014-07-20")
	(&d).AddEntry(&diary.Record{EventTime: t1, Tags: []string{"testing"}})
	if len(*ByWeek(d)) > 0 {
		t.Errorf("Expected an empty diary when filtering by week on a diary with an old entry. Got %v", ByWeek(d))
	}

	t2, err:= time.Parse(shortForm, "2014-07-30")
	(&d).AddEntry(&diary.Record{EventTime: t2, Text: "My name is Coyote", Tags: []string{"testing"}})
	if len(*ByWeek(d)) != 1 {
		t.Errorf("Expected a diary with one entry when filtering by week on a diary with an fresh entry. Err %s, Got %v", err, d)
	}
}

	