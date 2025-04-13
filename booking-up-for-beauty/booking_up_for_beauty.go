package booking

import (
	"time"
	"fmt"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	parsedDate, _ := time.Parse(layout, date)
	dateNow := time.Now()
	return dateNow.After(parsedDate)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout := "Monday, January 2, 2006 15:04:05"
	t, _ := time.Parse(layout, date)
	if t.Hour() < 18 && t.Hour() >= 12 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	t, _ := time.Parse(layout, date)
	description := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", t.Weekday().String(), t.Month().String(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return description
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentDate := time.Now()
	t := time.Date(currentDate.Year(),time.September,15,0,0,0,0,time.UTC)
	return t
}
