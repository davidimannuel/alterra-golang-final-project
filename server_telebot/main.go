package main

import (
	"keep-remind-app/configs"
	"keep-remind-app/drivers/ocr"
	"keep-remind-app/drivers/telebot"
	"log"
)

func main() {

	config, err := configs.LoadConfigs()
	if err != nil {
		log.Println("error config", config)
	}
	bot := config.TeleBOT

	updates, _ := bot.GetUpdatesChan()
	for update := range updates {
		if update.Message.Text == "" && len(update.Message.Photo) == 0 {
			continue
		}

		log.Printf("[%s] %s,offset %d", update.Message.From.Username, update.Message.Text, bot.UpdateConfig.Offset)
		if len(update.Message.Photo) > 0 {
			bot.SendMessage(telebot.SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , you are sending file ",
			})
			file, err := bot.GetFile(update.Message.Photo[len(update.Message.Photo)-1].FileID)
			if err != nil {
				bot.SendMessage(telebot.SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   "hello from bot , sorry failed to get file ",
				})
			}
			bytes, _ := bot.DownloadFileBytes(file.FilePath)
			text, _ := ocr.GetImageTextFromImageBytes(bytes)
			bot.SendMessage(telebot.SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , your text image note \n" + text,
			})
		} else {
			bot.SendMessage(telebot.SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , you are typing ," + update.Message.Text,
			})
		}
	}
}
