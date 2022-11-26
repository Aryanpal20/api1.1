package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateInformation(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var info Information
	json.NewDecoder(r.Body).Decode(&info)
	Database.Create(&info)
	json.NewEncoder(w).Encode(info)

}

func GetInformations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var information []Information
	Database.Find(&information)
	json.NewEncoder(w).Encode(information)

}

func GetInformationById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var information Information
	Database.First(&information, mux.Vars(r)["userid"])
	json.NewEncoder(w).Encode(information)

}

func UpdateInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var information Information
	Database.First(&information, mux.Vars(r)["userid"])
	json.NewDecoder(r.Body).Decode(&information)
	if err := Database.Model(&information).Where("user_id = ?", information.UserId).Update("task", information.Task).Error; err != nil {
		fmt.Printf("update err != nil; %v\n", err)
	}
	if err := Database.Model(&information).Where("user_id = ?", information.UserId).Update("date", information.Date).Error; err != nil {
		fmt.Printf("update err != nil; %v\n", err)
	}
	Database.Save(information)
	json.NewEncoder(w).Encode(&information)

}

func DeleteInformation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var information Information
	Database.Delete(&information, mux.Vars(r)["userid"])
	json.NewEncoder(w).Encode("information is deleted ")

}
