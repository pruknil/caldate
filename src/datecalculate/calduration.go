package datecalculate

import "fmt"

type Request struct {
	StartDate string `json:"startdate,omitempty"`
	EndDate   string `json:"enddate,omitempty"`
}

type Response struct {
	From          string `json:"from,omitempty"`
	To            string `json:"to,omitempty"`
	Totalday      string `json:"totalday,omitempty"`
	Humanreadday  string `json:"humanreadday,omitempty"`
	Second        string `json:"second,omitempty"`
	Minute        string `json:"minute,omitempty"`
	Hour          string `json:"hour,omitempty"`
	Week          string `json:"week,omitempty"`
	Percentofyear string `json:"percentofyear,omitempty"`
}

func CalDuration(startDate string, endDate string) Response {

	startDateWithoutSlash := removeSlashFromStringDate(startDate)
	endDateWithoutSlash := removeSlashFromStringDate(endDate)
	calculateDayInt := calculateDay(startDateWithoutSlash, endDateWithoutSlash)

	return Response{
		From:         FormatDate(startDateWithoutSlash),
		To:           FormatDate(endDateWithoutSlash),
		Totalday:     fmt.Sprintf("%d", calculateDayInt),
		Humanreadday: calculateTotalYMD(startDateWithoutSlash, endDateWithoutSlash),
		Second:       calculateSecond(calculateDayInt),
		Minute:       calculateMinutes(calculateDayInt),
		Hour:         calculateHours(calculateDayInt),
		Week:         calculateWeek(calculateDayInt),
	}
}
