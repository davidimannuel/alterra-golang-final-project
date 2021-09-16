package configs

import (
	"keep-remind-app/drivers/postgresql"
	"keep-remind-app/drivers/telebot"

	"keep-remind-app/server/middlewares"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Configs struct {
	AppHost    string
	AppTimeout time.Duration
	DB         *gorm.DB
	Redis      *redis.Client
	JWT        middlewares.ConfigJWT
	TeleBOT    *telebot.BotAPI
}

func LoadConfigs() (res Configs, err error) {
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
	configs.AppTimeout = time.Duration(viper.GetInt("APP_TIMEOUT")) * time.Second
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

	configs.Redis = redis.NewClient(&redis.Options{
		Addr:     viper.GetString("REDIS_HOST"),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       viper.GetInt("REDIS_DB"),
	})

	configs.JWT = middlewares.ConfigJWT{
		SecretJWT:       viper.GetString("JWT_SECRET"),
		ExpiresDuration: viper.GetInt("JWT_EXPIRED") * int(time.Second),
	}

	configs.TeleBOT = telebot.NewBot(viper.GetString("TELEGRAM_BOT_TOKEN"), telebot.UpdateConfig{
		Limit:   viper.GetInt("TELEGRAM_BOT_LIMIT"),
		Timeout: viper.GetInt("TELEGRAM_BOT_TIMEOUT"),
	})
	return configs, err
}
