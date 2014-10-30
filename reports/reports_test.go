package reports

import (
	"github.com/jsilvela/diary"
	"testing"
	"time"
)

func Test_Latest(t *testing.T) {
	const short_form = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(short_form, "2014-07-17")
	t2, _ := time.Parse(short_form, "2014-07-30")
	(&d).Add_entry(&diary.Record{
		Tags:       []string{"A", "B"},
		Event_time: t1,
		Text:       "Blah"})

	(&d).Add_entry(&diary.Record{
		Tags:       []string{"C", "B"},
		Event_time: t2,
		Text:       "Bleh"})

	lt := Latest(d)
	if *lt["B"] != t2 {
		t.Errorf("Latest entry of B incorrect, was %s", lt["B"])
	}

	if len(lt) != 3 {
		t.Error("Wrong number of tags counted. Should be 3. %q", lt)
	}
}

func Test_Tags(t *testing.T) {
	const short_form = "2006-01-02"

	var d diary.Diary
	t1, _ := time.Parse(short_form, "2014-07-17")
	t2, _ := time.Parse(short_form, "2014-07-30")
	(&d).Add_entry(&diary.Record{
		Tags:       []string{"A", "B"},
		Event_time: t1,
		Text:       "Blah"})

	(&d).Add_entry(&diary.Record{
		Tags:       []string{"C", "B"},
		Event_time: t2,
		Text:       "Bleh"})

	tags := Tags(d)

	if len(tags) != 3 {
		t.Errorf("Should see three tags. Got: %s", tags)
	}
}
