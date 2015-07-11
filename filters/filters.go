package filters

import (
	"github.com/jsilvela/diary"
	"time"
)

// Time constants
const (
	Week  = time.Hour * 24 * 7
	Month = time.Hour * 24 * 30
)

// ByRange gets the subset of the diary that falls in a time range
func ByRange(d diary.Diary, from time.Time, to time.Time) *diary.Diary {
	var dr diary.Diary
	for _, r := range d {
		if r.EventTime.After(from) && r.EventTime.Before(to) {
			(&dr).AddEntry(r)
		}
	}
	return &dr
}

// ByWeek gets events that fall in the last week
func ByWeek(d diary.Diary) *diary.Diary {
	now := time.Now()
	return ByRange(d, now.Add(-Week), now)
}

// ByMonth get events that fall in the last month
func ByMonth(d diary.Diary) *diary.Diary {
	now := time.Now()
	return ByRange(d, now.Add(-Month), now)
}

// ByTag returns the entries that have that a given tag
func ByTag(d diary.Diary, tag string) *diary.Diary {
	var emptyD diary.Diary
	for _, r := range d {
		for _, t := range r.Tags {
			if t == tag {
				(&emptyD).AddEntry(r)
			}
		}
	}
	return &emptyD
}
