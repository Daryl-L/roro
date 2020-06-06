package roro

import "time"

// Weekday returns the day of the week specified by r.
func (r Roro) WeekDay() time.Weekday {
	return r.o.Weekday()
}

// Weekday returns the day of the ISO 8601 week specified by r.
func (r Roro) ISOWeekDay() (day int) {
	day = int(r.WeekDay())

	if day == 0 {
		day = 7
	} else {
		day--
	}

	return
}

// StartOfDay returns the start day of the week in which r occurs.
func (r Roro) StartOfWeek() (startDay Roro) {
	startDay = r.StartOfDay().AddDate(0, 0, -int(r.WeekDay()))

	return
}

// StartOfDay returns the start day of the ISO 8601 week in which r occurs.
func (r Roro) ISOStartOfWeek() (startDay Roro) {
	startDay = r.StartOfDay().AddDate(0, 0, -r.ISOWeekDay())

	return
}

// StartOfDay returns the end day of the week in which r occurs.
func (r Roro) EndOfWeek() (startDay Roro) {
	startDay = r.StartOfDay().AddDate(0, 0, int(r.WeekDay()))

	return
}

// StartOfDay returns the end day of the ISO 8601 week in which r occurs.
func (r Roro) ISOEndOfWeek() (startDay Roro) {
	startDay = r.StartOfDay().AddDate(0, 0, r.ISOWeekDay())

	return
}

// DaysInWeek returns the dates of the week in which r occurs.
func (r Roro) DaysInWeek() (days []Roro) {
	var start = Today()
	days = make([]Roro, 7)

	for i := 0; i < 7; i++ {
		days[i] = start.AddDate(0, 0, i+1)
	}

	return
}

// ISOWeek returns the ISO 8601 year and week number in which t occurs.
// Week ranges from 1 to 53. Jan 01 to Jan 03 of year n might belong to
// week 52 or 53 of year n-1, and Dec 29 to Dec 31 might belong to week 1
// of year n+1.
func (r Roro) ISOWeek() (int, int) {
	return r.o.ISOWeek()
}

// IsSameWeek checks if the given date d is the same week as date r.
func (r Roro) IsSameWeek(d Roro) bool {
	var (
		rWeek, dWeek int
	)

	rWeek, _ = r.ISOWeek()
	dWeek, _ = d.ISOWeek()

	return rWeek == dWeek
}

// IsSameWeek checks if the date r is the same week as the current moment.
func (r Roro) IsCurrentWeek() bool {
	var (
		cWeek, rWeek int
	)

	cWeek, _ = Now().ISOWeek()
	rWeek, _ = r.ISOWeek()

	return cWeek == rWeek
}

// IsSameWeek checks if the date r is the same week as the current moment next week.
func (r Roro) IsNextWeek() bool {
	var (
		cWeek, rWeek int
	)

	cWeek, _ = Now().ISOWeek()
	rWeek, _ = r.ISOWeek()

	return rWeek > cWeek
}

// IsSameWeek checks if the date r is the same week as the current moment last week.
func (r Roro) IsLastWeek() bool {
	var (
		cWeek, rWeek int
	)

	cWeek, _ = Now().ISOWeek()
	rWeek, _ = r.ISOWeek()

	return rWeek < cWeek
}
