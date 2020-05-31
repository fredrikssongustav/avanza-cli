package utils

import (
	"fmt"
	"github.com/guptarohit/asciigraph"
)

func Chart(timeSeries [][]float64) {
	var chartY []float64

	for _, t := range timeSeries {
		chartY = append(chartY,t[1])
	}

	graph := asciigraph.Plot(chartY)

	fmt.Println(graph)
}
