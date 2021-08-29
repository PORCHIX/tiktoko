package main

import (
	"os"

	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"github.com/joho/godotenv"
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

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
