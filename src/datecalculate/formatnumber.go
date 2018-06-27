package datecalculate


func formatNumber(number string) string {

	if len(number) <= 3 {
		return number
	} else {

		return formatNumber( number[0:len(number)-3] ) + "," + number[len(number)-3:]
	}

}
