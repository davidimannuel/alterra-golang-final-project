package telebot

import (
	"context"
	redisDomain "keep-remind-app/businesses/redis"
	"keep-remind-app/businesses/telegramUser"
	"keep-remind-app/drivers/ocr"
	"keep-remind-app/drivers/telebot"
	"log"

	"github.com/spf13/cast"
)

type TeleBotUsecase struct {
	bot          *telebot.BotAPI
	res          telegramUser.TelegramUserUsecase
	telegramUser telegramUser.TelegramUserUsecase
	redisUsecase redisDomain.RedisUsecase
}

var (
	commandsWhiteList = []string{
		"/start", "/login", "/notes",
	}

	responseCommands = map[string]string{
		"/start": "Please Login First",
		"/login": "Please Input Your OTP",
		"/notes": "Here Your Notes : \n",
	}
)

func NewTelebotUseCase(bot *telebot.BotAPI, redisUsecase redisDomain.RedisUsecase, telegramUser telegramUser.TelegramUserUsecase) *TeleBotUsecase {
	return &TeleBotUsecase{
		bot:          bot,
		telegramUser: telegramUser,
		redisUsecase: redisUsecase,
	}
}

func (uc *TeleBotUsecase) CommandManagement(update telebot.UpdatesResponse) {
	msg := update.Message.Text
	username := update.Message.From.Username
	latestAction, _ := uc.redisUsecase.Get(context.Background(), telegramUser.LastActionTelegram+username)
	if latestAction == telegramUser.LatestActionOtp {
		err := uc.telegramUser.Activated(context.Background(), username, msg)
		log.Println(err)
		message := "Success to Login"
		if err != nil {
			message = "Sorry failed to login"
		}
		uc.bot.SendMessage(telebot.SendMessageConfig{
			ChatID: update.Message.Chat.ID,
			Text:   message,
		})
	} else {
		message, exist := responseCommands[msg]
		if !exist {
			message = "Sorry, command doest exist"
		}
		uc.bot.SendMessage(telebot.SendMessageConfig{
			ChatID: update.Message.Chat.ID,
			Text:   message,
		})
	}
	// save last chat id from users
	uc.redisUsecase.Set(context.Background(), telegramUser.LastChatTelegram+update.Message.From.Username, cast.ToInt(update.Message.Chat.ID), 0)
	return
}

func (uc *TeleBotUsecase) SaveNoteFromImage(update telebot.UpdatesResponse) {
	file, err := uc.bot.GetFile(update.Message.Photo[len(update.Message.Photo)-1].FileID)
	if err != nil {
		uc.bot.SendMessage(telebot.SendMessageConfig{
			ChatID: update.Message.Chat.ID,
			Text:   "hello from bot , sorry failed to get file ",
		})
	}
	bytes, _ := uc.bot.DownloadFileBytes(file.FilePath)
	text, _ := ocr.GetImageTextFromImageBytes(bytes)
	uc.bot.SendMessage(telebot.SendMessageConfig{
		ChatID: update.Message.Chat.ID,
		Text:   "note saved , your text image note \n" + text,
	})
}
