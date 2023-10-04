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

	dbURL := "postgres://unreasonable2003:rfj4LJm5nTpR@ep-bold-bread-62779844.ap-southeast-1.aws.neon.tech/neondb"

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
}
