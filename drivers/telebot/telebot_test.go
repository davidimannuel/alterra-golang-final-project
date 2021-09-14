package telebot

import (
	"keep-remind-app/drivers/ocr"
	"log"
	"net/url"
	"testing"
)

func TestMakeRequest(t *testing.T) {
	bot := NewBot("1994661667:AAGrNP4X_tO3HPdMJBw47ytZ2q7ikVZX8VA", UpdateConfig{
		Limit:   20,
		Timeout: 20,
	})
	param := url.Values{}
	param.Add("offset", "1")
	t.Log(param.Encode())
	t.Log(bot.MakeRequest(getMeMethodURL, param))
}
func TestGetUpdatesChan(t *testing.T) {
	bot := NewBot("1994661667:AAGrNP4X_tO3HPdMJBw47ytZ2q7ikVZX8VA", UpdateConfig{
		Limit:   20,
		Timeout: 20,
	})

	updates, _ := bot.GetUpdatesChan()
	for update := range updates {
		if update.Message.Text == "" && len(update.Message.Photo) == 0 {
			continue
		}

		log.Printf("[%s] %s,offset %d", update.Message.From.Username, update.Message.Text, bot.UpdateConfig.Offset)
		if len(update.Message.Photo) > 0 {
			bot.SendMessage(SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , you are sending file ",
			})
			file, err := bot.GetFile(update.Message.Photo[len(update.Message.Photo)-1].FileID)
			if err != nil {
				bot.SendMessage(SendMessageConfig{
					ChatID: update.Message.Chat.ID,
					Text:   "hello from bot , sorry failed to get file ",
				})
			}
			bytes, _ := bot.DownloadFileBytes(file.FilePath)
			text, _ := ocr.GetImageTextFromImageBytes(bytes)
			bot.SendMessage(SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , your image note " + text,
			})
		} else {
			bot.SendMessage(SendMessageConfig{
				ChatID: update.Message.Chat.ID,
				Text:   "hello from bot , you are typing ," + update.Message.Text,
			})
		}
	}

}
