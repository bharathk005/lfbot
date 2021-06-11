package lfserve

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func parseRequest(r *http.Request) (*Update, error) {
	var update Update
	if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
		log.Printf("could not decode incoming update %s", err.Error())
		return nil, err
	}
	return &update, nil
}

func HandleTelegramWebHook(w http.ResponseWriter, r *http.Request) {

	// Parse incoming request

	var update, err = parseRequest(r)
	resp := ""
	dest := update.Message.Chat.Id
	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		resp = "Ahh! Snap. Somthing is wrong in here.. Sorry " + update.Message.Chat.FirstName
	} else if update.Message.Text[0] == '/' {
		if update.Message.Text[1:] == "start" {
			resp = "Welcome! This is a safe place to chat with Random people Anonymously!" +
				"\nCommands: \n/new to start a new chat\n/like to let the other person know that you like the chat\nKeep it Simple!"
		} else if update.Message.Text[1:] == "new" {
			// logic for new chat
			// notify pair that conv ended
		} else if update.Message.Text[1:] == "like" {
			// logic for liking the chat
			// notify pair the conv liked
		}
	} else {
		resp = update.Message.Text
		pair := GetPair(update.Message.Chat.Id)
		if pair == -1 {
			resp = "We are currently waiting for a Random Sapien!"
		} else {
			dest = pair
		}
	}

	var telegramResponseBody, errTelegram = sendTextToTelegramChat(dest, resp)
	if errTelegram != nil {
		log.Printf("got error %s from telegram, reponse body is %s", errTelegram.Error(), telegramResponseBody)
	} else {
		log.Printf("%s -- %d", resp, update.Message.Chat.Id)
	}
}

func sendTextToTelegramChat(chatId int, text string) (string, error) {

	var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TMP_TOKEN") + "/sendMessage"
	response, err := http.PostForm(
		telegramApi,
		url.Values{
			"chat_id": {strconv.Itoa(chatId)},
			"text":    {text},
		})

	if err != nil {
		log.Printf("error when posting text to the chat: %s", err.Error())
		return "", err
	}
	defer response.Body.Close()

	var bodyBytes, errRead = ioutil.ReadAll(response.Body)
	if errRead != nil {
		log.Printf("error in parsing telegram answer %s", errRead.Error())
		return "", err
	}
	bodyString := string(bodyBytes)

	return bodyString, nil
}
