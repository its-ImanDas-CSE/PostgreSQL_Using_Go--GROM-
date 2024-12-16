// ------------------------------GORM (Golang Object Relational Mapper) --------------------
package main

import (
	"fmt"

	"gorm.io/driver/postgres" // it enables communication with a PostgreSQL database.
	"gorm.io/gorm"            // GORM library, GORM library provides an easy way to interact with databases using objects instead of writing raw SQL queries.
)

// Define a struct (the blueprint of the table in the database).
type Student struct {
	ID     int    `gorm:"primaryKey"` // Primary key
	Name   string `gorm:"size:100"`   // String column with max size
	Age    int
	DOB    string
	Course string
	City   string
}

// In Go, the syntax func (Student) TableName() string { defines a method for the Student struct.
// It tells GORM how to map the Student struct to a specific table name in the database.
// Setting the Table Name
func (Student) TableName() string { // This function explicitly specifies that the table name in the database is student.
	return "student" // Explicitly set table name to "student"
}

/*
func (Student) TableName() string {
	return "employee" // Explicitly set table name to "employee", if we want to set the table name to employee.
}
*/

func main() {

	// Define PostgreSQL connection string
	dsn := "host = localhost user=postgres password=Virat@2#Virat@2# dbname=test port=8899 sslmode=disable"

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	fmt.Println("Connected to PostgreSQL with GORM!")

	// Automigrate: Create or insert row in the Database Schema(tables), if the table is created it will insert it or it will create a new table.
	// If the table doesn't exist, GORM will create it. If it exists, it will insert a new row or update its columns (if needed).
	err = db.AutoMigrate(&Student{}) // here student is struct that we created.
	if err != nil {
		panic("Failed to migrate database: " + err.Error())
	}

	fmt.Println("Database migrated successfully!")

	// inserting a new student record
	student := Student{ID: 112, Name: "Rohan", Age: 22, DOB: "2001-04-14", Course: "MBA", City: "Pune"}
	result := db.Create(&student) // Inserts the new student record into the student table.
	if result.Error != nil {
		panic("Failed to insert student: " + result.Error.Error())
	}

	fmt.Println("Inserted student successfully!")
	fmt.Println("Rows affected:", result.RowsAffected)
	// RowsAffected: Shows how many rows were modified or added (1 in this case since only one student is inserted).
	update()
}
