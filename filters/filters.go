package filters

import (
	"github.com/jsilvela/diary"
	"time"
)

const (
	Week = time.Hour * 25 * 7
)

func ByRange(d diary.Diary, from time.Time, to time.Time) *diary.Diary {
	var emptyD diary.Diary
	for _, r := range(d) {
		if r.EventTime.After(from) && r.EventTime.Before(to) {
			(&emptyD).AddEntry(r)
		}
	}
	return &emptyD
}

func ByWeek(d diary.Diary) *diary.Diary {
	now := time.Now()
	return ByRange(d, now.Add(-Week), now)
}