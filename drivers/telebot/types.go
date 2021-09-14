package telebot

type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

type UpdatesChan <-chan UpdatesResponse

type SendMessageConfig struct {
	ChatID int
	Text   string
}
