package main

import (
	"fmt"
	"time"
)

type StatsRecord struct {
	*Sensor
	ID        int       `storm:"id,increment"`
	Timestamp time.Time `storm:"index"`
}

func SaveRecord(sensor *Sensor) error {
	record := StatsRecord{
		Sensor:    sensor,
		Timestamp: time.Now(),
	}
	return db.Save(&record)
}

func (s *StatsRecord) String() string {
	return fmt.Sprintf("ID - %v\nTimestamp - %v\ntemp - %v\nhumidity - %v\nCO2 - %v",
		s.ID, s.Timestamp, s.temp, s.humidity, s.co2)
}
