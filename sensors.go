package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type Sensor struct {
	temp     float64
	humidity float64
	co2      int
}

func NewSensor() *Sensor {
	return &Sensor{
		temp:     0.0,
		humidity: 0.0,
		co2:      0,
	}
}

func (s *Sensor) String() string {
	return fmt.Sprintf("temp - %v\nhumidity - %v\nCO2 - %v", s.temp, s.humidity, s.co2)
}

func (s *Sensor) FreshString() string {
	err := s.Update()
	if err != nil {
		return "no way"
	}
	return s.String()
}

func (s *Sensor) Update() error {
	resp, err := http.Get(config.Url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return err
	}
	p := doc.Find("p").Nodes

	s.temp, _ = strconv.ParseFloat(strings.TrimRight(
		strings.TrimLeft(p[0].FirstChild.Data, `temp - `), ` C`), 64)
	s.humidity, _ = strconv.ParseFloat(strings.TrimRight(
		strings.TrimLeft(p[1].FirstChild.Data, `humidity - `), ` %`), 64)
	s.co2, _ = strconv.Atoi(strings.TrimRight(strings.TrimLeft(p[2].FirstChild.Data, `CO2 - `), ` ppm`))

	return nil
}
