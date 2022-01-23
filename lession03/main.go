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

var (
	person = &Person{
		Name:  "Jason",
		Email: "me@nguyenthdat.com",
	}
	books = []Book{
		{Title: "a", Author: "b", CallNumber: 1, PersonID: 1},
		{Title: "c", Author: "d", CallNumber: 2, PersonID: 1},
		{Title: "e", Author: "f", CallNumber: 3, PersonID: 1},
	}
)

func main() {
	// Loading env
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	dbport := os.Getenv("DBPORT")

	//Connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbport)

	//Connect to database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Successfully connect to database")
	}
	db.AutoMigrate(&Person{})
	db.AutoMigrate(&Book{})
	db.Create(&person)
	for i := range books {
		db.Create(&books[i])
	}
}
