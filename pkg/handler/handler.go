package handler

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"tik-tok-video-downloader/tiktok"
)

type Handler struct {
	bot *tgbotapi.BotAPI
}

type response struct {
	chattable tgbotapi.Chattable
	message   *tgbotapi.Message
	filename string
}

func NewHandler(bot *tgbotapi.BotAPI) *Handler {
	return &Handler{bot: bot}
}

func (h *Handler) HandleMessage(commands []string, r response) {
	if r.message == nil && r.chattable == nil {
		return
	}
	for _, v := range commands {
		if v == r.message.Text {
			h.bot.Send(r.chattable)
		}
	}
	if len(commands) == 1 && commands[0] == "" {
		h.bot.Send(r.chattable)
		if err := os.Remove(r.filename); err != nil{
			logrus.Printf("Unable to remove %s", r.filename)
		}
	}
	return
}

func MessageStart(message *tgbotapi.Message) response {
	res := tgbotapi.NewMessage(message.Chat.ID, "Hi! I am Tik Tok video downloader. Every time when you will send me a link to a tiktok video I will download mp4 version of it. Thanks for using me, peace! ")
	res.ReplyToMessageID = message.MessageID
	return response{res, message, ""}
}

func MessageText(message *tgbotapi.Message) response {
	for _, m := range strings.Split(message.Text, " ") {
		logrus.Printf(m)
		filename, err := tiktok.DownloadTikTokVideo(m)
		if err != nil {
			logrus.Printf(err.Error())
			continue
		}
		path := "./" + filename
		//path := filename

		videoCfg := tgbotapi.NewVideoUpload(message.Chat.ID, path)
		videoCfg.Caption = "Downloaded via @ttto_mov_bot"
		videoCfg.ReplyToMessageID = message.MessageID
		return response{videoCfg, message, filename}
	}
	return response{nil, nil, ""}
}
