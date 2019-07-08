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
		if sensor.Temp < config.Thresholds.Temp[0] || sensor.Temp > config.Thresholds.Temp[1] {
			sendSticker("CAADAgAD9AEAAnELQgX2FdOwiYZwbgI")
			sendMsg(fmt.Sprintf("Temp - %.2f", sensor.Temp))

		}
		if sensor.Humidity < config.Thresholds.Humidity[0] || sensor.Humidity > config.Thresholds.Humidity[1] {
			sendSticker("CAADAgADNQADuhxDEv_cQsYUoBwZAg")
			sendMsg(fmt.Sprintf("Humidity - %.2f", sensor.Humidity))
		}
		if sensor.CO2 < config.Thresholds.CO2[0] || sensor.CO2 > config.Thresholds.CO2[1] {
			//sendSticker("CAADAgAD6wEAAnELQgXabYB2DWrLYAI")
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
