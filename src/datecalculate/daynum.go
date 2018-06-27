package datecalculate

import (
	"time"
)

const hourPerDay  = 24
const includeStartDay  = 1

func dayNum(startDate string,endDate string) int {

	startDay,startMonth,startYear := convertDate(startDate)
	endDay,endMonth,endYear := convertDate(endDate)

	startDateAsTime := time.Date(startYear,time.Month(startMonth),startDay,0,0,0,0,time.UTC)
	endDateAsTime := time.Date(endYear,time.Month(endMonth),endDay,0,0,0,0,time.UTC)

	diff := endDateAsTime.Sub(startDateAsTime)

	return (int(diff.Hours())/hourPerDay) + includeStartDay
}