package datecalculate

import (
	"fmt"
)

const HOURS_PER_DAY  = 24

func calculateHours(daynumber int) string {

	return fmt.Sprintf("%d hours",daynumber * HOURS_PER_DAY)
}
