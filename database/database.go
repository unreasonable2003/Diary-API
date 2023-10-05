package database

import (
	"fmt"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Connect() {
	var err error

	dbURL := "postgresql://example"

	// Parse the URL
	connector, err := pq.ParseURL(dbURL)
	if err != nil {
		panic(err)
	}

	// Open a connection to the database
	Database, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  connector,
		PreferSimpleProtocol: true, // Set this to true if needed
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		fmt.Println("Successfully connected to the database")
	}

	// Check the database connection
	err = pingDatabase(Database)
	if err != nil {
		panic(err)
	}
}

// Function to ping the database to check the connection
func pingDatabase(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		return err
	}

	fmt.Println("Connected to the database")
	return nil
}
