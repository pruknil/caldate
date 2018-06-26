package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
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
		From:          "4/1/2018",
		To:            "4/7/2018",
		Totalday:      "182",
		Humanreadday:  "6 months, 1 day",
		Second:        "15,724,800",
		Minute:        "262,080",
		Hour:          "4368",
		Week:          "26",
		Percentofyear: "49.86",
	}
	expectedResponseString, _ := json.Marshal(expectedResponse)

	if string(actual) != string(expectedResponseString) {
		t.Errorf("expected %s but got %s", expectedResponseString, actual)
	}

}
