package roro

const (
	SecondsInMinute = 60
	SecondsInHour   = SecondsInMinute * MinutesInHour
	SecondsInDay    = SecondsInHour * HoursInDay
)

// Unix returns t as a Unix time, the number of seconds elapsed
// since January 1, 1970 UTC. The result does not depend on the
// location associated with t.
// Unix-like operating systems often record time as a 32-bit
// count of seconds, but since the method here returns a 64-bit
// value it is valid for billions of years into the past or future.
func (r Roro) Unix() int64 {
	return r.o.Unix()
}
