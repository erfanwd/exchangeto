package database

import (
	"fmt"
	"log"
	"telegram-todolist/config"
	"telegram-todolist/models"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", config.Config("POSTGRES_HOST"), config.Config("POSTGRES_USER"), config.Config("POSTGRES_PASSWORD"), config.Config("POSTGRES_DATABASE_NAME"), config.Config("POSTGRES_PORT"))
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB.AutoMigrate(&models.Task{}, &models.User{}, &models.Exchange{}, &models.Reminder{})

	return DB
}
