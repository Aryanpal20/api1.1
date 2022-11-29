package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateInfo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var info Manager
	json.NewDecoder(r.Body).Decode(&info)
	Database.Create(&info)
	json.NewEncoder(w).Encode(info)

}

func GetInfos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var info []Manager
	Database.Find(&info)
	json.NewEncoder(w).Encode(info)

}

func GetInfoById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var info Manager
	Database.First(&info, mux.Vars(r)["mrg_id"])
	json.NewEncoder(w).Encode(info)

}

func UpdateInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var info Manager
	Database.First(&info, mux.Vars(r)["mrg_id"])
	json.NewDecoder(r.Body).Decode(&info)
	if err := Database.Model(&info).Where("mrg_id = ?", info.Mrg_Id).Update("task", info.Task).Error; err != nil {
		fmt.Printf("update err != nil; %v\n", err)
	}
	if err := Database.Model(&info).Where("mrg_id = ?", info.Mrg_Id).Update("date", info.Date).Error; err != nil {
		fmt.Printf("update err != nil; %v\n", err)
	}
	if err := Database.Model(&info).Where("mrg_id = ?", info.Mrg_Id).Update("due_date", info.Due_Date).Error; err != nil {
		fmt.Printf("update err != nil; %v\n", err)
	}
	// if err := Database.Model(&info).Where("user_id = ?", info.UserId).Update("email", info.Email).Error; err != nil {
	// 	fmt.Printf("update err != nil; %v\n", err)
	// }
	Database.Save(info)
	json.NewEncoder(w).Encode(info)

}

func DeleteInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var info Manager
	Database.Delete(&info, mux.Vars(r)["mrg_id"])
	json.NewEncoder(w).Encode("Information is deleted ")

}
