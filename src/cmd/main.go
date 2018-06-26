package main

import (
	"api"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/duration", api.Calculate).Methods("POST")
	fs := http.FileServer(http.Dir("./html"))
	router.PathPrefix("/").Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fs.ServeHTTP(w, r)
	})).Methods("GET")

	log.Fatal(http.ListenAndServe(":8888", router))
}
