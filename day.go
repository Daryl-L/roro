package roro

import "time"

// Day returns the day of the month specified by t.
func (r Roro) Day() int {
	return r.o.Day()
}

// StartOfDay returns the start time of the day in which r occurs.
func (r Roro) StartOfDay() (start Roro) {
	year, month, day := r.o.Date()

	start = Date(year, month, day, 0, 0, 0, 0, r.o.Location())

	return
}

// EndOfDay returns the end time of the day in which r occurs.
func (r Roro) EndOfDay() (end Roro) {
	end = r.StartOfDay().AddDate(0, 0, 1).Add(time.Duration(-1))

	return
}