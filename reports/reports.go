package reports

import (
	"github.com/jsilvela/diary"
	"sort"
	"time"
)

func Latest(d diary.Diary) map[string]*time.Time {

	tagTime := make(map[string]*time.Time)
	sort.Stable(sort.Reverse(d))
	for _, r := range(d) {
		for _, tag := range(r.Tags) {
			if tagTime[tag] == nil {
				tagTime[tag] = &r.EventTime
			}
		}
	}
	
	return tagTime
}
	