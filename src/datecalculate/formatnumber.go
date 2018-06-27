package datecalculate

const MONEY_SEPERATOR = 3

func formatNumber(number string) string {

	if len(number) <= MONEY_SEPERATOR {
		return number
	} else {

		return formatNumber( number[0:len(number)-MONEY_SEPERATOR] ) + "," + number[len(number)-MONEY_SEPERATOR:]
	}

}
