package diary

import (
	"testing"
	"time"
)

func Test_Add_entry(t *testing.T) {
	var d Diary
	(&d).Add_entry(&Record{Tags: []string{"hello", "there"}})
	if len(d) != 1 {
		t.Errorf("After adding one entry to empty diary, expected it to have 1 entry. "+
			"Got %d", len(d))
	}
	if d[0].Tags[0] != "hello" {
		t.Errorf("After storing single entry with tag 'hello', expected it to be the "+
			"first tag of the first entry. Was %s", d[0].Tags[0])
	}
}

func Test_Add_entry_respects_existing_Written_time(t *testing.T) {
	var d Diary
	r := Record{Tags: []string{"hello", "there"}}
	if !r.Written_time.IsZero() {
		t.Errorf("Unexpected non-default value for Written_time: %v", r.Written_time)
	}

	(&d).Add_entry(&r)
	t1 := r.Written_time
	if t1.IsZero() {
		t.Errorf("Unexpected default value for Written_time: %v", t1)
	}

	(&d).Add_entry(&r)
	t2 := r.Written_time
	if t2 != t1 {
		t.Errorf("Add_entry modified existing Add_entry date: %s from %s", t2, t1)
	}
}

func Test_Read_empty(t *testing.T) {
	d, err := Read("")
	if err == nil {
		t.Errorf("Should have errored when trying to read from empty file."+
			" Read this instead %v", d)
	}
}

func Test_Latest(t *testing.T) {
	var d Diary
	const base_time = "Jan 2 2006 15:04:05"
	t1, _ := time.Parse(base_time, "Jan 2 2006 15:04:05")
	t2, _ := time.Parse(base_time, "Jan 2 2016 15:04:05")
	t3, _ := time.Parse(base_time, "Jan 3 2006 15:04:05")

	(&d).Add_entry(&Record{
		Tags:         []string{"hello", "there"},
		Event_time:   t3,
		Written_time: t2})
	(&d).Add_entry(&Record{
		Tags:         []string{"bye", "there"},
		Event_time:   t1,
		Written_time: t1})

	latest_hp := d.Latest_happened()
	latest_wr := d.Latest_written()

	if latest_hp.Event_time != t3 {
		t.Errorf("Bad ordering. Latest happened should not have been: %v", latest_hp)
	}
	if latest_wr.Written_time != t2 {
		t.Errorf("Bad ordering. Latest written should not have been: %v", latest_wr)
	}
}
