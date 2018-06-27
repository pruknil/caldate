package api

import (
	"encoding/json"
	"net/http"
)

func Calculate(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var req Request
	err := decoder.Decode(&req)

	if err != nil {
		panic(err)
	}

	var resp Response
	resp.From = "4/1/2018"
	resp.To = "4/7/2018"
	resp.Totalday = "182"
	resp.Humanreadday = "6 months, 1 day"
	resp.Second = "15,724,800"
	resp.Minute = "262,080"
	resp.Hour = "4368"
	resp.Week = "26"
	resp.Percentofyear = "49.86"

	json.NewEncoder(w).Encode(resp)
	return
}

type Request struct {
	StartDate string `json:"startdate,omitempty"`
	EndDate   string `json:"enddate,omitempty"`
}

type Response struct {
	From          string `json:"from,omitempty"`
	To            string `json:"to,omitempty"`
	Totalday      string `json:"totalday,omitempty"`
	Humanreadday  string `json:"humanreadday,omitempty"`
	Second        string `json:"second,omitempty"`
	Minute        string `json:"minute,omitempty"`
	Hour          string `json:"hour,omitempty"`
	Week          string `json:"week,omitempty"`
	Percentofyear string `json:"percentofyear,omitempty"`
}
