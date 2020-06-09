package roro

import "time"

// Day returns the day of the month specified by t.
func (r Roro) Day() int {
	return r.o.Day()
}

// Date returns the year, month, and day in which t occurs.
func (r Roro) Date() (int, time.Month, int) {
	return r.o.Date()
}

// YearDay returns the day of the year specified by t, in the range [1,365] for non-leap years,
// and [1,366] in leap years.
func (r Roro) YearDay() int {
	return r.o.YearDay()
}

// StartOfDay returns the start time of the day in which r occurs.
func (r Roro) StartOfDay() (start Roro) {
	year, month, day := r.Date()

	start = Date(year, month, day, 0, 0, 0, 0, r.o.Location())

	return
}

// EndOfDay returns the end time of the day in which r occurs.
func (r Roro) EndOfDay() (end Roro) {
	end = r.StartOfDay().AddDate(0, 0, 1).Add(time.Duration(-1))

	return
}

// DiffDays returns the positive days difference between r and d.
func (r Roro) DiffDays(d Roro) int64 {
	if r.After(d) {
		return (r.Unix() - d.Unix()) / SecondsInDay
	} else {
		return (d.Unix() - r.Unix()) / SecondsInDay
	}
}

// IsSameDay checks if the given date is as same day as the instance.
func (r Roro) IsSameDay(d Roro) bool {
	return r.Year() == d.Year() && r.YearDay() == d.YearDay()
}

// IsTomorrow checks if the given instance is as same day as current moment tomorrow.
func (r Roro) IsTomorrow() bool {
	var diffSecond = time.Duration(r.Unix() - Now().Unix())

	return diffSecond > 0 && diffSecond > time.Hour * 24 && diffSecond < time.Hour * 24 * 2
}

// IsYesterday checks if the given instance is as same day as current moment yesterday.
func (r Roro) IsYesterday() bool {
	var diffSecond = time.Duration(Now().Unix() - r.Unix())

	return diffSecond > 0 && diffSecond > time.Hour * 24 && diffSecond < time.Hour * 24 * 2
}
