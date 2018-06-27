package datecalculate

import "testing"

func TestCalculateWeek182Shouldbe26(t *testing.T) {
	days := 182
	expected := "26 weeks"

	actual := calculateWeek(days)

	if expected != actual {
		t.Errorf("Expected= %s Actual= %s",expected,actual)
	}

}


func TestCalculateWeek15083Shouldbe2154WeeksAnd5Days(t *testing.T) {
	days := 15083
	expected := "2154 weeks and 5 days"

	actual := calculateWeek(days)

	if expected != actual {
		t.Errorf("Expected= %s Actual= %s",expected,actual)
	}

}
