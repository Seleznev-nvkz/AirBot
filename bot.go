package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"path/filepath"
)

const helpText = `
/subscribe
/unsubscribe
/check
/graph
/help`

func handleCommand(message *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")

	switch message.Command() {
	case "subscribe":
		config.appendSubscriber(message.Chat.ID)
		log.Printf("Subsribed %v", message.Chat.ID)
		msg.Text = "You successful subscribed."
	case "help":
		msg.Text = helpText
	case "unsubscribe":
		config.removeSubscriber(message.Chat.ID)
		log.Printf("Unubsribed %v", message.Chat.ID)
		msg.Text = "You successful unsubscribed."
	case "check":
		msg.Text = sensor.FreshString()
	case "graph":
		buildGraph()
		absPath, _ := filepath.Abs("./plot.png")
		sendPhoto(absPath)
		return
	default:
		msg.Text = "I don't know that command"
	}

	if _, err := bot.Send(msg); err != nil {
		log.Panic(err)
	}
}

func sendSticker(fileId string) {
	for _, id := range config.Subscribers {
		log.Println(id)
		if _, err := bot.Send(tgbotapi.NewStickerShare(id, fileId)); err != nil {
			log.Panic(err)
		}
	}
}

func sendMsg(msg string) {
	for _, id := range config.Subscribers {
		log.Println(id)
		if _, err := bot.Send(tgbotapi.NewMessage(id, msg)); err != nil {
			log.Panic(err)
		}
	}
}

func sendPhoto(fileId string) {
	for _, id := range config.Subscribers {
		log.Println(id)
		if _, err := bot.Send(tgbotapi.NewPhotoShare(id, fileId)); err != nil {
			log.Panic(err)
		}
	}
}

func listen() {
	log.Printf("%s running", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, _ := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if update.Message.IsCommand() {
			handleCommand(update.Message)
			continue
		}

		if update.Message.Entities != nil {
			for _, i := range *update.Message.Entities {
				if i.Type == "mention" { // type tgbotapi.MessageEntity has no field or method IsMention
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, sensor.FreshString()))
				}
			}
		}
	}
}
