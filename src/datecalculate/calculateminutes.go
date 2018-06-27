package datecalculate

import (
	"fmt"
	"strconv"
)

const HOUR_DAY = 24
const MINUTE_HOUR = 60

func calculateMinutes(days int) string {
	minutes := strconv.Itoa(days * HOUR_DAY * MINUTE_HOUR)

	return fmt.Sprintf(formatNumber(minutes) + " minutes")
}
