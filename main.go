package main

import (
	"fmt"
	"log"

	"rest_echo/api/models"
	"rest_echo/bootstrap"
	"rest_echo/db/gorm"
	"rest_echo/router"
)

func main() {
	fmt.Println("Welcome to the server")

	log.Printf("ENV : %s", bootstrap.App.ENV)

	e := router.New()

	// init database
	gorm.Init()
	// autoDropTables()
	// autoCreateTables()
	autoMigrateTables()

	e.Start(":8000")
}

// autoCreateTables: create database tables using GORM
func autoCreateTables() {
	if !gorm.DBManager().HasTable(&models.User{}) {
		gorm.DBManager().CreateTable(&models.User{})
	}

	// seeder
	if bootstrap.App.ENV == "dev" {
		// var users []models.User = []models.User{
		// 	models.User{Name: "iman", Email: "iman@sepulsa.com", Password: "123", Roleid: 1, Address: "kopo"},
		// }

		// for _, user := range users {
		// 	gorm.DBManager().Create(&user)
		// }
	}
}

// autoMigrateTables: migrate table columns using GORM
func autoMigrateTables() {
	gorm.DBManager().AutoMigrate(&models.User{})
}

// auto drop tables on dev mode
func autoDropTables() {
	if bootstrap.App.ENV == "dev" {
		gorm.DBManager().DropTableIfExists(&models.User{}, &models.User{})
	}
}
