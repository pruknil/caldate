package datecalculate

import (
	"testing"
)

func TestCalculateHoursInput182ShouldBe4368Hours(t *testing.T) {
	expected := "4368 hours"
	numberOfDay := 182
	actual := calculateHours(numberOfDay)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}

}
