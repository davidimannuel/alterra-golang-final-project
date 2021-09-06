package postgresql

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type Config struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func (config *Config) InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		config.DBHost,
		config.DBUsername,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error database connection")
	}
	return db
}
