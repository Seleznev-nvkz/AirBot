package main

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThresholds(t *testing.T) {
	configPath := "/tmp/config.yaml"
	config = NewConfig(configPath)
	config.DBPath = "/tmp/test_parse.db"
	sensor = NewSensor()
	defer os.Remove(configPath)

	refreshThresholds = time.Millisecond * 50 // for test auto-refresh

	// default
	sensor.Temp = config.Thresholds.Temp[1]
	sensor.Humidity = config.Thresholds.Humidity[1]
	sensor.CO2 = config.Thresholds.CO2[1]
	sensor.validate()

	assert.Equal(t, config.currentThresholds.CO2[1], config.Thresholds.CO2[1])
	assert.Equal(t, config.currentThresholds.Humidity[1], config.Thresholds.Humidity[1])
	assert.Equal(t, config.currentThresholds.Temp[1], config.Thresholds.Temp[1])

	// up
	sensor.Temp++
	sensor.Humidity++
	sensor.CO2++

	config.tempUp()
	config.humidityUp()
	config.co2Up()

	assert.Equal(t, config.currentThresholds.Temp[1], config.Thresholds.Temp[1]+config.Steps.Temp)
	assert.Equal(t, config.currentThresholds.Humidity[1], config.Thresholds.Humidity[1]+config.Steps.Humidity)
	assert.Equal(t, config.currentThresholds.CO2[1], config.Thresholds.CO2[1]+config.Steps.CO2)

	// auto-reset
	time.Sleep(refreshThresholds * 2)
	assert.Equal(t, config.currentThresholds.CO2[1], config.Thresholds.CO2[1])
	assert.Equal(t, config.currentThresholds.Humidity[1], config.Thresholds.Humidity[1])
	assert.Equal(t, config.currentThresholds.Temp[1], config.Thresholds.Temp[1])
}
