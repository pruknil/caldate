package datecalculate

import "testing"

func TestCalculateSecond182Shouldbe15724800(t *testing.T) {
	days := 182
	expexted := "15,724,800 seconds"

	actual := calculateSecond(days)
	if expexted != actual {
		t.Errorf("Expected: %s Actual: %s",expexted,actual)
	}

}