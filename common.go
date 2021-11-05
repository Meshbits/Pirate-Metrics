package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"sync"
	"time"
)

var BTC_RATES ConversionRates

var (
	MARKETS_AVAILABLE     = map[string][]string{}
	BTC_PRICE_SOURCE      string
	BINANCE_SECONDS       int = 60 * 2
	KUCOIN_SECONDS        int = 60 * 2
	TRADEOGRE_SECONDS     int = 60
	SAFETRADE_SECONDS     int = 60
	COINGECKO_SECONDS     int = 60 * 2
	FIXER_SECONDS         int = 60 * 60 * 4
	DISPLAY_RATES_SECONDS int = 30
	mutex                 sync.Mutex
)

type ConversionRates struct {
	Success   bool   `json:"success"`
	Timestamp int64  `json:"timestamp"`
	Base      string `json:"base"`
	Date      string `json:"date,omitempty"`
	Market    string `json:"market,omitempty"`
	Rates     struct {
		AED float64 `json:"AED"`
		AFN float64 `json:"AFN"`
		ALL float64 `json:"ALL"`
		AMD float64 `json:"AMD"`
		ANG float64 `json:"ANG"`
		AOA float64 `json:"AOA"`
		ARS float64 `json:"ARS"`
		AUD float64 `json:"AUD"`
		AWG float64 `json:"AWG"`
		AZN float64 `json:"AZN"`
		BAM float64 `json:"BAM"`
		BBD float64 `json:"BBD"`
		BDT float64 `json:"BDT"`
		BGN float64 `json:"BGN"`
		BHD float64 `json:"BHD"`
		BIF float64 `json:"BIF"`
		BMD float64 `json:"BMD"`
		BND float64 `json:"BND"`
		BOB float64 `json:"BOB"`
		BRL float64 `json:"BRL"`
		BSD float64 `json:"BSD"`
		BTC float64 `json:"BTC"`
		BTN float64 `json:"BTN"`
		BWP float64 `json:"BWP"`
		BYN float64 `json:"BYN"`
		BYR float64 `json:"BYR"`
		BZD float64 `json:"BZD"`
		CAD float64 `json:"CAD"`
		CDF float64 `json:"CDF"`
		CHF float64 `json:"CHF"`
		CLF float64 `json:"CLF"`
		CLP float64 `json:"CLP"`
		CNY float64 `json:"CNY"`
		COP float64 `json:"COP"`
		CRC float64 `json:"CRC"`
		CUC float64 `json:"CUC"`
		CUP float64 `json:"CUP"`
		CVE float64 `json:"CVE"`
		CZK float64 `json:"CZK"`
		DJF float64 `json:"DJF"`
		DKK float64 `json:"DKK"`
		DOP float64 `json:"DOP"`
		DZD float64 `json:"DZD"`
		EGP float64 `json:"EGP"`
		ERN float64 `json:"ERN"`
		ETB float64 `json:"ETB"`
		EUR float64 `json:"EUR"`
		FJD float64 `json:"FJD"`
		FKP float64 `json:"FKP"`
		GBP float64 `json:"GBP"`
		GEL float64 `json:"GEL"`
		GGP float64 `json:"GGP"`
		GHS float64 `json:"GHS"`
		GIP float64 `json:"GIP"`
		GMD float64 `json:"GMD"`
		GNF float64 `json:"GNF"`
		GTQ float64 `json:"GTQ"`
		GYD float64 `json:"GYD"`
		HKD float64 `json:"HKD"`
		HNL float64 `json:"HNL"`
		HRK float64 `json:"HRK"`
		HTG float64 `json:"HTG"`
		HUF float64 `json:"HUF"`
		IDR float64 `json:"IDR"`
		ILS float64 `json:"ILS"`
		IMP float64 `json:"IMP"`
		INR float64 `json:"INR"`
		IQD float64 `json:"IQD"`
		IRR float64 `json:"IRR"`
		ISK float64 `json:"ISK"`
		JEP float64 `json:"JEP"`
		JMD float64 `json:"JMD"`
		JOD float64 `json:"JOD"`
		JPY float64 `json:"JPY"`
		KES float64 `json:"KES"`
		KGS float64 `json:"KGS"`
		KHR float64 `json:"KHR"`
		KMF float64 `json:"KMF"`
		KPW float64 `json:"KPW"`
		KRW float64 `json:"KRW"`
		KWD float64 `json:"KWD"`
		KYD float64 `json:"KYD"`
		KZT float64 `json:"KZT"`
		LAK float64 `json:"LAK"`
		LBP float64 `json:"LBP"`
		LKR float64 `json:"LKR"`
		LRD float64 `json:"LRD"`
		LSL float64 `json:"LSL"`
		LTL float64 `json:"LTL"`
		LVL float64 `json:"LVL"`
		LYD float64 `json:"LYD"`
		MAD float64 `json:"MAD"`
		MDL float64 `json:"MDL"`
		MGA float64 `json:"MGA"`
		MKD float64 `json:"MKD"`
		MMK float64 `json:"MMK"`
		MNT float64 `json:"MNT"`
		MOP float64 `json:"MOP"`
		MRO float64 `json:"MRO"`
		MUR float64 `json:"MUR"`
		MVR float64 `json:"MVR"`
		MWK float64 `json:"MWK"`
		MXN float64 `json:"MXN"`
		MYR float64 `json:"MYR"`
		MZN float64 `json:"MZN"`
		NAD float64 `json:"NAD"`
		NGN float64 `json:"NGN"`
		NIO float64 `json:"NIO"`
		NOK float64 `json:"NOK"`
		NPR float64 `json:"NPR"`
		NZD float64 `json:"NZD"`
		OMR float64 `json:"OMR"`
		PAB float64 `json:"PAB"`
		PEN float64 `json:"PEN"`
		PGK float64 `json:"PGK"`
		PHP float64 `json:"PHP"`
		PKR float64 `json:"PKR"`
		PLN float64 `json:"PLN"`
		PYG float64 `json:"PYG"`
		QAR float64 `json:"QAR"`
		RON float64 `json:"RON"`
		RSD float64 `json:"RSD"`
		RUB float64 `json:"RUB"`
		RWF float64 `json:"RWF"`
		SAR float64 `json:"SAR"`
		SBD float64 `json:"SBD"`
		SCR float64 `json:"SCR"`
		SDG float64 `json:"SDG"`
		SEK float64 `json:"SEK"`
		SGD float64 `json:"SGD"`
		SHP float64 `json:"SHP"`
		SLL float64 `json:"SLL"`
		SOS float64 `json:"SOS"`
		SRD float64 `json:"SRD"`
		STD float64 `json:"STD"`
		SVC float64 `json:"SVC"`
		SYP float64 `json:"SYP"`
		SZL float64 `json:"SZL"`
		THB float64 `json:"THB"`
		TJS float64 `json:"TJS"`
		TMT float64 `json:"TMT"`
		TND float64 `json:"TND"`
		TOP float64 `json:"TOP"`
		TRY float64 `json:"TRY"`
		TTD float64 `json:"TTD"`
		TWD float64 `json:"TWD"`
		TZS float64 `json:"TZS"`
		UAH float64 `json:"UAH"`
		UGX float64 `json:"UGX"`
		USD float64 `json:"USD"`
		UYU float64 `json:"UYU"`
		UZS float64 `json:"UZS"`
		VEF float64 `json:"VEF"`
		VND float64 `json:"VND"`
		VUV float64 `json:"VUV"`
		WST float64 `json:"WST"`
		XAF float64 `json:"XAF"`
		XAG float64 `json:"XAG"`
		XAU float64 `json:"XAU"`
		XCD float64 `json:"XCD"`
		XDR float64 `json:"XDR"`
		XOF float64 `json:"XOF"`
		XPF float64 `json:"XPF"`
		YER float64 `json:"YER"`
		ZAR float64 `json:"ZAR"`
		ZMK float64 `json:"ZMK"`
		ZMW float64 `json:"ZMW"`
		ZWL float64 `json:"ZWL"`
	} `json:"rates"`
}

