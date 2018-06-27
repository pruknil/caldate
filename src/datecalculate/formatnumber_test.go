package datecalculate

import "testing"

func TestFormatNumberFrom1000Shouldbe1comma000(t *testing.T) {
	number := "1000"
	expected := "1,000"

	actual := formatNumber(number)

	if actual != expected {
		t.Errorf("Expected %s Actual %s",expected,actual)
	}

}

func TestFormatNumberFrom1000000Shouldbe1comma000comma000(t *testing.T) {
	number := "1000000"
	expected := "1,000,000"

	actual := formatNumber(number)

	if actual != expected {
		t.Errorf("Expected %s Actual %s",expected,actual)
	}

}