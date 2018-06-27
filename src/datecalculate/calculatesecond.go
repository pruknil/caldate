package datecalculate

import (
	"strconv"
	"fmt"
)

const MINUTES_PER_HOUR = 60
const SECONDS_PER_MINUTE = 60

func calculateSecond(days int) string {
	seconds := strconv.Itoa(days*HOURS_PER_DAY*MINUTES_PER_HOUR*SECONDS_PER_MINUTE)
	return fmt.Sprintf("%s seconds",formatNumber(seconds))
}
