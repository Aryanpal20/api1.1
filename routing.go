package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()

	r.HandleFunc("/infos", GetInformations).Methods("GET")
	r.HandleFunc("/info/{userid}", GetInformationById).Methods("GET")
	r.HandleFunc("/info", CreateInformation).Methods("POST")
	r.HandleFunc("/info/{userid}", UpdateInformation).Methods("PUT")
	r.HandleFunc("/info/{userid}", DeleteInformation).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
