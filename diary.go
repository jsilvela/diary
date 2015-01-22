// Package diary defines the Record and Diary structures,
// defines their file persistence scheme (JSON)
// and hides these choices from calling modules.
package diary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Record struct {
	Event_time   time.Time
	Written_time time.Time
	Tags         []string
	Text         string
}

type Diary []*Record

// The Len, Swap and Less functions allow sorting
func (a Diary) Len() int           { return len(a) }
func (a Diary) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Diary) Less(i, j int) bool { return a[i].Event_time.Before(a[j].Event_time) }

func (r *Record) String() string {
	y, m, d := r.Event_time.Date()
	return fmt.Sprintf("time: %d-%d-%d\ntags: %s\ntext: %s\n",
		y, m, d, r.Tags, r.Text)
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

func (a *Diary) Add_entry(r *Record) {
	if r.Written_time.IsZero() {
		r.Written_time = time.Now()
	}
	*a = append(*a, r)
}

func (a Diary) Latest_happened() (r *Record) {
	var latest time.Time
	var rec *Record
	for _, e := range a {
		if latest.Before(e.Event_time) {
			latest, rec = e.Event_time, e
		}
	}
	return rec
}

func (a Diary) Latest_written() (r *Record) {
	var latest time.Time
	var rec *Record
	for _, e := range a {
		if latest.Before(e.Written_time) {
			latest, rec = e.Written_time, e
		}
	}
	return rec
}
