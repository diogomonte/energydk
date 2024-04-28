package main

import (
	"encoding/json"
	"github.com/a-h/templ"
	"github.com/montediogo/energydk/src/energy"
	"github.com/montediogo/energydk/src/ui"
	"log"
	"net/http"
)

type ErrorMessage struct {
	message string
}

func listEnergy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	prices, err := energy.ListEnergyPrices()
	if err != nil {
		errorMessage := ErrorMessage{message: err.Error()}
		_ = json.NewEncoder(w).Encode(errorMessage)
	}
	_ = json.NewEncoder(w).Encode(prices)
}

func main() {
	indexComponent := ui.Index()
	http.Handle("/", templ.Handler(indexComponent))
	//http.HandleFunc("/energy", listEnergy)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("server error", err)
	}
}
