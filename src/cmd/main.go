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
	log.Fatal(http.ListenAndServe(":8888", router))

}
