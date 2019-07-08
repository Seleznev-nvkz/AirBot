package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSaveRecord(t *testing.T) {
	db = newDB("/tmp/test_parse.db")
	defer os.Remove("/tmp/test_parse.db")

	testSensor := &Sensor{
		Temp:     12.0,
		Humidity: 13.0,
		CO2:      15,
	}

	SaveRecord(testSensor)
	records := GetStatsByLastDay()
	assert.Equal(t, len(records), 1)
	assert.Equal(t, records[0].Sensor.String(), testSensor.String())
}
