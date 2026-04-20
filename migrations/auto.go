package main

import (
	"os"
	"purple-school/internal/link"
	"purple-school/internal/user"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
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

	db.AutoMigrate(&link.Link{}, &user.User{})
}
