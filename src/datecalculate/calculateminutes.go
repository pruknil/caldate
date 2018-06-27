package datecalculate

import (
	"fmt"
	"strconv"
)

func calculateMinutes(days int) string {
	minutes := strconv.Itoa(days * HOURS_PER_DAY * MINUTES_PER_HOUR)

	return fmt.Sprintf(formatNumber(minutes) + " minutes")
}
