package main

import (
	"fmt"
	"log"
	"time"
)

func runChecker() {
	ticker := time.NewTicker(config.Interval)

	for range ticker.C {
		sensor.Update()
		if sensor.temp < config.Thresholds.Temp[0] || sensor.temp > config.Thresholds.Temp[1] {
			sendSticker("CAADAgAD9AEAAnELQgX2FdOwiYZwbgI")
			sendMsg(fmt.Sprintf("Temp - %.2f", sensor.temp))

		}
		if sensor.humidity < config.Thresholds.Humidity[0] || sensor.humidity > config.Thresholds.Humidity[1] {
			sendSticker("CAADAgADNQADuhxDEv_cQsYUoBwZAg")
			sendMsg(fmt.Sprintf("Humidity - %.2f", sensor.humidity))
		}
		if sensor.co2 < config.Thresholds.CO2[0] || sensor.co2 > config.Thresholds.CO2[1] {
			//sendSticker("CAADAgAD6wEAAnELQgXabYB2DWrLYAI")
			sendSticker("CAADAgAD9AEAAnELQgX2FdOwiYZwbgI")
			sendMsg(fmt.Sprintf("CO2 - %v", sensor.co2))
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
