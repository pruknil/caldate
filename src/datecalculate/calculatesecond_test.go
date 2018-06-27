package datecalculate

import "testing"

func TestCalculateSecond182Shouldbe15724800(t *testing.T) {
	days := 182
	expected := "15,724,800 seconds"
	actual := calculateSecond(days)
	if expected != actual {
		t.Errorf("Expected: %s Actual: %s", expected, actual)
	}

}
