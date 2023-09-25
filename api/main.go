package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

const SOURCE_CODE_URL = "https://github.com/st-matskevich/tg-mini-app-template/"

func main() {
	log.Println("Starting API service")

	webAppURL := os.Getenv("TELEGRAM_WEB_APP_URL")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	bot, err := gotgbot.NewBot(botToken, nil)
	if err != nil {
		log.Fatalf("Telegram Bot API initialization error: %v", err)
	}
	log.Println("Telegram Bot API initialized")

	http.HandleFunc("/bot", CreateBotEndpointHandler(bot, webAppURL))

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// According to the https://core.telegram.org/bots/api#setwebhook webhook will receive JSON-serialized Update structure
// Handler created by this function parses Update structure and replies to any message with welcome text and inline keyboard to open Mini App
func CreateBotEndpointHandler(bot *gotgbot.Bot, appURL string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s route", r.URL.Path)
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusNotImplemented)
			return
		}

		update := gotgbot.Update{}
		err := json.NewDecoder(r.Body).Decode(&update)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if update.Message == nil {
			http.Error(w, "Bot update didn't include a message", http.StatusBadRequest)
			return
		}

		message := "Welcome to the Telegram Mini App Template Bot\nTap the button below to open mini app or bot source code"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: appURL}},
				}, {
					{Text: "Open source code", Url: SOURCE_CODE_URL},
				}},
			},
		}

		if _, err := bot.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	}
}
