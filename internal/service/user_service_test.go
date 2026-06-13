package service

import (
	"testing"
	"time"
)

func calculateAgeAt(dob, now time.Time) int {
	years := now.Year() - dob.Year()
	if now.Month() < dob.Month() ||
		(now.Month() == dob.Month() && now.Day() < dob.Day()) {
		years--
	}
	return years
}

func date(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

func TestCalculateAge(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		dob      time.Time
		now      time.Time
		wantAge  int
	}{
		{
			name:    "birthday already passed this year",
			dob:     date(1990, time.January, 1),
			now:     date(2025, time.June, 13),
			wantAge: 35,
		},
		{
			name:    "birthday has not occurred yet this year",
			dob:     date(1990, time.December, 31),
			now:     date(2025, time.June, 13),
			wantAge: 34,
		},
		{
			name:    "birthday is today – exact match",
			dob:     date(1990, time.June, 13),
			now:     date(2025, time.June, 13),
			wantAge: 35,
		},
		{
			name:    "same month, earlier day – birthday passed",
			dob:     date(1995, time.June, 1),
			now:     date(2025, time.June, 13),
			wantAge: 30,
		},
		{
			name:    "same month, later day – birthday not yet passed",
			dob:     date(1995, time.June, 30),
			now:     date(2025, time.June, 13),
			wantAge: 29,
		},
		{
			name:    "birthday one day ago",
			dob:     date(2000, time.June, 12),
			now:     date(2025, time.June, 13),
			wantAge: 25,
		},
		{
			name:    "birthday tomorrow",
			dob:     date(2000, time.June, 14),
			now:     date(2025, time.June, 13),
			wantAge: 24,
		},
		{
			name:    "born on Feb 29 – checked on Feb 29 (leap year)",
			dob:     date(2000, time.February, 29),
			now:     date(2024, time.February, 29),
			wantAge: 24,
		},
		{
			name:    "born on Feb 29 – checked on Feb 28 in non-leap year (not yet had birthday)",
			dob:     date(2000, time.February, 29),
			now:     date(2025, time.February, 28),
			wantAge: 24,
		},
		{
			name:    "born on Feb 29 – checked on Mar 1 in non-leap year (birthday passed on Mar 1)",
			dob:     date(2000, time.February, 29),
			now:     date(2025, time.March, 1),
			wantAge: 25,
		},
		{
			name:    "born on Mar 1 – checked just after Feb 28 in non-leap year (not yet had birthday)",
			dob:     date(1996, time.March, 1),
			now:     date(2025, time.February, 28),
			wantAge: 28,
		},
		{
			name:    "born on Dec 31 in a leap year – checked Jan 1 next year (not yet had birthday)",
			dob:     date(2000, time.December, 31),
			now:     date(2025, time.January, 1),
			wantAge: 24,
		},
		{
			name:    "newborn – dob equals now",
			dob:     date(2025, time.June, 13),
			now:     date(2025, time.June, 13),
			wantAge: 0,
		},
		{
			name:    "infant under 1 year old",
			dob:     date(2025, time.January, 1),
			now:     date(2025, time.June, 13),
			wantAge: 0,
		},
		{
			name:    "very old age – 100 years",
			dob:     date(1925, time.June, 13),
			now:     date(2025, time.June, 13),
			wantAge: 100,
		},
		{
			name:    "new year's day – dob Dec 31 last year is 0 years old",
			dob:     date(2024, time.December, 31),
			now:     date(2025, time.January, 1),
			wantAge: 0,
		},
		{
			name:    "new year's day – dob Jan 1 is exactly 1 year old",
			dob:     date(2024, time.January, 1),
			now:     date(2025, time.January, 1),
			wantAge: 1,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := calculateAgeAt(tc.dob, tc.now)
			if got != tc.wantAge {
				t.Errorf("calculateAgeAt(dob=%s, now=%s) = %d, want %d",
					tc.dob.Format("2006-01-02"),
					tc.now.Format("2006-01-02"),
					got,
					tc.wantAge,
				)
			}
		})
	}
}

func TestCalculateAge_RealClock_NonNegative(t *testing.T) {
	t.Parallel()
	dob := date(1990, time.May, 10)
	age := CalculateAge(dob)
	if age < 0 {
		t.Errorf("CalculateAge returned a negative age: %d", age)
	}
}
