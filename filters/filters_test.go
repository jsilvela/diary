package filters_test

import (
	"github.com/jsilvela/diary"
	"github.com/jsilvela/diary/filters"
	"testing"
	"time"
)

const (
	Day = time.Hour * 24
)

func Test_By_range(t *testing.T) {
	var d diary.Diary

	now := time.Now()
	last_week := now.Add(-1 * filters.Week)

	t1 := now.Add(-2 * filters.Week)
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"testing"}})

	if len(*filters.By_range(d, last_week, now)) > 0 {
		t.Errorf("Expected empty when filtering by week on a diary with an old entry."+
			" Got %v", filters.By_week(d))
	}

	t2 := now.Add(-2 * Day)
	(&d).Add_entry(&diary.Record{
		Event_time: t2,
		Text:       "My name is Coyote",
		Tags:       []string{"testing"}})
	if len(*filters.By_range(d, last_week, now)) != 1 {
		t.Errorf("Expected one result when filtering by week on a "+
			"diary with a fresh entry. Got %v", *filters.By_week(d))
	}
}

func Test_By_tag(t *testing.T) {
	var d diary.Diary
	const shortForm = "2006-01-02"

	t1, _ := time.Parse(shortForm, "2014-07-20")
	t2, _ := time.Parse(shortForm, "2014-07-30")
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"testing"}})
	(&d).Add_entry(&diary.Record{Event_time: t1, Tags: []string{"B"}})
	(&d).Add_entry(&diary.Record{Event_time: t2, Tags: []string{"B", "C"}})

	if len(*filters.By_tag(d, "B")) != 2 {
		t.Errorf("Baaad")
	}
	if len(*filters.By_tag(d, "C")) != 1 {
		t.Errorf("woorse")
	}
}
