package main

import (
	"bytes"
	"github.com/wcharczuk/go-chart"
	"log"
	"os"
	"time"
)

func buildGraph(mode string) {
	recentStats, err := GetRecentStats()
	if err != nil {
		log.Panic(err)
	}
	var (
		temp      []float64
		co2       []float64
		humidity  []float64
		timestamp []time.Time
		zeros     []float64
	)

	for _, v := range recentStats {
		t := v.Sensor.Temp
		temp = append(temp, t)
		zeros = append(zeros, t/2)
		if mode == "co2" {
			co2 = append(co2, float64(v.Sensor.CO2))
		} else {
			co2 = append(co2, float64(v.Sensor.CO2/10))
		}
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

	zerosTs := chart.TimeSeries{
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorTransparent,
		},
		XValues: timestamp,
		YValues: zeros,
	}
	var series []chart.Series

	switch mode {
	case "temp":
		series = []chart.Series{
			temperatureTs,
		}
	case "co2":
		series = []chart.Series{
			co2Ts,
		}
	case "hum":
		series = []chart.Series{
			humidityTs,
		}
	default:
		series = []chart.Series{
			temperatureTs,
			co2Ts,
			humidityTs,
		}
	}
	series = append(series, zerosTs)

	graph := chart.Chart{
		Background: chart.Style{
			Padding: chart.Box{
				Top: 30,
			},
		},
		XAxis: chart.XAxis{
			NameStyle:      chart.StyleShow(),
			Style:          chart.StyleShow(),
			ValueFormatter: chart.TimeValueFormatterWithFormat("15:04"),
		},
		YAxis: chart.YAxis{
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		Series: series,
	}

	graph.Elements = []chart.Renderable{
		chart.Legend(&graph,
			chart.Style{FontSize: 9.0,
				FillColor:   chart.ColorLightGray,
				StrokeColor: chart.ColorTransparent,
			}),
	}

	buffer := bytes.NewBuffer([]byte{})
	err = graph.Render(chart.PNG, buffer)
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
