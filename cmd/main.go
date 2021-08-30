package main

import (
	"os"
	"tik-tok-video-downloader/pkg/handler"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("Error loading .env file: %s", err.Error())
	}

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TOKEN"))
	logrus.Printf(os.Getenv("TOKEN"))
	if err != nil {
		logrus.Fatalf(err.Error())
	}

	bot.Debug = true

	logrus.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	h := handler.NewHandler(bot)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		h.HandleMessage([]string{"/start", "/about"}, handler.MessageStart(update.Message))
		h.HandleMessage([]string{""}, handler.MessageText(update.Message))
	}
}
