package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

func ArrrPrice(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println("mux vars:", vars)
	// fmt.Println("symbols:", vars["symbols"])
	// fmt.Printf("Market: %v\nPair: ARRR/%v\n", strings.ToLower(vars["market"]), strings.ToUpper(vars["pair"]))
	var arrrPrice []byte
	var _arrrPrice pirateRates
	switch strings.ToLower(vars["market"]) {
	case "tradeogre":
		_arrrPrice = ARRR_TO_RATES
	case "kucoin":
		switch strings.ToUpper(vars["pair"]) {
		case "ARRR-BTC":
			_arrrPrice = ARRR_KC_RATES_BTC
		case "ARRR-USDT":
			_arrrPrice = ARRR_KC_RATES_USDT
		}
	}
	arrrPrice, _ = json.Marshal(_arrrPrice)

	if vars["symbols"] != "" {
		symbols := strings.Split(vars["symbols"], ",")
		var result interface{}
		json.Unmarshal([]byte(arrrPrice), &result)
		var rates []map[string]float64
		var qr queriedRates
		for _, symbol := range symbols {
			// fmt.Println(symbol)
			for i, v := range result.(map[string]interface{})["rates"].(map[string]interface{}) {
				if strings.Compare(i, symbol) == 0 {
					// fmt.Println("i -", i)
					// fmt.Println("v -", v)
					rates = append(rates, map[string]float64{i: v.(float64)})
				}
			}
		}
		// fmt.Println(rates)
		qr.Timestamp = _arrrPrice.Timestamp
		qr.Rates = rates
		qr.Base = _arrrPrice.Base
		qr.Market = _arrrPrice.Market
		// fmt.Println(qr)

		queryRes, _ := json.Marshal(qr)
		// fmt.Println(string(queryRes))
		w.Header().Set("Content-Type", "application/json")
		w.Write(queryRes)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(arrrPrice)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<p>For API documentation visit <a href="https://github.com/Meshbits/Pirate-Metrics">Pirate Metrics Repository on Github</a>.</p>`)
}

func main() {
	fixerAPIToken := flag.String("fixerAPIToken", "", "Get the API token from https://fixer.io. API Token from a free sign-up is sufficient to use with this application.")
	flag.Parse()
	fmt.Println("fixerAPIToken:", *fixerAPIToken)
	if *fixerAPIToken == "" {
		return
	}

	go fixer(*fixerAPIToken)
	go CGeckoBTCAPI()
	go ArrrToAPI()
	go ArrrBtcKcAPI()
	go ArrrUsdtKcAPI()

	go displayRates()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/v1/prices/{market}/{pair}", ArrrPrice).Methods("GET").Queries("symbols", "{symbols}")
	r.HandleFunc("/v1/prices/{market}/{pair}", ArrrPrice).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

	// fmt.Scanln()

}

func displayRates() {
	for {
		// fmt.Printf("FIXER_RATES:\n %v\n", FIXER_RATES)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Printf("BTC USD: %v\n", BTC_RATES.Rates.USD)
		fmt.Printf("BTC NZD: %v\n", BTC_RATES.Rates.NZD)
		fmt.Printf("TO-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_TO_RATES.Rates.BTC)
		fmt.Printf("TO-ARRR/BTC: ARRR (USD): %v\n", ARRR_TO_RATES.Rates.USD)
		fmt.Printf("KC-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_KC_RATES_BTC.Rates.BTC)
		fmt.Printf("KC-ARRR/BTC: ARRR (USD): %v\n", ARRR_KC_RATES_BTC.Rates.USD)
		fmt.Printf("KC-ARRR/USDT: ARRR (BTC): %.8f\n", ARRR_KC_RATES_USDT.Rates.BTC)
		fmt.Printf("KC-ARRR/USDT: ARRR (USD): %v\n", ARRR_KC_RATES_USDT.Rates.USD)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		sleepSeconds := 900
		fmt.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
