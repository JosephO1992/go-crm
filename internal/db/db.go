package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Client *gorm.DB
}

func NewDatabase() (*Database, error) {

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return &Database{}, fmt.Errorf("could not connect to database %w", err)
	}

	fmt.Println("Connected to the database	")

	return &Database{
		Client: db,
	}, nil
}

func MigrateDatabase(d *Database) error {
	fmt.Println("Migrating our database")
	var err = d.Client.AutoMigrate(&User{})
	return err
}
