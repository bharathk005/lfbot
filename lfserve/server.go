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

	if err != nil {
		log.Printf("error parsing update, %s", err.Error())
		resp = "<BOT> Ahh! I can listen to texts only.. Sorry, "
		sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
	} else if len(update.Message.Text) < 1 {

	} else if update.Message.Text[0] == '/' {
		if update.Message.Text[1:] == "start" {
			resp = "Welcome! This is a safe place to chat with Random people Anonymously!" +
				"\nCommands: \n/new to start a new chat\n/like to let the other person know that you like the chat\nKeep it Simple!"
			sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
			pair := GetPair(update.Message.Chat.Id)
			if pair == -1 {
				resp = "<BOT> I am currently waiting for a Random Sapien to join! You will recieve a ping when someone joins"
				sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
				findNewPair(update.Message.Chat.Id, update.Message.Chat.FirstName)
			} else {
				resp = "<BOT> I have found a person you can chat with. Start Chatting! "
				sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
				sendTextToTelegramChat(pair, resp, update.Message.Chat.FirstName)
			}
		} else if update.Message.Text[1:] == "new" {
			pair, affected := NewPair(update.Message.Chat.Id)
			if pair == -1 {
				resp = "<BOT> I am currently waiting for a Random Sapien to join! You will recieve a ping when someone joins"
				sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
			} else {
				resp = "<BOT> I have found a new person you can chat with. Start Chatting! "
				sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
				sendTextToTelegramChat(pair, resp, update.Message.Chat.FirstName)
			}

			// notify pair that conv ended
			if affected != -1 {
				resp = "<BOT> I am sad to inform you that the other sapien left. I am finding a different sapien for you.. "
				sendTextToTelegramChat(affected, resp, update.Message.Chat.FirstName)
				findNewPair(affected, update.Message.Chat.FirstName)
			}

		} else if update.Message.Text[1:] == "like" {
			pair := GetPair(update.Message.Chat.Id)
			if pair == -1 {
				resp = "<BOT> Sorry no one is listening"
				sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
			} else {
				resp = "The other person likes ❤ this chat!\n/like to like back"
				sendTextToTelegramChat(pair, resp, update.Message.Chat.FirstName)
			}
		}
	} else {
		resp = update.Message.Text
		pair := GetPair(update.Message.Chat.Id)
		if pair == -1 {
			resp = "<BOT> I am currently waiting for a Random Sapien to join! You will recieve a ping when someone joins "
			sendTextToTelegramChat(update.Message.Chat.Id, resp, update.Message.Chat.FirstName)
			findNewPair(update.Message.Chat.Id, update.Message.Chat.FirstName)
		} else {
			sendTextToTelegramChat(pair, resp, update.Message.Chat.FirstName)
		}
	}
}

func findNewPair(chatId int64, name string) {
	pair, _ := NewPair(chatId)
	if pair != -1 {
		resp := "<BOT> I have found a person you can chat with. Start Chatting! "
		sendTextToTelegramChat(chatId, resp, name)
		sendTextToTelegramChat(pair, resp, "")
	}
}

func sendTextToTelegramChat(chatId int64, text string, name string) {

	if false {
		var telegramApi string = "https://api.telegram.org/bot" + os.Getenv("TMP_TOKEN") + "/sendMessage"
		response, err := http.PostForm(
			telegramApi,
			url.Values{
				"chat_id": {strconv.FormatInt(chatId, 10)},
				"text":    {text},
			})

		if err != nil {
			log.Printf("error when posting text to the chat: %s", err.Error())
			return
		}
		defer response.Body.Close()

		var _, errRead = ioutil.ReadAll(response.Body)
		if errRead != nil {
			log.Printf("error in parsing telegram answer %s", errRead.Error())
			return
		}
	}
	log.Printf("%d:%s -- %s", chatId, name, text)
}
