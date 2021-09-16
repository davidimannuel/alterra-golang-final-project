package main

import (
	_redisUc "keep-remind-app/businesses/redis"

	_teleBotUc "keep-remind-app/businesses/telebot"
	_telegramUserUc "keep-remind-app/businesses/telegramUser"
	_telegramUserRepo "keep-remind-app/repositories/telegramUser"

	"keep-remind-app/configs"
	_redisRepo "keep-remind-app/repositories/redis"

	"log"
)

func main() {

	config, err := configs.LoadConfigs()
	if err != nil {
		log.Println("error config", config)
	}
	bot := config.TeleBOT
	redisRepo := _redisRepo.NewRedisRepository(config.Redis)
	redisUc := _redisUc.NewRedisUsecase(redisRepo)
	telegramUserRepo := _telegramUserRepo.NewTelegramUserRepository(config.DB)
	telegramUserUc := _telegramUserUc.NewTelegramUserUsecase(telegramUserRepo, redisUc)
	telebotUc := _teleBotUc.NewTelebotUseCase(bot, redisUc, telegramUserUc)
	updates, _ := bot.GetUpdatesChan()
	for update := range updates {
		if update.Message.Text == "" && len(update.Message.Photo) == 0 {
			continue
		}
		log.Printf("[%s] %s,offset %d", update.Message.From.Username, update.Message.Text, bot.UpdateConfig.Offset)
		telebotUc.CommandManagement(update)
	}

}
