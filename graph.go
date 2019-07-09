package main

import (
	"github.com/Arafatk/glot"
	"time"
)

func buildGraph() {
	var (
		temp []float64
		co2  []float64
	)

	for _, v := range GetRecentStats() {
		temp = append(temp, v.Sensor.Temp)
		co2 = append(co2, float64(v.Sensor.CO2/10))
	}

	timeRange := getRange(len(temp))
	dimensions := 2
	plot, _ := glot.NewPlot(dimensions, false, false)

	_ = plot.SetTitle(time.Now().Format(time.Stamp))

	pointGroupName := "CO2 x10"
	style := "lines"
	points := [][]float64{timeRange, co2}
	_ = plot.AddPointGroup(pointGroupName, style, points)

	pointGroupName = "Temp"
	points = [][]float64{timeRange, temp}
	_ = plot.AddPointGroup(pointGroupName, style, points)
	_ = plot.SavePlot(config.PlotPath)

	// w8 for replot
	time.Sleep(time.Second * 2)
}

func getRange(len int) []float64 {
	var timeRange []float64
	for i := -1.0; i >= float64(-len); i-- {
		timeRange = append(timeRange, i)
	}
	return timeRange
}
