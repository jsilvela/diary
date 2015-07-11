package reports

import (
	"encoding/csv"
	"fmt"
	"github.com/jsilvela/diary"
	"log"
	"os"
	"sort"
	"strings"
	"time"
)

// Latest maps each tag to its latest occurrence
func Latest(d diary.Diary) map[string]time.Time {

	tagTime := make(map[string]time.Time)
	sort.Stable(sort.Reverse(d))
	for _, r := range d {
		for _, tag := range r.Tags {
			if tagTime[tag].IsZero() {
				tagTime[tag] = r.EventTime
			}
		}
	}

	return tagTime
}

// Tags finds all the Tags in the diary
func Tags(d diary.Diary) []string {

	var tags []string
	m := Latest(d)
	for k := range m {
		tags = append(tags, k)
	}
	return tags
}

// TimeSeries generates CSV output with the time series data
func TimeSeries(d diary.Diary) []string {
	file, err := os.Create("report.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	lines := make([]string, d.Len())
	sort.Stable(sort.Reverse(d))
	for i, r := range d {
		t := r.EventTime.Format("Mon 2 Jan 2006")
		tags := strings.Join(r.Tags, ",")
		writer.Write([]string{t, tags, r.Text})
		lines[i] = fmt.Sprintf("%s => %s\n", t, tags)
	}
	writer.Flush()
	return lines
}