type queriedRates struct {
	// Success   bool        `json:"success"`
	Timestamp int64       `json:"timestamp"`
	Base      string      `json:"base"`
	Market    string      `json:"market"`
	Rates     interface{} `json:"rates"`
}

// APIQuery holds the URL and params which is sent to APICall to query rates
type APIQuery struct {
	Url    string `json:"method"`
	Params string `json:"params"`
}

func APICall(q *APIQuery) (string, error) {

	// url := "http://data.fixer.io/api/latest?access_key=YOUR_FIXER_API_KEY&format=1"
	// url := "https://tradeogre.com/api/v1/ticker/BTC-ARRR"
	url := q.Url + q.Params
	// fmt.Println(q.Url)
	// fmt.Println(url)
	method := "GET"

	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	// fmt.Println("body:", string(body))

	s := string(body)
	return s, nil
}

// RPCResultMap using golang's own http package
func RPCResultMap(url string, params interface{}) (interface{}, error) {

	// fmt.Printf("params -- %+v\n", params)

	paramsStr := fmt.Sprintf("%v", params)

	query := APIQuery{
		Url:    url,
		Params: string(paramsStr),
	}
	// fmt.Printf("%+v\n", query)

	getJSON, err := APICall(&query)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	// fmt.Printf("getJSON %T\n", getJSON)
	// fmt.Printf("getJSON %v\n", getJSON)

	var result interface{}
	json.Unmarshal([]byte(getJSON), &result)
	// fmt.Println("Error:", result["error"])
	// fmt.Println("success:", result["success"])
	// if result.(map[string]interface{})["error"] != nil {
	// 	// fmt.Printf("error is not nil: %v\n", result["error"])
	// 	return nil, errors.New(result.(map[string]interface{})["error"].(string))
	// }
	return result, nil
	// return "", nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}

func updateMarketsAvailable(market, pair string, wg *sync.WaitGroup) {
	mutex.Lock()
	if _, ok := MARKETS_AVAILABLE[market]; ok {
		for _, ma := range MARKETS_AVAILABLE[market] {
			if ma != pair {
				MARKETS_AVAILABLE[market] = append(MARKETS_AVAILABLE[market], pair)
			}
		}
	} else {
		MARKETS_AVAILABLE[market] = append(MARKETS_AVAILABLE[market], pair)
	}
	mutex.Unlock()
	wg.Done()
}
