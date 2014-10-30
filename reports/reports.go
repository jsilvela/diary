package reports

import (
	"github.com/jsilvela/diary"
	"sort"
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
