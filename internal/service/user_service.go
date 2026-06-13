package service

import "time"

// CalculateAge takes a Date of Birth and returns the current age in years.
// Age is decremented by 1 if the birthday has not yet occurred in the current year.
func CalculateAge(dob time.Time) int {
	now := time.Now()
	years := now.Year() - dob.Year()

	// Use month+day comparison — YearDay() is incorrect across leap years
	// because Feb 29 shifts all subsequent day-numbers by 1 in non-leap years.
	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}
	return years
}
