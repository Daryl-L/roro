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

// Date returns the Time corresponding to
//	yyyy-mm-dd hh:mm:ss + nsec nanoseconds
// in the appropriate zone for that time in the given location.
//
// The month, day, hour, min, sec, and nsec values may be outside
// their usual ranges and will be normalized during the conversion.
// For example, October 32 converts to November 1.
//
// A daylight savings time transition skips or repeats times.
// For example, in the United States, March 13, 2011 2:15am never occurred,
// while November 6, 2011 1:15am occurred twice. In such cases, the
// choice of time zone, and therefore the time, is not well-defined.
// Date returns a time that is correct in one of the two zones involved
// in the transition, but it does not guarantee which.
//
// Date panics if loc is nil.
func Date(year int, month time.Month, day int, hour, min, sec, nsec int, loc *time.Location) (date Roro) {
	date.o = time.Date(year, month, day, hour, min, sec, nsec, loc)

	return
}

