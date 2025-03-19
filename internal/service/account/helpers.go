package service

import (
	"time"
)

// GetLastTaxYearStartDate returns the last occurrence of the tax year start date
// (defaulting to 6th April if not provided).
func GetLastTaxYearStartDate(taxYearStartMonth int, taxYearStartDay int) time.Time {
	now := time.Now()
	year := now.Year()

	// Create the tax year start date for the current year
	taxYearStartDate := time.Date(year, time.Month(taxYearStartMonth), taxYearStartDay, 0, 0, 0, 0, time.UTC)

	// If today is before the tax start date, move to the previous year's date
	if now.Before(taxYearStartDate) {
		taxYearStartDate = time.Date(year-1, time.Month(taxYearStartMonth), taxYearStartDay, 0, 0, 0, 0, time.UTC)
	}

	return taxYearStartDate
}

// parseMonthDay - accepts a date string in the format MM-DD and returns the month and day as integers, or an error
func parseMonthDay(dateStr string) (int, int, error) {
	t, err := time.Parse("01-02", dateStr)
	if err != nil {
		return 0, 0, err
	}
	return int(t.Month()), t.Day(), nil
}
