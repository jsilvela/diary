package diary

import "testing"

func TestAddEntry(t *testing.T) {
	var d Diary
	(&d).AddEntry(&Record{Tags: []string{"hello", "there"}})
	if len(d) != 1 {
		t.Errorf("After adding one entry to empty diary, expected it to have 1 entry. Got %d", len(d))
	}
	if d[0].Tags[0] != "hello" {
		t.Errorf("After storing single entry with tag 'hello', expected it to be the first tag of the first entry. Was %s", d[0].Tags[0])
	}
}
