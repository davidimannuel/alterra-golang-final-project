package postgresql

import (
	"fmt"
	"keep-remind-app/repositories/note"
	"keep-remind-app/repositories/user"
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
		"postgres://%s:%s@%s:%s/%s?sslmode=disable&TimeZone=UTC",
		config.DBUsername,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		}), &gorm.Config{})
	if err != nil {
		log.Fatal("Error database connection")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.Model{}, &note.Model{}) // migrate db
}
