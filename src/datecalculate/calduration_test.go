package datecalculate

import (
	"testing"
)

func TestCalDurationInput412018And472018ShouldBeStruct(t *testing.T) {
	startDate := "4/1/2018"
	endDate := "4/7/2018"
	expected := Response{
		From:          "Thursday, 4 January 2018",
		To:            "Wednesday, 4 July 2018",
		Totalday:      "182",
		Humanreadday:  "6 months, 1 day",
		Second:        "15,724,800 seconds",
		Minute:        "262,080 minutes",
		Hour:          "4368 hours",
		Week:          "26 weeks",
		Percentofyear: "49.86% of 2018",
	}
	actual := CalDuration(startDate, endDate)
	if actual != expected {
		t.Errorf("Expected:\n %s but got\n %s", expected, actual)
	}

}
