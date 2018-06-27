package datecalculate

import "testing"

func TestFormatDateInput04012018ShouldBeThursday4January2018(t *testing.T) {
	expected := "Thursday, 4 January 2018"
	startDate := "04012018"

	actualDate := formatDate(startDate)

	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}

func TestFormatDateInput04072018ShouldBeWednesday4July2018(t *testing.T) {
	expected := "Wednesday, 4 July 2018"
	startDate := "04072018"

	actualDate := formatDate(startDate)

	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
