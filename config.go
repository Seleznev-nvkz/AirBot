package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Subscribers []int64       `yaml:"subscribers"`
	Interval    time.Duration `yaml:"interval"`
	Url         string        `yaml:"url"`
	DBPath      string        `yaml:"dbPath"`
	PlotPath    string        `yaml:"plotPath"`
	Thresholds  Thresholds    `yaml:"thresholds"`
	Steps       Steps         `yaml:"steps"`

	currentThresholds Thresholds
	path              string
}

type Thresholds struct {
	Temp     [2]float64 `yaml:"temp"`
	Humidity [2]float64 `yaml:"humidity"`
	CO2      [2]int     `yaml:"co2"`
}

type Steps struct {
	Temp     float64 `yaml:"temp"`
	Humidity float64 `yaml:"humidity"`
	CO2      int     `yaml:"co2"`
}

func NewConfig(path string) *Config {
	// create default config
	c := &Config{
		Subscribers: []int64{},
		Url:         "http://192.168.88.192",
		Interval:    time.Minute,
		DBPath:      "db",
		Thresholds: Thresholds{
			Temp:     [2]float64{20.0, 34.0},
			Humidity: [2]float64{0.0, 80.0},
			CO2:      [2]int{0, 1200},
		},
		Steps: Steps{
			Temp:     1.0,
			Humidity: 2.0,
			CO2:      100,
		},
		PlotPath: "/data/plot.png",

		currentThresholds: Thresholds{
			Temp:     [2]float64{20.0, 34.0},
			Humidity: [2]float64{0.0, 80.0},
			CO2:      [2]int{0, 1200},
		},
		path: path,
	}

	// update from file
	bytes, err := ioutil.ReadFile(c.path)
	if err == nil {
		err = yaml.Unmarshal(bytes, c)
		if err != nil {
			log.Panic(err)
		}
	}

	group, ok := os.LookupEnv("GROUP_ID")
	if ok {
		subscriber, err := strconv.ParseInt(group, 10, 64)
		if err == nil {
			c.appendSubscriber(subscriber)
		}
	}

	c.save()
	log.Println("Config loaded")
	return c
}

func (c *Config) save() error {
	bytes, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.path, bytes, 0644)
}

func (c *Config) tempUp() {
	c.currentThresholds.Temp[1] += c.Steps.Temp
	c.resetThresholds()
}

func (c *Config) humidityUp() {
	c.currentThresholds.Humidity[1] += c.Steps.Humidity
	c.resetThresholds()
}

func (c *Config) co2Up() {
	c.currentThresholds.CO2[1] += c.Steps.CO2
	c.resetThresholds()
}

var refreshThresholds = time.Hour * 2
var refreshReady = true

func (c *Config) resetThresholds() {
	if refreshReady {
		refreshReady = false
		go func(timeout time.Duration) {
			time.Sleep(timeout)
			c.currentThresholds = c.Thresholds
		}(refreshThresholds)
	}
}

func (c *Config) appendSubscriber(newSubscriber int64) {
	exists := false
	for _, subscriber := range c.Subscribers {
		if subscriber == newSubscriber {
			exists = true
			break
		}
	}

	if !exists {
		c.Subscribers = append(c.Subscribers, newSubscriber)
		c.save()
	}
}

func (c *Config) removeSubscriber(subscriber int64) {
	for i, sub := range c.Subscribers {
		if sub == subscriber {
			c.Subscribers[i] = c.Subscribers[len(c.Subscribers)-1]
			c.Subscribers = c.Subscribers[:len(c.Subscribers)-1]
			break
		}
	}
	c.save()
}
