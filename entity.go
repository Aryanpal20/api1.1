// here we can create employee entity

package main

type Employee struct {
	Empid     int    `json:"empid"`
	Empname   string `json:"empname"`   // json format for key value pair
	Empsalary int    `json:"empsalary"` // json format for key value pair
	// Email     string `json:"email"`     // json format for key value pair

}
