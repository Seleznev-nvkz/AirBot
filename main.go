package main

import (
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

var sensor *Sensor
var bot *tgbotapi.BotAPI
var config *Config
var db *DB

func main() {
	var configPath string

	token, ok := os.LookupEnv("TOKEN")
	if !ok {
		panic("Set TOKEN")
	}
	bot, _ = tgbotapi.NewBotAPI(token)
	bot.Debug = true

	configPath, ok = os.LookupEnv("CONFIG_PATH")
	if !ok {
		configPath = "config.yaml"
	}
	config = NewConfig(configPath)
	db = newDB(config.DBPath)
	sensor = NewSensor()

	go runChecker()
	listen()
}
