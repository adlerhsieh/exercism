package leap

const testVersion = 2

// Check if a year is a leap year
func IsLeapYear(year int) bool {
	switch {
	case year%400 == 0:
		return true
	case year%100 == 0:
		return false
	case year%4 == 0:
		return true
	}
	return false
}
