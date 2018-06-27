package datecalculate

import (
	"fmt"
)

const HOUR_DAY = 24
const MINUTE_HOUR = 60

func calculateMinutes(days int) string {
	minutes := days * HOUR_DAY * MINUTE_HOUR
	minutesString := fmt.Sprintf("%v", minutes)

	return fmt.Sprintf(minutesString + " seconds")
}
