package datecalculate

import "fmt"

const DAYS_PER_WEEK  = 7

func calculateWeek(days int) string {

	if days%DAYS_PER_WEEK == 0 {
		return fmt.Sprintf("%d weeks",days/DAYS_PER_WEEK)
	}

	return fmt.Sprintf("%d weeks and %d days",days/DAYS_PER_WEEK,days%DAYS_PER_WEEK)
}
