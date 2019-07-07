package main

import (
	"github.com/Arafatk/glot"
	"time"
)

func buildGraph() {

	//data := GetRecentStats()
	timeRange := getRange(5)

	dimensions := 2
	plot, _ := glot.NewPlot(dimensions, false, false)

	_ = plot.SetTitle(time.Now().Format(time.Stamp))

	pointGroupName := "CO2 x10"
	style := "lines"
	points := [][]float64{timeRange, {555 / 10, 941 / 10, 949 / 10, 1032 / 10, 1340 / 10}}
	_ = plot.AddPointGroup(pointGroupName, style, points)

	pointGroupName = "Temp"
	points = [][]float64{timeRange, {26, 26, 23, 21, 19}}
	_ = plot.AddPointGroup(pointGroupName, style, points)

	_ = plot.SavePlot("plot.png")

}

func getRange(len int) []float64 {
	var timeRange []float64
	for i := -1.0; i >= float64(-len); i-- {
		timeRange = append(timeRange, i)
	}
	return timeRange
}
