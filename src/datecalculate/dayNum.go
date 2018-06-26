package datecalculate

import (
	"time"
)

func dayNum(startdate string,enddate string) int {

	startDate,startMonth,startYear := convertDate(startdate)
	endDate,endMonth,endYear := convertDate(enddate)

	startDateAsTime := time.Date(startYear,time.Month(startMonth),startDate,0,0,0,0,time.UTC)
	endDateAsTime := time.Date(endYear,time.Month(endMonth),endDate,0,0,0,0,time.UTC)

	diff := endDateAsTime.Sub(startDateAsTime)

	return (int(diff.Hours())/24) + 1

}