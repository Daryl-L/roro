package roro

import "time"

// Weekday returns the day of the week specified by r.
func (r Roro) WeekDay() (time.Weekday) {
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