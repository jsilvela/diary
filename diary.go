// Package diary defines the Record and Diary structures,
// defines their file persistence scheme (JSON)
// and hides these choices from calling modules.
package diary

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"time"
)

// Record represents an entry in the diary
type Record struct {
	EventTime   time.Time `json:"eventTime"`
	WrittenTime time.Time `json:"writtenTime"`
	Tags        []string  `json:"tags"`
	Text        string    `json:"text"`
}

// Diary represents a diary
type Diary []Record

// The Len, Swap and Less functions allow sorting
func (a Diary) Len() int           { return len(a) }
func (a Diary) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a Diary) Less(i, j int) bool { return a[i].EventTime.Before(a[j].EventTime) }

func (r Record) String() string {
	y, m, d := r.EventTime.Date()
	return fmt.Sprintf("time: %d-%d-%d\ntags: %s\ntext: %s\n",
		y, m, d, r.Tags, r.Text)
}

// Write diary onto file
func Write(writer io.Writer, d Diary) error {
	mar, err := json.MarshalIndent(d, "", "\t")
	if err != nil {
		return err
	}
	_, e := writer.Write(mar)
	return e
}

// Read diary from file
func Read(reader io.Reader) (Diary, error) {
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var reqs Diary
	err = json.Unmarshal(bytes, &reqs)
	if err != nil {
		return nil, err
	}
	return reqs, nil
}

// AddEntry adds a new record to the diary
func (a *Diary) AddEntry(r Record) {
	rcopy := r
	if rcopy.WrittenTime.IsZero() {
		rcopy.WrittenTime = time.Now()
		if rcopy.WrittenTime.IsZero() {
			panic("Whhhaaa?")
		}
	}
	*a = append(*a, rcopy)
}

// LatestHappened gets the record with the latest event time
func (a Diary) LatestHappened() (r Record) {
	var latest Record
	for _, e := range a {
		if latest.EventTime.Before(e.EventTime) {
			latest = e
		}
	}
	return latest
}

// LatestWritten gets event written last
func (a Diary) LatestWritten() (r Record) {
	var latest Record
	for _, e := range a {
		if latest.WrittenTime.Before(e.WrittenTime) {
			latest = e
		}
	}
	return latest
}
