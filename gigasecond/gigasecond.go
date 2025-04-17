// Package gigasecond tells you what time is it after 1.000.000.000 seconds from a given date
package gigasecond

// import path for the time package from the standard library
import "time"

// AddGigasecond returns what time is it after 1.000.000.000 seconds from a given date
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Second * 1000000000)
	return t
}
