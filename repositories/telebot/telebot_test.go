package telebot

import (
	"testing"
)

func TestGetUpdates(t *testing.T) {
	obj := teleBotRepository{}
	t.Log(obj.GetUpdates())
}
