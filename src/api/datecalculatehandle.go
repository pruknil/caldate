package api

import (
	"encoding/json"
	"net/http"
)

// our main function

func Calculate(w http.ResponseWriter, r *http.Request) {

	var resp Response
	//fmt.Printf(r.URL.Query().Get("point"))
	//resp.Point = r.URL.Query().Get("point")
	//resp.Grade = grade
	//resp.
	resp.Humanreadday = "hahah"
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
	Totalday      string `json:"to,omitempty"`
	Humanreadday  string `json:"humanreadday,omitempty"`
	Second        string `json:"second,omitempty"`
	Minute        string `json:"minute,omitempty"`
	Hour          string `json:"hour,omitempty"`
	Week          string `json:"week,omitempty"`
	Percentofyear string `json:"percentofyear,omitempty"`
}
