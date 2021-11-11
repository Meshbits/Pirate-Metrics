package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

func indexPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<p>For API documentation visit <a href="https://github.com/Meshbits/Pirate-Metrics">Pirate Metrics Repository on Github</a>.</p>`)
}

func PriceAPI(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// fmt.Println("mux vars:", vars)
	// fmt.Println("symbols:", vars["symbols"])
	// fmt.Printf("Market: %v\nPair: %v\n", strings.ToLower(vars["market"]), strings.ToUpper(vars["pair"]))
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
			_convRates = ARRR_BTC_KC_RATES
		case "ARRR-USDT":
			_convRates = ARRR_USDT_KC_RATES
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "binance":
		switch strings.ToUpper(vars["pair"]) {
		case "BTC-USDT":
			if strings.ToLower(BTC_PRICE_SOURCE) != strings.ToLower(vars["market"]) {
				w.Header().Set("Content-Type", "application/json")
				fmt.Fprintf(w, `{"success": false,"error": "This market not enabled. Please try another market"}`)
				return
			}
			_convRates = BTC_USDT_BINANCE_RATES
		case "KMD-BTC":
			_convRates = KMD_BTC_BINANCE_RATES
		default:
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"success": false,"error": "Pair not found. Please try another pair"}`)
			return
		}
	case "coingecko":
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
	case "safetrade":
		switch strings.ToUpper(vars["pair"]) {
		case "VRSC-BTC":
			_convRates = VRSC_BTC_SAFETRADE_RATES
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

	var wg sync.WaitGroup
	wg.Add(1)
	go fixer(*fixerAPIToken, &wg)

	switch strings.ToLower(*btcPriceSource) {
	case "binance":
		wg.Add(1)
		go BtcUsdtBinanceAPI(&wg)
		BTC_PRICE_SOURCE = "Binance"
	case "coingecko":
		wg.Add(1)
		go BtcUsdCoinGeckoAPI(&wg)
		BTC_PRICE_SOURCE = "CoinGecko"
	}

	wg.Add(6)
	go ArrrToAPI(&wg)
	go ArrrBtcKcAPI(&wg)
	go ArrrUsdtKcAPI(&wg)
	go KmdBtcBinanceAPI(&wg)
	go ZecBtcBinanceAPI(&wg)
	go VrscBtcSafeTradeAPI(&wg)
	wg.Wait()

	go displayRates()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", indexPage)
	r.HandleFunc("/v1/price/{market}/{pair}", PriceAPI).Methods("GET").Queries("symbols", "{symbols}")
	r.HandleFunc("/v1/price/{market}/{pair}", PriceAPI).Methods("GET")
	r.HandleFunc("/v1/price/markets", Markets).Methods("GET")

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))

}

func displayRates() {
	for {
		// fmt.Printf("FIXER_RATES:\n %v\n", FIXER_RATES)
		log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		log.Println("BTC Price Source is:", BTC_PRICE_SOURCE)
		log.Printf("BTC USD: %.6f\n", BTC_RATES.Rates.USD)
		log.Printf("BTC NZD: %.6f\n", BTC_RATES.Rates.NZD)
		log.Printf("TradeOgre-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_TO_RATES.Rates.BTC)
		log.Printf("TradeOgre-ARRR/BTC: ARRR (USD): %.6f\n", ARRR_TO_RATES.Rates.USD)
		log.Printf("KuCoin-ARRR/BTC: ARRR (BTC): %.8f\n", ARRR_BTC_KC_RATES.Rates.BTC)
		log.Printf("KuCoin-ARRR/BTC: ARRR (USD): %.6f\n", ARRR_BTC_KC_RATES.Rates.USD)
		log.Printf("KuCoin-ARRR/USDT: ARRR (BTC): %.8f\n", ARRR_USDT_KC_RATES.Rates.BTC)
		log.Printf("KuCoin-ARRR/USDT: ARRR (USD): %.6f\n", ARRR_USDT_KC_RATES.Rates.USD)
		log.Printf("Binance-KMD/BTC: KMD (BTC) %.8f\n", KMD_BTC_BINANCE_RATES.Rates.BTC)
		log.Printf("Binance-KMD/BTC: KMD (USD) %.6f\n", KMD_BTC_BINANCE_RATES.Rates.USD)
		log.Printf("Binance-ZEC/BTC: ZEC (BTC) %.8f\n", ZEC_BTC_BINANCE_RATES.Rates.BTC)
		log.Printf("Binance-ZEC/BTC: ZEC (USD) %.6f\n", ZEC_BTC_BINANCE_RATES.Rates.USD)
		log.Printf("SafeTrade-VRSC/BTC: VRSC (BTC) %.8f\n", VRSC_BTC_SAFETRADE_RATES.Rates.BTC)
		log.Printf("SafeTrade-VRSC/BTC: VRSC (USD) %.6f\n", VRSC_BTC_SAFETRADE_RATES.Rates.USD)
		log.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		sleepSeconds := DISPLAY_RATES_SECONDS
		log.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
