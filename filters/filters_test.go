package filters

import (
	"github.com/jsilvela/diary"
	"testing"
	"time"
)

const (
	Day = time.Hour * 24
)

func TestBy_week(t *testing.T) {
	var d diary.Diary

	t1 := time.Now().Add(-2 * Week)
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"testing"}})
	if len(*By_week(d)) > 0 {
		t.Errorf("Expected empty when filtering by week on a diary with an old entry."+
			" Got %v", By_week(d))
	}

	t2 := time.Now().Add(-2 * Day)
	(&d).Add_entry(&diary.Record{
		Event_time: t2,
		Text:       "My name is Coyote",
		Tags:       []string{"testing"}})
	if len(*By_week(d)) != 1 {
		t.Errorf("Expected one result when filtering by week on a "+
			"diary with a fresh entry. Got %v", *By_week(d))
	}
}

func TestBy_tag(t *testing.T) {
	var d diary.Diary
	const shortForm = "2006-01-02"

	t1, _ := time.Parse(shortForm, "2014-07-20")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"testing"}})
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"B"}})
	(&d).Add_entry(&diary.Record{Event_time: t2, Tags: []string{"B", "C"}})

	if len(*By_tag(d, "B")) != 2 {
		t.Errorf("Baaad")
	}
	if len(*By_tag(d, "C")) != 1 {
		t.Errorf("woorse")
	}
}
