package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

var sensor *Sensor
var bot *tgbotapi.BotAPI
var config *Config
var db *DB

func init() {
	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		panic("Set TOKEN")
	}

	configPath, ok = os.LookupEnv("CONFIG_PATH")
	if !ok {
		configPath = "config.yaml"
	}

	config = NewConfig()
	db = newDB(config.DBPath)
	sensor = NewSensor()
	bot, _ = tgbotapi.NewBotAPI(token)
	bot.Debug = true
}

func main() {
	go runChecker()
	listen()
}
