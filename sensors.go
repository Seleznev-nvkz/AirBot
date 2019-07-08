package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type Sensor struct {
	Temp     float64
	Humidity float64
	CO2      int
}

func NewSensor() *Sensor {
	return &Sensor{
		Temp:     0.0,
		Humidity: 0.0,
		CO2:      0,
	}
}

func (s *Sensor) String() string {
	return fmt.Sprintf("temp - %v\nhumidity - %v\nCO2 - %v", s.Temp, s.Humidity, s.CO2)
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

	s.Temp, _ = strconv.ParseFloat(strings.TrimRight(
		strings.TrimLeft(p[0].FirstChild.Data, `temp - `), ` C`), 64)
	s.Humidity, _ = strconv.ParseFloat(strings.TrimRight(
		strings.TrimLeft(p[1].FirstChild.Data, `humidity - `), ` %`), 64)
	s.CO2, _ = strconv.Atoi(strings.TrimRight(strings.TrimLeft(p[2].FirstChild.Data, `CO2 - `), ` ppm`))

	return nil
}
