package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDurationInput412018And472018ShouldbeJSON(t *testing.T) {
	var requestBody = Request{
		StartDate: "4/1/2018",
		EndDate:   "4/7/2018"}

	url := "/duration"
	requestBodyString, _ := json.Marshal(requestBody)
	request := httptest.NewRequest("POST", url, bytes.NewBuffer(requestBodyString))
	responseRecord := httptest.NewRecorder()
	Calculate(responseRecord, request)
	result := responseRecord.Result()
	actual, _ := ioutil.ReadAll(result.Body)
	expectedResponse := Response{
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
	expectedResponseString, _ := json.Marshal(expectedResponse)

	if strings.TrimSpace(string(actual)) != strings.TrimSpace(string(expectedResponseString)) {
		t.Errorf("expected %s\n but got\n %s", actual, expectedResponseString)
	}

}
