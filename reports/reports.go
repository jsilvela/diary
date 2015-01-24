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

func Latest(d diary.Diary) map[string]*time.Time {

	tag_time := make(map[string]*time.Time)
	sort.Stable(sort.Reverse(d))
	for _, r := range d {
		for _, tag := range r.Tags {
			if tag_time[tag] == nil {
				tag_time[tag] = &r.Event_time
			}
		}
	}

	return tag_time
}

func Tags(d diary.Diary) []string {

	var tags []string
	m := Latest(d)
	for k := range m {
		tags = append(tags, k)
	}
	return tags
}

// CSV output
func Time_series(d diary.Diary) []string {
	file, err := os.Create("report.csv")
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)
	lines := make([]string, d.Len())
	sort.Stable(sort.Reverse(d))
	for i, r := range d {
		t := r.Event_time.Format("Mon 2 Jan 2006")
		tags := strings.Join(r.Tags, ",")
		writer.Write([]string{t, tags, r.Text})
		lines[i] = fmt.Sprintf("%s => %s\n", t, tags)
	}
	writer.Flush()
	return lines
}
