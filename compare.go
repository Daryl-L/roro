package roro

// Equal reports whether t and u represent the same time instant.
// Two times can be equal even if they are in different locations.
// For example, 6:00 +0200 and 4:00 UTC are Equal.
// See the documentation on the Time type for the pitfalls of using == with
// Time values; most code should use Equal instead.
func (r Roro) Equal(u Roro) bool {
	return r.o.Equal(u.o)
}

// Before reports whether the time instant t is before u.
func (r Roro) Before(u Roro) bool {
	return r.o.Before(u.o)
}

// After reports whether the time instant t is after u.
func (r Roro) After(u Roro) bool {
	return r.o.After(u.o)
}

// IsZero reports whether t represents the zero time instant,
// January 1, year 1, 00:00:00 UTC.
func (r Roro) IsZero(u Roro) bool {
	return r.o.IsZero()
}