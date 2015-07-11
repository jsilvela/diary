package diary_test

import (
	"github.com/jsilvela/diary"
	"testing"
	"time"
)

func Test_AddEntry(t *testing.T) {
	var d diary.Diary
	(&d).AddEntry(diary.Record{Tags: []string{"hello", "there"}})
	if len(d) != 1 {
		t.Errorf("After adding one entry to empty diary, expected it to have 1 entry. "+
			"Got %d", len(d))
	}
	if d[0].Tags[0] != "hello" {
		t.Errorf("After storing single entry with tag 'hello', expected it to be the "+
			"first tag of the first entry. Was %s", d[0].Tags[0])
	}
}

func Test_AddEntry_respectsExisting_Written_time(t *testing.T) {
	var d diary.Diary
	r := diary.Record{Tags: []string{"hello", "there"}}
	if !r.WrittenTime.IsZero() {
		t.Errorf("Unexpected non-default value for WrittenTime: %v", r.WrittenTime)
	}

	(&d).AddEntry(r)
	t1 := r.WrittenTime
	if !t1.IsZero() {
		t.Error("AddEntry modified the record argument")
	}
	el1 := d[0]
	if el1.WrittenTime.IsZero() {
		t.Error("AddEntry didn't write the TimeWritten timestamp")
	}

	r2 := diary.Record{WrittenTime: time.Now(), EventTime: time.Now()}
	(&d).AddEntry(r2)
	el2 := d[1]
	if el2.WrittenTime != r2.WrittenTime {
		t.Errorf("AddEntry modified existing AddEntry date: %s from %s", r2, el2)
	}
}

func Test_ReadEmpty(t *testing.T) {
	d, err := diary.Read("")
	if err == nil {
		t.Errorf("Should have errored when trying to read from empty file."+
			" Read this instead %v", d)
	}
}

func Test_Latest(t *testing.T) {
	var d diary.Diary
	const baseTime = "Jan 2 2006 15:04:05"
	t1, _ := time.Parse(baseTime, "Jan 2 2006 15:04:05")
	t2, _ := time.Parse(baseTime, "Jan 2 2016 15:04:05")
	t3, _ := time.Parse(baseTime, "Jan 3 2006 15:04:05")

	(&d).AddEntry(diary.Record{
		Tags:        []string{"hello", "there"},
		EventTime:   t3,
		WrittenTime: t2})
	(&d).AddEntry(diary.Record{
		Tags:        []string{"bye", "there"},
		EventTime:   t1,
		WrittenTime: t1})

	latestHp := d.LatestHappened()
	latestWr := d.LatestWritten()

	if latestHp.EventTime != t3 {
		t.Errorf("Bad ordering. Latest happened should not have been: %v", latestHp)
	}
	if latestWr.WrittenTime != t2 {
		t.Errorf("Bad ordering. Latest written should not have been: %v", latestWr)
	}
}
