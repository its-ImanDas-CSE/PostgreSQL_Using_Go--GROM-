package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* type Student struct {
	ID     int    `gorm:"primaryKey"`
	Name   string `gorm:"size:100"`
	Age    int
	DOB    string
	Course string
	City   string
}   */

/* func (Student) TableName() string {
	return "student" // Explicitly set table name to "student"
} */

func update() {
	// Define PostgreSQL connection string
	dsn := "host=localhost user=postgres password=Virat@2#Virat@2# dbname=test port=8899 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Ask the user for the ID of the student whose age they want to update
	var userID int
	fmt.Print("Enter student ID to update age: ")
	_, err = fmt.Scan(&userID) // Read user input for student ID
	if err != nil {
		log.Fatal("Invalid input for student ID: ", err)
	}

	// Find the student by their ID
	var student Student                  // Declaring the student variable type Student(struct).
	result := db.First(&student, userID) // This is the GORM query that fetches the first record in the student table that matches the given condition.
	// The '&' symbol is used to pass the address of the student variable (a pointer), so that GORM can fill it with the data retrieved from the database.
	// The fetch data is stored in the student variable.
	if result.Error != nil {
		fmt.Println("Student not found!")
		return
	}

	// Ask the user for the new age
	var newAge int
	fmt.Print("Enter new age for student: ")
	_, err = fmt.Scan(&newAge) // Read user input for new age
	if err != nil {
		log.Fatal("Invalid input for age: ", err)
	}

	// Update the student's age, it is the student variable that we created above, to store the fetch data.
	student.Age = newAge

	// Save the updated student variable record to the database
	result = db.Save(&student) // This is the GORM method used to save the student record to the database.
	if result.Error != nil {
		fmt.Println("Failed to update student:", result.Error)
		return
	}

	// Success message
	fmt.Println("Student age updated successfully!") // successfully done.
}
