// here we can do database connection

package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// global variable for creating emp and access emp

var Database *gorm.DB

var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/detail" // urldsn is a driver where the database connected  127.0.0.1:3306
var err error

func DataMigration() {

	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())

		panic("connection failed")
	}
	Database.AutoMigrate(Employee{})
}
