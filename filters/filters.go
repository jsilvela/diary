package filters

import (
	"github.com/jsilvela/diary"
	"time"
)

const (
	Week = time.Hour * 24 * 7
)

func By_range(d diary.Diary, from time.Time, to time.Time) *diary.Diary {
	var emptyD diary.Diary
	for _, r := range d {
		if r.Event_time.After(from) && r.Event_time.Before(to) {
			(&emptyD).Add_entry(r)
		}
	}
	return &emptyD
}

func By_week(d diary.Diary) *diary.Diary {
	now := time.Now()
	return By_range(d, now.Add(-Week), now)
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
