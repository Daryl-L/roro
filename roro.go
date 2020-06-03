package roro

import "time"

type Roro struct {
	o time.Time
}

// Now returns the current local time.
func Now() (now Roro) {
	now = Roro{time.Now()}

	return
}

