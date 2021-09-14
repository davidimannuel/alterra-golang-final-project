package telebot

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

var botUrl = "https://api.telegram.org/bot1994661667:AAGrNP4X_tO3HPdMJBw47ytZ2q7ikVZX8VA"

type teleBotRepository struct {
}

// UpdateConfig contains information about a GetUpdates request.
type UpdateConfig struct {
	Offset  int
	Limit   int
	Timeout int
}

func (repo teleBotRepository) GetUpdates() (res TeleBotResponse, err error) {
	url := botUrl + "/getupdates"
	// offset = last_id + 1
	resp, err := http.Get(url)
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	log.Println(string(body))
	err = json.Unmarshal(body, &res)
	if err != nil {
		return
	}
	return
}
