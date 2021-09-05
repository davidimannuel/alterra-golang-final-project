package orm

import (
	"fmt"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

type ORM struct {
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
}

func (orm *ORM) InitDB() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=Asia/Jakarta",
		orm.DBHost,
		orm.DBUsername,
		orm.DBPassword,
		orm.DBName,
		orm.DBPort,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Error orm connection")
	}
	return db
}
