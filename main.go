package main

import (
	"flag"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"os"
)

var sensor *Sensor
var bot *tgbotapi.BotAPI
var config *Config
var db *DB

func init() {
	var configPath string

	if flag.Lookup("test.v") == nil {
		token, ok := os.LookupEnv("TOKEN")
		if !ok {
			panic("Set TOKEN")
		}

		configPath, ok = os.LookupEnv("CONFIG_PATH")
		if !ok {
			configPath = "config.yaml"
		}

		bot, _ = tgbotapi.NewBotAPI(token)
		bot.Debug = true
	} else {
		configPath = "/tmp/config.yaml"
	}

	config = NewConfig(configPath)
	db = newDB(config.DBPath)
	sensor = NewSensor()
}

func main() {
	go runChecker()
	listen()
}
