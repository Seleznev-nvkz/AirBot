package main

import (
	"fmt"
	"log"
	"time"
)

const (
	TEMPERATURE = iota
	HUMIDITY
	CO2
)

func runChecker() {
	ticker := time.NewTicker(config.Interval)

	for range ticker.C {
		sensor.Update()
		switch sensor.validate() {
		case TEMPERATURE:
			sendSticker("CAADAgAD9AEAAnELQgX2FdOwiYZwbgI")
			sendMsg(fmt.Sprintf("Temp - %.2f", sensor.Temp))
		case HUMIDITY:
			sendSticker("CAADAgADNQADuhxDEv_cQsYUoBwZAg")
			sendMsg(fmt.Sprintf("Humidity - %.2f", sensor.Humidity))
		case CO2:
			sendSticker("CAADAgAD9AEAAnELQgX2FdOwiYZwbgI")
			sendMsg(fmt.Sprintf("CO2 - %v", sensor.CO2))
		}

		now := time.Now()
		if now.Minute() == 30 && now.Hour() == 19 {
			sendSticker("CAADAgADEQADrQWwDJL18PQEXEkiAg")
		}

		err := SaveRecord(sensor)
		if err != nil {
			log.Println(err)
		}
	}
}
