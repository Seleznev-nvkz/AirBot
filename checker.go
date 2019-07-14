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
		err := sensor.Update()
		if err != nil {
			log.Panic(err)
		}
		now := time.Now()
		day := now.Weekday().String()

		if config.WorkFinish > now.Hour() && now.Hour() > config.WorkStart {
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
		}

		if day != "Sunday" && day != "Saturday" {
			if now.Minute() == 30 && now.Hour() == 19 {
				sendSticker("CAADAgADEQADrQWwDJL18PQEXEkiAg")
			}
		}

		err = SaveRecord(sensor)
		if err != nil {
			log.Println(err)
		}
	}
}
