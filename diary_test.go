package diary_test

import (
	"bytes"
	"github.com/jsilvela/diary"
	"strings"
	"testing"
	"time"
)

func TestAddEntry(t *testing.T) {
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

func TestAddEntry_respects_existing_WrittenTime(t *testing.T) {
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

func TestReadEmpty(t *testing.T) {
	empty := bytes.NewBufferString("")
	d, err := diary.Read(empty)
	if err == nil {
		t.Errorf("Should have errored when trying to read from empty file."+
			" Read this instead %v", d)
	}
}

func TestRead(t *testing.T) {
	buf := bytes.NewBufferString(`[
	{
		"eventTime": "2014-09-21T00:00:00Z",
		"writtenTime": "2014-03-11T09:42:31.379879883+02:00",
		"tags": [
			"myTag"
		],
		"text": "myText."
	}]`)
	d, err := diary.Read(buf)
	if err != nil {
		t.Error(err)
	}
	if len(d) != 1 {
		t.Errorf("Should have found one entry, found %d", len(d))
	}
	if d[0].Text != "myText." || d[0].Tags[0] != "myTag" {
		t.Errorf("Unexpected value parsed for entry: %v", d[0])
	}
}

func TestWrite(t *testing.T) {
	var d diary.Diary
	const baseTime = "Jan 2 2006 15:04:05"
	t3, _ := time.Parse(baseTime, "Jan 3 2006 15:04:05")

	(&d).AddEntry(diary.Record{
		Tags:        []string{"hello", "there"},
		EventTime:   t3,
		WrittenTime: t3,
		Text:        "MyText."})
	var b []byte
	buf := bytes.NewBuffer(b)
	err := diary.Write(buf, d)
	if err != nil {
		t.Error(err)
	}
	expected := `[
        {
                "eventTime": "2006-01-03T15:04:05Z",
                "writtenTime": "2006-01-03T15:04:05Z",
                "tags": [
                        "hello",
                        "there"
                ],
                "text": "MyText."
        }
]`
	if strings.Replace(strings.Replace(buf.String(), " ", "", -1), "\t", "", -1) !=
		strings.Replace(expected, " ", "", -1) {
		t.Errorf("Unexpected output:\n%s\n%s",
			strings.Replace(strings.Replace(buf.String(), " ", "", -1), "\t", "", -1),
			strings.Replace(expected, " ", "", -1))
	}
}

func TestLatestTimestamps(t *testing.T) {
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
