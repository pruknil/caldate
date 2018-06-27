package datecalculate

import (
	"fmt"
	"time"
)

func calculateTotalYMD(startDate string, endDate string) string {
	var stringDate string

	startDays, startMonths, startYears := convertDate(startDate)
	endDays, endMonths, endYears := convertDate(endDate)
	years := endYears - startYears
	months := endMonths - startMonths
	days := endDays - startDays + 1
	if days < 0 {
		newDate := time.Date(startYears, time.Month(startMonths), 32, 0, 0, 0, 0, time.UTC)
		days += 32 - newDate.Day()
		months--
	}
	if months < 0 {
		months += 12
		years--
	}

	if years > 0 {
		stringDate = fmt.Sprintf("%v years, %v months, %v days", years, months, days)
	} else if months > 0 {
		stringDate = fmt.Sprintf("%v months, %v day", months, days)
	}
	return stringDate
}
