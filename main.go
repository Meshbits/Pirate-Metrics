package main

import (
	"flag"
	"fmt"
	"time"
)

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

	displayRates()

	// fmt.Scanln()

}

func displayRates() {
	for {
		// fmt.Printf("FIXER_RATES:\n %v\n", FIXER_RATES)
		// fmt.Printf("BTC_RATES:\n %v\n", BTC_RATES)
		fmt.Printf("ARRR_RATES:\n %v\n", ARRR_RATES)

		sleepSeconds := 10
		fmt.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
