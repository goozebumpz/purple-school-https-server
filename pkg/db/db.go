package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"purple-school/configs"
)

type DB struct {
	db *gorm.DB
}

func CreateDbConnection(conf *configs.Config) *DB {
	db, err := gorm.Open(postgres.Open(conf.Db.DSN), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return &DB{db}
}
