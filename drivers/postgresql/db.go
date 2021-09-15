package postgresql

import (
	"fmt"
	"keep-remind-app/repositories/label"
	"keep-remind-app/repositories/note"
	"keep-remind-app/repositories/telegramUser"
	"keep-remind-app/repositories/user"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second * 20, // Slow SQL threshold
			LogLevel:                  logger.Silent,    // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)
	db, err := gorm.Open(postgres.New(
		postgres.Config{
			DSN: dsn,
		}), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal("Error database connection")
	}

	return db
}

func MigrateDB(db *gorm.DB) {
	db.AutoMigrate(&user.UserModel{}, &note.NoteModel{}, &label.LabelModel{}, &telegramUser.TelegramUserModel{}) // migrate db
}
