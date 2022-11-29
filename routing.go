package main

import (
	"log"

	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()

	r.HandleFunc("/mrg", GetInfos).Methods("GET")
	r.HandleFunc("/mrg/{mrg_id}", GetInfoById).Methods("GET")
	r.HandleFunc("/mrg", CreateInfo).Methods("POST")
	r.HandleFunc("/mrg/{mrg_id}", UpdateInfo).Methods("PUT")
	r.HandleFunc("/mrg/{mrg_id}", DeleteInfo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
