package datecalculate

import "strconv"

func CalculateHours(daynumber int) string {

	numberofHr := strconv.Itoa(daynumber * 24)
	return (numberofHr)
}
