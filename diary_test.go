package diary

import (
	"testing"
	"time"
)

func TestAddEntry(t *testing.T) {
	var d Diary
	(&d).AddEntry(&Record{Tags: []string{"hello", "there"}})
	if len(d) != 1 {
		t.Errorf("After adding one entry to empty diary, expected it to have 1 entry. Got %d", len(d))
	}
	if d[0].Tags[0] != "hello" {
		t.Errorf("After storing single entry with tag 'hello', expected it to be the first tag of the first entry. Was %s", d[0].Tags[0])
	}
}

func TestAddEntryRespectsExistingWrittenTime(t *testing.T) {
	var d Diary
	r := Record{Tags: []string{"hello", "there"}}
	if !r.WrittenTime.IsZero() {
		t.Errorf("Unexpected non-default value for WrittenTime: %v", r.WrittenTime)
	}

	(&d).AddEntry(&r)
	t1 := r.WrittenTime
	if t1.IsZero() {
		t.Errorf("Unexpected default value for WrittenTime: %v", t1)
	}

	(&d).AddEntry(&r)
	t2 := r.WrittenTime
	if t2 != t1 {
		t.Errorf("AddEntry modified existing WriteEntry date: %s from %s", t2, t1)
	}
}

func TestReadEmpty(t *testing.T) {
	d, err := Read("")
	if err == nil {
		t.Errorf("Should have errored when trying to read from empty file. Read this instead %v", d)
	}
}

func TestLatest(t *testing.T) {
	var d Diary
	t1, _ := time.Parse("Jan 2 2006 15:04:05", "Jan 2 2006 15:04:05")
	t2, _ := time.Parse("Jan 2 2006 15:04:05", "Jan 2 2016 15:04:05")
	t3, _ := time.Parse("Jan 2 2006 15:04:05", "Jan 3 2006 15:04:05")

	(&d).AddEntry(&Record{
		Tags:        []string{"hello", "there"},
		EventTime:   t3,
		WrittenTime: t2})
	(&d).AddEntry(&Record{
		Tags:        []string{"bye", "there"},
		EventTime:   t1,
		WrittenTime: t1})

	latestH := d.LatestHappened()
	latestW := d.LatestWritten()

	if latestH.EventTime != t3 {
		t.Errorf("Bad ordering. Latest happened should not have been: %v", latestH)
	}
	if latestW.WrittenTime != t2 {
		t.Errorf("Bad ordering. Latest written should not have been: %v", latestW)
	}
}
