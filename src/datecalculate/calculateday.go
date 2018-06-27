package datecalculate

import (
	"time"
)

const INCLUDEING_START_DAY = 1

func calculateDay(startDate string, endDate string) int {

	startDay, startMonth, startYear := convertDate(startDate)
	endDay, endMonth, endYear := convertDate(endDate)

	startDateAsTime := time.Date(startYear, time.Month(startMonth), startDay, 0, 0, 0, 0, time.UTC)
	endDateAsTime := time.Date(endYear, time.Month(endMonth), endDay, 0, 0, 0, 0, time.UTC)

	diff := endDateAsTime.Sub(startDateAsTime)

	return (int(diff.Hours()) / HOURS_PER_DAY) + INCLUDEING_START_DAY
}

