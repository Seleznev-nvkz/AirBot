package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecord(t *testing.T) {
	configPath := "/tmp/config.yaml"
	config = NewConfig(configPath)
	config.DBPath = "/tmp/test_parse.db"
	db = newDB(config.DBPath)

	defer os.Remove(config.DBPath)
	defer os.Remove(configPath)

	testSensor := &Sensor{
		Temp:     12.0,
		Humidity: 13.0,
		CO2:      15,
	}

	_ = SaveRecord(testSensor)
	records := GetRecentStats()
	assert.Equal(t, len(records), 1)
	assert.Equal(t, records[0].Sensor.String(), testSensor.String())
}
