package datecalculate

import "testing"

func TestCalcualteMinutesInput472018ShouldBe262080(t *testing.T) {
	expected := "262,080 minutes"
	days := 182

	actual := calculateMinutes(days)

	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}
}
