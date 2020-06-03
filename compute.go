package roro

import "time"

// Add returns the time r+d.
func (r Roro) Add(d time.Duration) (res Roro) {
	res.o = r.o.Add(d)

	return
}

// AddDate returns the time corresponding to adding the
// given number of years, months, and days to r.
// For example, AddDate(-1, 2, 3) applied to January 1, 2011
// returns March 4, 2010.
//
// AddDate normalizes its result in the same way that Date does,
// so, for example, adding one month to October 31 yields
// December 1, the normalized form for November 31.
func (r Roro) AddDate(years, months, days int) (res Roro) {
	res.o = r.o.AddDate(years, months, days)

	return
}