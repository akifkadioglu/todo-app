package controllers

import (
	"database/sql"
	"fmt"
	config "myapp/config"
	"myapp/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Query struct {
	Exist string `json:"exist"`
}

var dbAddress = config.GoDotEnvVariable("DB_USERNAME") + ":" +
	config.GoDotEnvVariable("DB_PASSWORD") +
	"@tcp(" +
	config.GoDotEnvVariable("DB_HOST") +
	":" +
	config.GoDotEnvVariable("DB_PORT") +
	")/"

func ConnectDatabase() gorm.DB {
	check := CheckDatabase()
	if check != "1" {
		CreateDatabase()
	}
	dsn := dbAddress + config.GoDotEnvVariable("DB_DATABASE") + "?charset=utf8mb4&parseTime=True&loc=Local"
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn))
	db.AutoMigrate(&models.Task{})
	if err != nil {
		panic(err.Error())
	}
	return *db

}

func CreateDatabase() {
	result := CheckDatabase()
	if result != "1" {
		db := DBConnection()
		_, err := db.Exec("CREATE DATABASE `" + config.GoDotEnvVariable("DB_DATABASE") + "`;")
		if err != nil {
			panic(err.Error())
		}
		fmt.Println("The (" + config.GoDotEnvVariable("DB_DATABASE") + ") database created")
	} else {
		fmt.Println("The (" + config.GoDotEnvVariable("DB_DATABASE") + ") database already exist.")
	}
}

func CheckDatabase() string {
	db := DBConnection()
	query, _ := db.Query("SELECT IF(EXISTS( SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '" + config.GoDotEnvVariable("DB_DATABASE") + "'), 1, 0) as exist")
	query.Next()
	var result Query
	err := query.Scan(&result.Exist)
	if err != nil {
		panic(err.Error())
	}
	return result.Exist
}

func DBConnection() sql.DB {
	db, err := sql.Open(config.GoDotEnvVariable("DB_CONNECTION"), dbAddress)
	if err != nil {
		panic(err.Error())
	}
	return *db
}
