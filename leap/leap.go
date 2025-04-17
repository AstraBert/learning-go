// The leap package tells you whether or not a year is leap
package leap

// IsLeapYear returns whether or not a year is leap
func IsLeapYear(year int) bool {
	if year % 4 != 0 {
		return false
	} else if year % 100 == 0 && year % 400 != 0 {
		return false
	} else {
		return true
	}
}
