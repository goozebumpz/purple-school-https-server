package main

import (
	"fmt"
	"os"
	"purple-school/internal/link"
	"purple-school/pkg/db"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	const f = "main"
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(postgres.Open(os.Getenv("DSN")), &gorm.Config{})
	sqlDB, err := db.DB()
	defer sqlDB.Close()

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&link.Link{})
}

func createDatabaseLink(db *db.DB, name string) error {
	res := db.Exec(fmt.Sprintf(`CREATE DATABASE %s`, name))
	err := res.Error

	if err != nil {
		isExitError :=
			err.Error() == fmt.Sprintf(`pq: database "%s" already exists`, name) ||
				err.Error() == fmt.Sprintf(`ERROR: database "%s" already exists (SQLSTATE 42P04)`, name)

		if isExitError {
			return nil
		}

		return err
	}

	return nil
}
