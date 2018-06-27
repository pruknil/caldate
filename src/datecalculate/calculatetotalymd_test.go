package datecalculate

import "testing"

func TestCalcualteTotalYMDInput01042018And04072018ShouldBe6Months1Day(t *testing.T) {
	expected := "6 months, 1 day"
	startDate := "04012018"
	endDate := "04072018"

	actual := calculateTotalYMD(startDate, endDate)

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}

func TestCalcualteTotalYMDInput11041977And21072018ShouldBe41Years3Months17Days(t *testing.T) {
	expected := "41 years, 3 months, 17 days"
	startDate := "11041977"
	endDate := "27072018"

	actual := calculateTotalYMD(startDate, endDate)

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
