package roro

// Year returns the year in which t occurs.
func (r Roro) Year() int {
	return r.o.Year()
}

// IsSameYear checks if the given date d is the same year as the date r.
func (r Roro) IsSameYear(d Roro) bool {
	return r.Year() == d.Year()
}

// IsCurrentYear checks if the date r is the same year as current moment.
func (r Roro) IsCurrentYear() bool {
	return Now().Year() == r.Year()
}

// IsLastYear checks if the date r is the same year as the current moment last year.
func (r Roro) IsLastYear() bool {
	return Now().Year()-r.Year() == 1
}

// IsLastYear checks if the date r is the same year as the current moment next year.
func (r Roro) IsNextYear() bool {
	return Now().Year()-r.Year() == -1
}
