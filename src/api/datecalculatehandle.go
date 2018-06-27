package api

import (
	"datecalculate"
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
	resp := datecalculate.CalDuration(req.StartDate, req.EndDate)
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
