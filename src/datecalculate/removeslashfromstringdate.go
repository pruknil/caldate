package datecalculate

import "fmt"

func RemoveSlashFromStringDate(date string) string {
	var day, month, year int
	fmt.Sscanf(date, "%d/%d/%d", &day, &month, &year)
	return fmt.Sprintf("%02d%02d%d", day, month, year)
}
