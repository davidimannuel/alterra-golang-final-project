package telebot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

type BotAPI struct {
	APIURL       string
	FilePathURL  string
	Token        string `json:"token"`
	Buffer       int    `json:"buffer"`
	UpdateConfig UpdateConfig
}

var (
	getMeMethodURL       = "/getMe"
	getUpdatesMethodURL  = "/getUpdates"
	getFileMethodURL     = "/getFile"
	sendMessageMethodURL = "/sendMessage"
)

func NewBot(token string, updateConfig UpdateConfig) *BotAPI {
	return &BotAPI{
		Token:       token,
		APIURL:      fmt.Sprintf("https://api.telegram.org/bot%s", token),
		FilePathURL: fmt.Sprintf("https://api.telegram.org/file/bot%s", token),
		UpdateConfig: UpdateConfig{
			Offset:  0,
			Limit:   updateConfig.Limit,
			Timeout: updateConfig.Timeout,
		},
	}
}

func (bot *BotAPI) MakeRequest(method string, params url.Values) (res APIResponse, err error) {
	apiURL := bot.APIURL + method + "?" + params.Encode()
	log.Println(apiURL)
	resp, err := http.Get(apiURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&res)
	if err != nil {
		return
	}
	log.Println(string(res.Result))
	return
}

func (bot *BotAPI) GetUpdates() (res []UpdatesResponse, err error) {
	params := url.Values{}
	if bot.UpdateConfig.Offset > 0 {
		params.Add("offset", strconv.Itoa(bot.UpdateConfig.Offset))
	}
	if bot.UpdateConfig.Limit > 0 {
		params.Add("limit", strconv.Itoa(bot.UpdateConfig.Limit))
	}
	if bot.UpdateConfig.Offset > 0 {
		params.Add("timeout", strconv.Itoa(bot.UpdateConfig.Timeout))
	}
	apiRes, err := bot.MakeRequest(getUpdatesMethodURL, params)
	err = json.Unmarshal(apiRes.Result, &res)
	if err != nil {
		return
	}
	return
}

func (bot *BotAPI) GetUpdatesChan() (UpdatesChan, error) {
	ch := make(chan UpdatesResponse, bot.Buffer)

	go func() {
		for {
			updates, err := bot.GetUpdates()
			if err != nil {
				log.Println(err)
				log.Println("Failed to get updates, retrying in 3 seconds...")
				time.Sleep(time.Second * 3)

				continue
			}
			for _, update := range updates {
				if update.UpdateID >= bot.UpdateConfig.Offset {
					bot.UpdateConfig.Offset = update.UpdateID + 1
					ch <- update
				}
			}
		}
	}()

	return ch, nil
}

func (bot *BotAPI) SendMessage(msgConfig SendMessageConfig) (res SendMessageResponse, err error) {
	params := url.Values{}
	params.Add("chat_id", strconv.Itoa(msgConfig.ChatID))
	params.Add("text", msgConfig.Text)
	apiRes, err := bot.MakeRequest(sendMessageMethodURL, params)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(apiRes.Result, &res)
	if err != nil {
		return
	}
	return
}

func (bot *BotAPI) GetFile(fileID string) (res GetFileResponse, err error) {
	params := url.Values{}
	params.Add("file_id", fileID)
	apiRes, err := bot.MakeRequest(getFileMethodURL, params)
	if err != nil {
		log.Println(err)
		return
	}
	err = json.Unmarshal(apiRes.Result, &res)
	if err != nil {
		return
	}
	return
}

func (bot *BotAPI) DownloadFile(filepath string) error {
	url := bot.FilePathURL + "/" + filepath
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func (bot *BotAPI) DownloadFileBytes(filepath string) (res []byte, err error) {
	url := bot.FilePathURL + "/" + filepath
	resp, err := http.Get(url)
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return res, err
	}
	res = buf.Bytes()

	return res, err
}
