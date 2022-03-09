package policy

import (
	"github.com/rclone/rclone/backend/union/upstream"
)

func init() {
	registerPolicy("seq", &Seq{})
}

// Seq stands for sequential
// Action category: same as all
// Create category: try each remote in the order they were passed in the configuration file, until one succeeds
// Search category: same as all.
type Seq struct {
	All
}

// TODO HOW TO HANDLE CASE WHERE USER WAENTS TO UPDATE A FILE WITH A MUCH LARGER FILE, WHICH NOW CAUSES THE BACKEND TO FAIL?
func (p *Seq) CreateAttempts(upstreams []*upstream.Fs) [][]*upstream.Fs {
	var attempts [][]*upstream.Fs
	for _, u := range upstreams {
		attempts = append(attempts, []*upstream.Fs{u})
	}
	return attempts
}

func (p *Seq) CreateEntriesAttempts(entries []upstream.Entry) [][]upstream.Entry {
	var attempts [][]upstream.Entry
	for _, e := range entries {
		attempts = append(attempts, []upstream.Entry{e})
	}
	return attempts
}
