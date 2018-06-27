package datecalculate

import (
	"testing"
)

func TestCallHrInput182Should4368(t *testing.T) {
	expected := "4368 hours"
	numberOfDay := 182
	actual := calculateHours(numberOfDay)
	if actual != expected {
		t.Errorf("expected %s but got %s", expected, actual)
	}

}
