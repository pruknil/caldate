package datecalculate

import "strconv"

func CallHr(daynumber int) string {

	numberofHr := strconv.Itoa(daynumber * 24)
	return (numberofHr)
}
