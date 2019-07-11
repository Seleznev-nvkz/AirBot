package main

import (
	"github.com/magiconair/properties/assert"
	"os"
	"testing"
)

func TestSensorValidate(t *testing.T) {
	configPath := "/tmp/config.yaml"
	config = NewConfig(configPath)
	sensor = NewSensor()
	defer os.Remove(configPath)

	sensor.Temp = config.Thresholds.Temp[1] + 1
	assert.Equal(t, TEMPERATURE, sensor.validate())
	sensor.Temp = config.Thresholds.Temp[1]

	sensor.Humidity = config.Thresholds.Humidity[1] + 1
	assert.Equal(t, HUMIDITY, sensor.validate())
	sensor.Humidity = config.Thresholds.Humidity[1]

	sensor.CO2 = config.Thresholds.CO2[1] + 1
	assert.Equal(t, CO2, sensor.validate())
}
