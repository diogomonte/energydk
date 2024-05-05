package main

import (
	"github.com/montediogo/energydk/src/chart"
	"github.com/montediogo/energydk/src/energy"
	"log"
	"net/http"
)

func httpserver(w http.ResponseWriter, _ *http.Request) {
	prices, _ := energy.ListEnergyPrices()

	xAxis := make([]string, 0, len(prices.AllPrices))
	values := make([]interface{}, 0, len(prices.AllPrices))

	for _, price := range prices.AllPrices {
		xAxis = append(xAxis, price.Time)
		values = append(values, price.PriceDkk)
	}
	barChart := chart.DrawBarChart("Energy prices", xAxis, values)
	barChart.Render(w)
}

func main() {
	http.HandleFunc("/", httpserver)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("server error", err)
	}
}
