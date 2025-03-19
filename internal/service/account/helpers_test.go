package service

import (
	"testing"
	"time"
)

func TestGetLastTaxYearStartDate(t *testing.T) {
	now := time.Now()

	tests := []struct {
		name     string
		taxMonth int
		taxDay   int
	}{
		{"Tax year starting April 6", 4, 6},
		{"Tax year starting January 1", 1, 1},
		{"Tax year starting December 31", 12, 31},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			year := now.Year()
			// Create the tax year start date for the current year.
			taxYearStartDate := time.Date(year, time.Month(tc.taxMonth), tc.taxDay, 0, 0, 0, 0, time.UTC)
			var expected time.Time
			// If today is before the tax start date, expect the previous year's date.
			if now.Before(taxYearStartDate) {
				expected = time.Date(year-1, time.Month(tc.taxMonth), tc.taxDay, 0, 0, 0, 0, time.UTC)
			} else {
				expected = taxYearStartDate
			}

			result := GetLastTaxYearStartDate(tc.taxMonth, tc.taxDay)
			if !result.Equal(expected) {
				t.Errorf("For tax start %d-%d: expected %v, got %v", tc.taxMonth, tc.taxDay, expected, result)
			}
		})
	}
}

func TestParseMonthDay(t *testing.T) {
	tests := []struct {
		name       string
		dateStr    string
		wantMonth  int
		wantDay    int
		shouldFail bool
	}{
		{"Valid date", "12-25", 12, 25, false},
		{"Valid date with leading zeros", "01-01", 1, 1, false},
		{"Invalid month", "13-01", 0, 0, true},
		{"Invalid format", "2022-01-01", 0, 0, true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			month, day, err := parseMonthDay(tc.dateStr)
			if tc.shouldFail {
				if err == nil {
					t.Errorf("Expected error for input %q, got none", tc.dateStr)
				}
				return
			}
			if err != nil {
				t.Errorf("Unexpected error for input %q: %v", tc.dateStr, err)
				return
			}
			if month != tc.wantMonth {
				t.Errorf("For input %q, got month %d; want %d", tc.dateStr, month, tc.wantMonth)
			}
			if day != tc.wantDay {
				t.Errorf("For input %q, got day %d; want %d", tc.dateStr, day, tc.wantDay)
			}
		})
	}
}
