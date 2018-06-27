package datecalculate

import "testing"

func TestDayNumShouldbe182(t *testing.T){
	startdate := "04012018"
	enddate := "04072018"
	expected := 182

	actual := calculateDay(startdate,enddate)

	if expected != actual {
		t.Errorf("Expected %d Actual %d",expected,actual)
	}

}