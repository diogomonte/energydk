package chart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func DrawBarChart(tile string, xAxis []string, values []interface{}) *charts.Bar {
	bar := charts.NewBar()
	bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title: tile,
	}))
	bar.SetXAxis(xAxis)
	bar.AddSeries("energy", generateBarItems(values))
	return bar
}

func generateBarItems(values []interface{}) []opts.BarData {
	items := make([]opts.BarData, 0)
	for _, value := range values {
		items = append(items, opts.BarData{Value: value})
	}
	return items
}
