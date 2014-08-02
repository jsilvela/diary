// Package diary defines the Record and Diary structures, defines their file persistence scheme (JSON)
// and hides these choices from calling modules.
package diary

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
func (a Diary) Less(i, j int) bool { return a[i].WrittenTime.Before(a[j].WrittenTime) }

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
	r.WrittenTime = time.Now()
	*a = append(*a, r)
}
