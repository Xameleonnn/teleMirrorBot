package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"teleBot/models"
)

// template for API "https://api.telegram.org/bot<token>/METHOD_NAME"

func token() string {
	token := flag.String(
		"token",
		"",
		"Token for telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("Missing token")
	}
	return *token
}

func errHandle(err error) {
	if err != nil {
		log.Println(err)
	}
}

func getUpdates(botApi string, offsetNum int) ([]models.Update, error) {
	resp, err := http.Get(botApi + "/getUpdates" + "?offset=" + strconv.Itoa(offsetNum))
	errHandle(err)
	defer resp.Body.Close()
	dataByte, err := ioutil.ReadAll(resp.Body)
	errHandle(err)
	var responseStruct models.RestResponse
	err = json.Unmarshal(dataByte, &responseStruct)
	errHandle(err)
	return responseStruct.Updates, nil
}

func respond(botApi string, update models.Update) error {
	incMessage := update.Message.Text
	chatId := update.Message.Chat.Chat_id
	var botMessage models.BotMessage
	botMessage.Chat_id = chatId
	botMessage.Text = incMessage
	buffer, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}
	_, err = http.Post(botApi+"/sendMessage", "application/json", bytes.NewBuffer(buffer))
	if err != nil {
		return err
	}
	return nil
}

func main() {
	t := token()
	botApi := "https://api.telegram.org/bot" + t
	offsetNum := 0
	for {
		updates, err := getUpdates(botApi, offsetNum)
		errHandle(err)
		for _, update := range updates {
			err = respond(botApi, update)
			errHandle(err)
			offsetNum = update.Update_id + 1
		}
		fmt.Println(updates)
	}
}
