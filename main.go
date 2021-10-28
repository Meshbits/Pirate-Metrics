package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	// url := "https://tradeogre.com/api/v1/ticker/BTC-ARRR"
	// params := ""
	// result, err := RPCResultMap(url, params)
	// if err != nil {
	// 	fmt.Printf("Got error: %v\n", err)
	// 	return
	// }
	// fmt.Println(result)
	// // fmt.Printf("initialprice: %v\n", result.(map[string]interface{})["initialprice"].(string))
	// // fmt.Printf("price: %v\n", result.(map[string]interface{})["price"].(string))

	fixerAPIToken := flag.String("fixerAPIToken", "", "Get the API token from https://fixer.io. API Token from a free sign-up is sufficient to use with this application.")
	flag.Parse()
	fmt.Println("fixerAPIToken:", *fixerAPIToken)
	if *fixerAPIToken == "" {
		return
	}

	go fixer(*fixerAPIToken)
	go CGeckoBTCAPI()

	go displayRates()

	fmt.Scanln()

}

func displayRates() {
	for {
		fmt.Printf("FIXER_RATES:\n %v\n", FIXER_RATES)
		fmt.Printf("BTC_RATES:\n %v\n", BTC_RATES)

		sleepSeconds := 10
		fmt.Printf("Will Display Rates every other %v seconds...\n\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
