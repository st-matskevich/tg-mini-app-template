package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/PaulSonOfLars/gotgbot/v2"
)

func main() {
	log.Println("Starting API service")

	webAppURL := os.Getenv("TELEGRAM_WEB_APP_URL")
	botToken := os.Getenv("TELEGRAM_BOT_TOKEN")
	botClient, err := gotgbot.NewBot(botToken, nil)
	if err != nil {
		log.Fatalf("Telegram Bot API initialization error: %v", err)
	}
	log.Println("Telegram Bot API initialized")

	http.HandleFunc("/bot", func(w http.ResponseWriter, r *http.Request) {
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

		message := "Welcome to the Telegram Mini App Template Bot\nTap the button below to open mini app"
		opts := &gotgbot.SendMessageOpts{
			ReplyMarkup: gotgbot.InlineKeyboardMarkup{
				InlineKeyboard: [][]gotgbot.InlineKeyboardButton{{
					{Text: "Open mini app", WebApp: &gotgbot.WebAppInfo{Url: webAppURL}},
				}},
			},
		}

		if _, err := botClient.SendMessage(update.Message.Chat.Id, message, opts); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
