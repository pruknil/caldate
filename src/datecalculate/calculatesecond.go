package datecalculate

import (
	"strconv"
	"fmt"
)

const HOUR_PER_DAY = 24
const MINUTE_PER_HOUR = 60
const SECOND_PER_MINUTE = 60

func calculateSecond(days int) string {
	seconds := strconv.Itoa(days*HOUR_PER_DAY*MINUTE_PER_HOUR*SECOND_PER_MINUTE)
	return fmt.Sprintf("%s seconds",formatNumber(seconds))
}
