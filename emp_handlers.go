// here we can create a handler in database

package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employees []Employee             // here we can use array beacuse we want list of employee
	Database.Find(&employees)            // here fetch all the recodrs from database
	json.NewEncoder(w).Encode(employees) // here we can return employee

}

func GetEmployeeById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"]) // here we can fetch the data by id
	json.NewEncoder(w).Encode(employee)           // here we can return employee

}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var employee Employee
	Database.First(&employee, mux.Vars(r)["eid"]) //  here we can feth the data by id
	json.NewDecoder(r.Body).Decode(&employee)     // here we can update the data which we can fetch
	Database.Save(&employee)                      // here we can save the data in table which we can update
	json.NewEncoder(w).Encode(employee)           // here we can return the data

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var emp Employee
	Database.Delete(&emp, mux.Vars(r)["eid"])
	json.NewEncoder(w).Encode("employee is deleted ")

}
