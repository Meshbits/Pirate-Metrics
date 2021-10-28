package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArrrPrice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	arrrPrice, _ := json.Marshal(ARRR_RATES)
	w.Write(arrrPrice)
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
	go TOgreARRRAPI()

	go displayRates()

	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", ArrrPrice).Methods("GET")

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
		fmt.Printf("ARRR USD: %v\n", ARRR_RATES.Rates.USD)
		fmt.Printf("ARRR NZD: %v\n", ARRR_RATES.Rates.NZD)
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")

		sleepSeconds := 10
		fmt.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
