package configs

import (
	"alterra-golang-final-project/drivers/postgresql"
	"log"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Configs struct {
	AppHost string
	DB      *gorm.DB
}

func LoadConfigs() (res *Configs, err error) {
	// read config
	viper.SetConfigFile("../.env")
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("Config file not found")
		} else {
			// Config file was found but another error was produced
			log.Println(err.Error())
		}
	}
	// init config
	configs := Configs{}
	// app host
	configs.AppHost = viper.GetString("APP_HOST")
	//init db
	db := postgresql.Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBUsername: viper.GetString("DB_USERNAME"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBPort:     viper.GetString("DB_PORT"),
		DBName:     viper.GetString("DB_NAME"),
	}
	configs.DB = db.InitDB()
	postgresql.MigrateDB(configs.DB)
	return &configs, err
}
