package main

import (
	_labelUc "keep-remind-app/businesses/label"
	_ocrUc "keep-remind-app/businesses/ocr"
	_redisUc "keep-remind-app/businesses/redis"
	_labelRepo "keep-remind-app/repositories/label"

	_noteUc "keep-remind-app/businesses/note"
	_teleBotUc "keep-remind-app/businesses/telebot"
	_telegramUserUc "keep-remind-app/businesses/telegramUser"
	_noteRepo "keep-remind-app/repositories/note"
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
	telegramUserRepo := _telegramUserRepo.NewTelegramUserRepository(config.DB)
	labelRepo := _labelRepo.NewLabelRepository(config.DB)
	noteRepo := _noteRepo.NewNoteRepository(config.DB)
	redisUc := _redisUc.NewRedisUsecase(redisRepo)
	ocrUc := _ocrUc.NewOCRUsecase()
	telegramUserUc := _telegramUserUc.NewTelegramUserUsecase(telegramUserRepo, redisUc)
	labelUc := _labelUc.NewLabelUsecase(labelRepo)
	noteUc := _noteUc.NewNoteUsecase(noteRepo, ocrUc, labelUc)
	telebotUc := _teleBotUc.NewTelebotUseCase(bot, redisUc, telegramUserUc, noteUc)
	updates, _ := bot.GetUpdatesChan()
	for update := range updates {
		if update.Message.Text == "" && len(update.Message.Photo) == 0 {
			continue
		}
		log.Printf("[%s] %s,offset %d", update.Message.From.Username, update.Message.Text, bot.UpdateConfig.Offset)
		telebotUc.CommandManagement(update)
	}

}
