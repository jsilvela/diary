// Package diary defines the Record and Diary structures, defines their file persistence scheme (JSON)
// and hides these choices from calling modules.
package diary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

type Record struct {
	EventTime   time.Time
	WrittenTime time.Time
	Tags        []string
	Text        string
}

type Diary []*Record

// The Len, Swap and Less functions allow sorting
func (a Diary) Len() int           { return len(a) }
func (a Diary) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Diary) Less(i, j int) bool { return a[i].EventTime.Before(a[j].EventTime) }

func (r *Record) String() string {
	y, m, d := r.EventTime.Date()
	return fmt.Sprintf("time: %d-%d-%d\ntags: %s\ntext: %s\n", y, m, d, r.Tags, r.Text)
}

func Write(filename string, d Diary) error {
	mar, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return err
	}
	e := ioutil.WriteFile(filename, mar, 0644)
	return e
}

func Read(filename string) (*Diary, error) {
	bytes, errfile := ioutil.ReadFile(filename)
	if errfile != nil {
		return nil, errfile
	}
	var reqs Diary
	err := json.Unmarshal(bytes, &reqs)
	if err != nil {
		return nil, err
	}
	return &reqs, nil
}

func (a *Diary) AddEntry(r *Record) {
	if r.WrittenTime.IsZero() {
		r.WrittenTime = time.Now()
		log.Println("WARN: modifying WrittenTime of entry")
	}
	*a = append(*a, r)
}

func (a Diary) LatestHappened() (r *Record) {
	var latest time.Time
	var rec *Record
	for _, e := range a {
		if latest.Before(e.EventTime) {
			latest, rec = e.EventTime, e
		}
	}
	return rec
}

func (a Diary) LatestWritten() (r *Record) {
	var latest time.Time
	var rec *Record
	for _, e := range a {
		if latest.Before(e.WrittenTime) {
			latest, rec = e.WrittenTime, e
		}
	}
	return rec
}
