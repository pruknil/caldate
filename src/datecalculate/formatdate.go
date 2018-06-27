package datecalculate

import (
	"fmt"
	"strconv"
	"time"
)

func formatDate(inputDate string) string {
	day, month, year := convertDate(inputDate)
	date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

	return fmt.Sprintf("%v, %v %v %v", date.Weekday(), date.Day(), date.Month(), date.Year())
}

func convertDate(inputDate string) (int, int, int) {
	day, _ := strconv.Atoi(inputDate[0:2])
	month, _ := strconv.Atoi(inputDate[2:4])
	year, _ := strconv.Atoi(inputDate[4:8])

	return day, month, year
}
