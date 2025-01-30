package infrastructure

import (
	"fmt"
	"go-http-server/utils"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres(config *utils.Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbPort, config.DbUser, config.DbPassword, config.DbName)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	return db
}
