package datecalculate

import "testing"

func TestDayNumShouldbe182(t *testing.T) {
	startdate := "04012018"
	enddate := "04072018"
	expected := 182

	actual := calculateDay(startdate, enddate)

	if expected != actual {
		t.Errorf("Expected %d Actual %d", expected, actual)
	}

}

func TestDayNumShouldbe15083(t *testing.T) {
	startdate := "11041977"
	enddate := "27072018"
	expected := 15083

	actual := calculateDay(startdate, enddate)

	if expected != actual {
		t.Errorf("Expected %d Actual %d", expected, actual)
	}

}
