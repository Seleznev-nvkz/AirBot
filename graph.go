package main

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"log"
	"os"
	"time"
)

func buildGraph() {
	var (
		temp      []float64
		co2       []float64
		humidity  []float64
		timestamp []time.Time
	)

	for _, v := range GetRecentStats() {
		temp = append(temp, v.Sensor.Temp)
		co2 = append(co2, float64(v.Sensor.CO2/10))
		humidity = append(humidity, v.Sensor.Humidity)
		timestamp = append(timestamp, v.Timestamp)
	}

	temperatureTs := chart.TimeSeries{
		Style: chart.Style{
			Show:      true,
			FillColor: chart.ColorBlue.WithAlpha(15),
		},
		Name:    "Temperature",
		XValues: timestamp,
		YValues: temp,
	}

	co2Ts := chart.TimeSeries{
		Style: chart.Style{
			Show:      true,
			FillColor: chart.ColorGreen.WithAlpha(15),
		},
		Name:    "CO2 x10",
		XValues: timestamp,
		YValues: co2,
	}

	humidityTs := chart.TimeSeries{
		Style: chart.Style{
			Show:      true,
			FillColor: chart.ColorBlack.WithAlpha(15),
		},
		Name:    "Humidity",
		XValues: timestamp,
		YValues: humidity,
	}

	graph := chart.Chart{
		XAxis: chart.XAxis{
			NameStyle:      chart.StyleShow(),
			Style:          chart.StyleShow(),
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis: chart.YAxis{
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: []chart.Series{
			temperatureTs,
			co2Ts,
			humidityTs,
		},
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph,
			chart.Style{FontSize: 9.0,
				FillColor:   chart.ColorLightGray,
				StrokeColor: chart.ColorTransparent,
			}),
	}

	buffer := bytes.NewBuffer([]byte{})
	err := graph.Render(chart.PNG, buffer)
	if err != nil {
		log.Fatal(err)
	}

	fo, err := os.Create(config.PlotPath)
	if err != nil {
		panic(err)
	}

	if _, err := fo.Write(buffer.Bytes()); err != nil {
		panic(err)
	}
}
