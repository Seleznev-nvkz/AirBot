package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"time"
)

var configPath string

type Config struct {
	Subscribers []int64       `yaml:"subscribers"`
	Interval    time.Duration `yaml:"interval"`
	Url         string        `yaml:"url"`
	DBPath      string        `yaml:"dbPath"`
	Thresholds  Thresholds    `yaml:"thresholds"`
}

type Thresholds struct {
	Temp     [2]float64 `yaml:"temp"`
	Humidity [2]float64 `yaml:"humidity"`
	CO2      [2]int     `yaml:"co2"`
}

func NewConfig() *Config {
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
	}

	bytes, err := ioutil.ReadFile(configPath)
	if err == nil {
		err = yaml.Unmarshal(bytes, c)
		if err != nil {
			log.Panic(err)
		}
	}

	url, ok := os.LookupEnv("URL")
	if ok {
		c.Url = url
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

	return ioutil.WriteFile(configPath, bytes, 0644)
}

//func (c *Config) init(configPath string) {
//	var err error
//	db, err = storm.Open(configPath)
//	defer db.Close()
//	if err != nil {
//		log.Panic(err)
//	}
//
//	db.Init(&Config{})
//	err = db.One("ID", 1, c)
//	log.Println(err)
//	log.Println(c.subscribers)
//
//	var ok bool
//	c.subscribers = append(c.subscribers, int64(123))
//	log.Println(c.subscribers)
//	err = db.Save(c)
//	log.Println(err)
//}

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
