package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DbUsername = "admin"
const DbPassword = "password"
const DbName = "shortner"
const DbHost = "dpg-cild24lgkuvinfk5eqtg-a"
const DbPort = "3306"

var Db *gorm.DB

// InitDb - Initialize database
func InitDb() *gorm.DB {
	Db = connectDB()
	return Db
}

// connectDB - Connect to database
func connectDB() *gorm.DB {
	// dsn := DbUsername + ":" + DbPassword + "@tcp" + "(" + DbHost + ":" + DbPort + ")/" + DbName + "?" + "parseTime=true&loc=Local"
	dsn := "host=dpg-cild24lgkuvinfk5eqtg-a user=admin password=hVNuRlkX9RDx6gktWYJWtC5CRYWL7Lsd dbname=shortner port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Error connecting to database : error=%v\n", err)
		return nil
	}

	return db
}
