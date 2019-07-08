package main

import (
	"fmt"
	"time"
)

type StatsRecord struct {
	Sensor
	ID        int       `storm:"id,increment"`
	Timestamp time.Time `storm:"index"`
}

func SaveRecord(sensor *Sensor) error {
	record := StatsRecord{
		Sensor:    *sensor,
		Timestamp: time.Now(),
	}
	return db.Save(&record)
}

func (s *StatsRecord) String() string {
	return fmt.Sprintf("ID - %v\nTimestamp - %v\ntemp - %v\nhumidity - %v\nCO2 - %v",
		s.ID, s.Timestamp, s.Temp, s.Humidity, s.CO2)
}

func GetStatsByLastDay() (res []StatsRecord) {
	now := time.Now()
	db.Range("Timestamp", now.AddDate(0, 0, -1), now, &res)
	return
}
