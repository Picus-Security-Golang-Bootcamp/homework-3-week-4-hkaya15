package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type POSTGRES struct {}

func (ps *POSTGRES) Create() (*gorm.DB, error) {
	res, err := CreatePostgreSQL()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func CreatePostgreSQL() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("PICUS_DB_HOST"),
		os.Getenv("PICUS_DB_PORT"),
		os.Getenv("PICUS_DB_USERNAME"),
		os.Getenv("PICUS_DB_NAME"),
		os.Getenv("PICUS_DB_PASSWORD"),
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true})
	if err != nil {
		return nil, fmt.Errorf("cannot open database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
