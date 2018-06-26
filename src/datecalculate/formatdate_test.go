package datecalculate

import "testing"

func TestFormatDateInput412018ShouldBeThurday4January2018(t *testing.T) {
	expected := "Thursday, 4 January 2018"
	startDate := "04012018"
	actualDate := FormatDate(startDate)
	if actualDate != expected {
		t.Errorf("expected %s but got %s", expected, actualDate)
	}
}
