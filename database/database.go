package database

import (
	"fmt"
	"os"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() *gorm.DB {
	DBHost := os.Getenv("DB_HOST")
	DBUser := os.Getenv("DB_USER")
	DBPassword := os.Getenv("DB_PASS")
	DBName := os.Getenv("DB_NAME")
	DBPort := os.Getenv("DB_PORT")

	DBDSN := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		DBHost, DBUser, DBPassword, DBName, DBPort,
	)

	db, err := gorm.Open(postgres.Open(DBDSN), &gorm.Config{})
	if err != nil {
		panic("failed connect to database")
	}

	if err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`).Error; err != nil {
		panic(err)
	}

	return db
}

func Commands(db *gorm.DB) {
	migrate := false

	for _, arg := range os.Args[1:] {
		if arg == "--migrate" {
			migrate = true
			break
		}
	}

	if migrate {
		err := db.AutoMigrate(
			entity.User{},
			entity.Post{},
			entity.PostImage{},
			entity.Comment{},
			entity.Like{},
		)
		if err != nil {
			panic("Failed to migrate database: " + err.Error())
		}

		fmt.Println("[MIGRATION] success migrate")
	}
}
