package filters

import (
	"github.com/jsilvela/diary"
	"time"
)

const (
	Week  = time.Hour * 24 * 7
	Month = time.Hour * 24 * 30
)

func By_range(d diary.Diary, from time.Time, to time.Time) *diary.Diary {
	var dr diary.Diary
	for _, r := range d {
		if r.Event_time.After(from) && r.Event_time.Before(to) {
			(&dr).Add_entry(r)
		}
	}
	return &dr
}

func By_week(d diary.Diary) *diary.Diary {
	now := time.Now()
	return By_range(d, now.Add(-Week), now)
}

func By_month(d diary.Diary) *diary.Diary {
	now := time.Now()
	return By_range(d, now.Add(-Month), now)
}

func By_tag(d diary.Diary, tag string) *diary.Diary {
	var emptyD diary.Diary
	for _, r := range d {
		for _, t := range r.Tags {
			if t == tag {
				(&emptyD).Add_entry(r)
			}
		}
	}
	return &emptyD
}
