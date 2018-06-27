package datecalculate

import "testing"

func TestRemoveSlashFromStringDateInput4Slash1Slash2018ShouldBe04012018(t *testing.T) {
	date := "4/1/2018"
	expected := "04012018"
	actual := RemoveSlashFromStringDate(date)
	if actual != expected {
		t.Errorf("Expected: %s but got %s", expected, actual)
	}
}

func TestRemoveSlashFromStringDateInput4Slash7Slash2018ShouldBe04072018(t *testing.T) {
	date := "4/7/2018"
	expected := "04072018"
	actual := RemoveSlashFromStringDate(date)
	if actual != expected {
		t.Errorf("Expected: %s but got %s", expected, actual)
	}
}
