package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const DbUsername = "admin"
const DbPassword = "password"
const DbName = "shortner"
const DbHost = "localhost"
const DbPort = "3306"

var Db *gorm.DB

// InitDb - Initialize database
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

// connectDB - Connect to database
func connectDB() *gorm.DB {
	dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"
	fmt.Println("dsn : ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}

	return db
}
