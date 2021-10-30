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

var MARKETS_AVAILABLE = map[string][]string{}

var BTC_PRICE_SOURCE string

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<p>For API documentation visit <a href="https://github.com/Meshbits/Pirate-Metrics">Pirate Metrics Repository on Github</a>.</p>`)
}

func PriceAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println("mux vars:", vars)
	// fmt.Println("symbols:", vars["symbols"])
	fmt.Printf("Market: %v\nPair: %v\n", strings.ToLower(vars["market"]), strings.ToUpper(vars["pair"]))
	var convRates []byte
	var _convRates ConversionRates
	switch strings.ToLower(vars["market"]) {
	case "fixer":
		switch strings.ToUpper(vars["pair"]) {
		case "USD":
			_convRates = FIXER_RATES
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "tradeogre":
		switch strings.ToUpper(vars["pair"]) {
		case "ARRR-BTC":
			_convRates = ARRR_TO_RATES
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "kucoin":
		switch strings.ToUpper(vars["pair"]) {
		case "ARRR-BTC":
			_convRates = ARRR_KC_RATES_BTC
		case "ARRR-USDT":
			_convRates = ARRR_KC_RATES_USDT
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "binance":
		fmt.Println("Binance selected", vars["market"])
		if strings.ToLower(BTC_PRICE_SOURCE) != strings.ToLower(vars["market"]) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "This market not enabled. Please try another market"}`)
			return
		}
		switch strings.ToUpper(vars["pair"]) {
		case "BTC-USDT":
			_convRates = BTC_BINANCE_RATES_USDT
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "coingecko":
		fmt.Println("CoinGecko selected", vars["market"])
		if strings.ToLower(BTC_PRICE_SOURCE) != strings.ToLower(vars["market"]) {
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "This market not enabled. Please try another market"}`)
			return
		}
		switch strings.ToUpper(vars["pair"]) {
		case "BTC-USD":
			_convRates = BTC_COINGECKO_RATES
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	default:
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, `{"success": false,"error": "Market not found. Please try another market"}`)
		return
	}
	convRates, _ = json.Marshal(_convRates)

	if vars["symbols"] != "" {
		symbols := strings.Split(vars["symbols"], ",")
		var result interface{}
		json.Unmarshal([]byte(convRates), &result)
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
		qr.Timestamp = _convRates.Timestamp
		qr.Rates = rates
		qr.Base = _convRates.Base
		qr.Market = _convRates.Market
		// fmt.Println(qr)

		queryRes, _ := json.Marshal(qr)
		// fmt.Println(string(queryRes))
		w.Header().Set("Content-Type", "application/json")
		w.Write(queryRes)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(convRates)
}

func Markets(w http.ResponseWriter, r *http.Request) {
	markets, _ := json.Marshal(MARKETS_AVAILABLE)
	w.Header().Set("Content-Type", "application/json")
	w.Write(markets)
}

func main() {
	fixerAPIToken := flag.String("fixerAPIToken", "", "Get the API token from https://fixer.io. API Token from a free sign-up is sufficient to use with this application.")
	btcPriceSource := flag.String("btcPriceSource", "Binance", "You can specify any single string for this option: Binance, CoinGecko")
	flag.Parse()
	fmt.Println("fixerAPIToken:", *fixerAPIToken)
	fmt.Println("btcPriceSource:", *btcPriceSource)
	if *fixerAPIToken == "" {
		return
	}

	go fixer(*fixerAPIToken)

	switch strings.ToLower(*btcPriceSource) {
	case "binance":
		go BtcUsdtBinanceAPI()
		BTC_PRICE_SOURCE = "Binance"
	case "coingecko":
		go BtcUsdCoinGeckoAPI()
		BTC_PRICE_SOURCE = "CoinGecko"
	}

	go ArrrToAPI()
	go ArrrBtcKcAPI()
	go ArrrUsdtKcAPI()

	go displayRates()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/v1/prices/{market}/{pair}", PriceAPI).Methods("GET").Queries("symbols", "{symbols}")
	r.HandleFunc("/v1/prices/{market}/{pair}", PriceAPI).Methods("GET")
	r.HandleFunc("/v1/prices/markets", Markets).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

	// fmt.Scanln()

}

func displayRates() {
	for {
		// fmt.Printf("FIXER_RATES:\n %v\n", FIXER_RATES)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("BTC Price Source is:", BTC_PRICE_SOURCE)
		fmt.Printf("BTC USD: %v\n", BTC_RATES.Rates.USD)
		fmt.Printf("BTC NZD: %v\n", BTC_RATES.Rates.NZD)
		fmt.Printf("TO-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_TO_RATES.Rates.BTC)
		fmt.Printf("TO-ARRR/BTC: ARRR (USD): %v\n", ARRR_TO_RATES.Rates.USD)
		fmt.Printf("KC-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_KC_RATES_BTC.Rates.BTC)
		fmt.Printf("KC-ARRR/BTC: ARRR (USD): %v\n", ARRR_KC_RATES_BTC.Rates.USD)
		fmt.Printf("KC-ARRR/USDT: ARRR (BTC): %.8f\n", ARRR_KC_RATES_USDT.Rates.BTC)
		fmt.Printf("KC-ARRR/USDT: ARRR (USD): %v\n", ARRR_KC_RATES_USDT.Rates.USD)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		sleepSeconds := 10
		fmt.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
