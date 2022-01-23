package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	Name  string
	Email string `gorm:"typevarchar(100); unique_index"`
	Books []Book
}

type Book struct {
	gorm.Model
	Title      string
	Author     string
	CallNumber int `gorm:"unique_index"`
	PersonID   int
}

func main() {
	// Loading env
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	dbport := os.Getenv("DBPORT")

	//Connection string
	dbURL := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbname, password, dbport)

	//Connect to database
	_, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connect to database")
	}
}
