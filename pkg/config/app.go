// Connect to the database

// Used MySQL here using gorm framework
package config

import (
	"fmt"
	"os"

	"github.com/Kaushal_26/go-bookstore/pkg/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func Connect() {
	current_working_directory, err := os.Getwd()
	utils.CheckError(err)
	godotenv.Load(current_working_directory + "..\\..\\.env")

	LOGIN_ID := os.Getenv("LOGIN_ID")
	PASSWORD := os.Getenv("PASSWORD")
	DBNAME := os.Getenv("DBNAME")
	gorm_connection := LOGIN_ID + ":" + PASSWORD + "@/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"

	d, err := gorm.Open("mysql", gorm_connection)
	utils.CheckError(err)

	fmt.Println("Database Connection Established\n")

	db = d
}

func GetDB() *gorm.DB {
	return db
}
