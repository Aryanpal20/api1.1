package main

type Manager struct {
	Mrg_Id   int    `json:"mrg_id"`
	Task     string `json:"task"`
	Date     string `json:"date"`
	Due_Date string `json:"due_date"`
}
