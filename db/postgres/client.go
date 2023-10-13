package postgres

import (
	"Loco/config"
	"Loco/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetConnection returns an instance of postgres database connection
func GetConnection() *gorm.DB {
	c := config.Postgres()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
		c.URL,
		c.User,
		c.Password,
		c.Database,
		c.Port,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("error in establishing connection with db", err)
		return nil
	}
	if err := db.AutoMigrate(&model.Transaction{}); err != nil {
		log.Fatal("error while creating transaction table")
	}
	return db
}
