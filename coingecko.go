package main

import (
	"fmt"
	"time"
)

var BTC_RATES bitcoinRates

type bitcoinRates struct {
	Base  string `json:"base"`
	Rates struct {
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

func CGeckoBTCAPI() {
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "https://api.coingecko.com/api/v3/coins/markets"
		params := `vs_currency=usd&ids=bitcoin&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=1h%2C%2024h%2C%207d%2C%2030d`
		result, err := RPCResultMap(url, params)
		if err != nil {
			fmt.Printf("Got error fetching Fixer rates: %v\n", err)
		}
		// fmt.Println(result)
		fmt.Printf("BTC Price (USD): %v\n", result.([]interface{})[0].(map[string]interface{})["current_price"])

		var btc bitcoinRates
		btc.Base = "USD"
		btc.Rates.AED = toFixed(FIXER_RATES.Rates.AED*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AFN = toFixed(FIXER_RATES.Rates.AFN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ALL = toFixed(FIXER_RATES.Rates.ALL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AMD = toFixed(FIXER_RATES.Rates.AMD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ANG = toFixed(FIXER_RATES.Rates.ANG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AOA = toFixed(FIXER_RATES.Rates.AOA*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ARS = toFixed(FIXER_RATES.Rates.ARS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AUD = toFixed(FIXER_RATES.Rates.AUD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AWG = toFixed(FIXER_RATES.Rates.AWG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.AZN = toFixed(FIXER_RATES.Rates.AZN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BAM = toFixed(FIXER_RATES.Rates.BAM*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BBD = toFixed(FIXER_RATES.Rates.BBD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BDT = toFixed(FIXER_RATES.Rates.BDT*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BGN = toFixed(FIXER_RATES.Rates.BGN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BHD = toFixed(FIXER_RATES.Rates.BHD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BIF = toFixed(FIXER_RATES.Rates.BIF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BMD = toFixed(FIXER_RATES.Rates.BMD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BND = toFixed(FIXER_RATES.Rates.BND*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BOB = toFixed(FIXER_RATES.Rates.BOB*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BRL = toFixed(FIXER_RATES.Rates.BRL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BSD = toFixed(FIXER_RATES.Rates.BSD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BTC = toFixed(FIXER_RATES.Rates.BTC*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 8)
		btc.Rates.BTN = toFixed(FIXER_RATES.Rates.BTN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BWP = toFixed(FIXER_RATES.Rates.BWP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BYN = toFixed(FIXER_RATES.Rates.BYN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BYR = toFixed(FIXER_RATES.Rates.BYR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.BZD = toFixed(FIXER_RATES.Rates.BZD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CAD = toFixed(FIXER_RATES.Rates.CAD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CDF = toFixed(FIXER_RATES.Rates.CDF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CHF = toFixed(FIXER_RATES.Rates.CHF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CLF = toFixed(FIXER_RATES.Rates.CLF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CLP = toFixed(FIXER_RATES.Rates.CLP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CNY = toFixed(FIXER_RATES.Rates.CNY*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.COP = toFixed(FIXER_RATES.Rates.COP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CRC = toFixed(FIXER_RATES.Rates.CRC*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CUC = toFixed(FIXER_RATES.Rates.CUC*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CUP = toFixed(FIXER_RATES.Rates.CUP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CVE = toFixed(FIXER_RATES.Rates.CVE*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.CZK = toFixed(FIXER_RATES.Rates.CZK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.DJF = toFixed(FIXER_RATES.Rates.DJF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.DKK = toFixed(FIXER_RATES.Rates.DKK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.DOP = toFixed(FIXER_RATES.Rates.DOP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.DZD = toFixed(FIXER_RATES.Rates.DZD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.EGP = toFixed(FIXER_RATES.Rates.EGP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ERN = toFixed(FIXER_RATES.Rates.ERN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ETB = toFixed(FIXER_RATES.Rates.ETB*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.EUR = toFixed(FIXER_RATES.Rates.EUR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.FJD = toFixed(FIXER_RATES.Rates.FJD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.FKP = toFixed(FIXER_RATES.Rates.FKP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GBP = toFixed(FIXER_RATES.Rates.GBP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GEL = toFixed(FIXER_RATES.Rates.GEL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GGP = toFixed(FIXER_RATES.Rates.GGP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GHS = toFixed(FIXER_RATES.Rates.GHS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GIP = toFixed(FIXER_RATES.Rates.GIP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GMD = toFixed(FIXER_RATES.Rates.GMD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GNF = toFixed(FIXER_RATES.Rates.GNF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GTQ = toFixed(FIXER_RATES.Rates.GTQ*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.GYD = toFixed(FIXER_RATES.Rates.GYD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.HKD = toFixed(FIXER_RATES.Rates.HKD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.HNL = toFixed(FIXER_RATES.Rates.HNL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.HRK = toFixed(FIXER_RATES.Rates.HRK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.HTG = toFixed(FIXER_RATES.Rates.HTG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.HUF = toFixed(FIXER_RATES.Rates.HUF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.IDR = toFixed(FIXER_RATES.Rates.IDR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ILS = toFixed(FIXER_RATES.Rates.ILS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.IMP = toFixed(FIXER_RATES.Rates.IMP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.INR = toFixed(FIXER_RATES.Rates.INR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.IQD = toFixed(FIXER_RATES.Rates.IQD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.IRR = toFixed(FIXER_RATES.Rates.IRR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ISK = toFixed(FIXER_RATES.Rates.ISK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.JEP = toFixed(FIXER_RATES.Rates.JEP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.JMD = toFixed(FIXER_RATES.Rates.JMD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.JOD = toFixed(FIXER_RATES.Rates.JOD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.JPY = toFixed(FIXER_RATES.Rates.JPY*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KES = toFixed(FIXER_RATES.Rates.KES*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KGS = toFixed(FIXER_RATES.Rates.KGS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KHR = toFixed(FIXER_RATES.Rates.KHR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KMF = toFixed(FIXER_RATES.Rates.KMF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KPW = toFixed(FIXER_RATES.Rates.KPW*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KRW = toFixed(FIXER_RATES.Rates.KRW*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KWD = toFixed(FIXER_RATES.Rates.KWD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KYD = toFixed(FIXER_RATES.Rates.KYD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.KZT = toFixed(FIXER_RATES.Rates.KZT*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LAK = toFixed(FIXER_RATES.Rates.LAK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LBP = toFixed(FIXER_RATES.Rates.LBP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LKR = toFixed(FIXER_RATES.Rates.LKR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LRD = toFixed(FIXER_RATES.Rates.LRD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LSL = toFixed(FIXER_RATES.Rates.LSL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LTL = toFixed(FIXER_RATES.Rates.LTL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LVL = toFixed(FIXER_RATES.Rates.LVL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.LYD = toFixed(FIXER_RATES.Rates.LYD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MAD = toFixed(FIXER_RATES.Rates.MAD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MDL = toFixed(FIXER_RATES.Rates.MDL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MGA = toFixed(FIXER_RATES.Rates.MGA*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MKD = toFixed(FIXER_RATES.Rates.MKD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MMK = toFixed(FIXER_RATES.Rates.MMK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MNT = toFixed(FIXER_RATES.Rates.MNT*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MOP = toFixed(FIXER_RATES.Rates.MOP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MRO = toFixed(FIXER_RATES.Rates.MRO*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MUR = toFixed(FIXER_RATES.Rates.MUR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MVR = toFixed(FIXER_RATES.Rates.MVR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MWK = toFixed(FIXER_RATES.Rates.MWK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MXN = toFixed(FIXER_RATES.Rates.MXN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MYR = toFixed(FIXER_RATES.Rates.MYR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.MZN = toFixed(FIXER_RATES.Rates.MZN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NAD = toFixed(FIXER_RATES.Rates.NAD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NGN = toFixed(FIXER_RATES.Rates.NGN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NIO = toFixed(FIXER_RATES.Rates.NIO*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NOK = toFixed(FIXER_RATES.Rates.NOK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NPR = toFixed(FIXER_RATES.Rates.NPR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.NZD = toFixed(FIXER_RATES.Rates.NZD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.OMR = toFixed(FIXER_RATES.Rates.OMR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PAB = toFixed(FIXER_RATES.Rates.PAB*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PEN = toFixed(FIXER_RATES.Rates.PEN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PGK = toFixed(FIXER_RATES.Rates.PGK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PHP = toFixed(FIXER_RATES.Rates.PHP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PKR = toFixed(FIXER_RATES.Rates.PKR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PLN = toFixed(FIXER_RATES.Rates.PLN*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.PYG = toFixed(FIXER_RATES.Rates.PYG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.QAR = toFixed(FIXER_RATES.Rates.QAR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.RON = toFixed(FIXER_RATES.Rates.RON*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.RSD = toFixed(FIXER_RATES.Rates.RSD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.RUB = toFixed(FIXER_RATES.Rates.RUB*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.RWF = toFixed(FIXER_RATES.Rates.RWF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SAR = toFixed(FIXER_RATES.Rates.SAR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SBD = toFixed(FIXER_RATES.Rates.SBD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SCR = toFixed(FIXER_RATES.Rates.SCR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SDG = toFixed(FIXER_RATES.Rates.SDG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SEK = toFixed(FIXER_RATES.Rates.SEK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SGD = toFixed(FIXER_RATES.Rates.SGD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SHP = toFixed(FIXER_RATES.Rates.SHP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SLL = toFixed(FIXER_RATES.Rates.SLL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SOS = toFixed(FIXER_RATES.Rates.SOS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SRD = toFixed(FIXER_RATES.Rates.SRD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.STD = toFixed(FIXER_RATES.Rates.STD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SVC = toFixed(FIXER_RATES.Rates.SVC*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SYP = toFixed(FIXER_RATES.Rates.SYP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.SZL = toFixed(FIXER_RATES.Rates.SZL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.THB = toFixed(FIXER_RATES.Rates.THB*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TJS = toFixed(FIXER_RATES.Rates.TJS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TMT = toFixed(FIXER_RATES.Rates.TMT*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TND = toFixed(FIXER_RATES.Rates.TND*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TOP = toFixed(FIXER_RATES.Rates.TOP*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TRY = toFixed(FIXER_RATES.Rates.TRY*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TTD = toFixed(FIXER_RATES.Rates.TTD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TWD = toFixed(FIXER_RATES.Rates.TWD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.TZS = toFixed(FIXER_RATES.Rates.TZS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.UAH = toFixed(FIXER_RATES.Rates.UAH*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.UGX = toFixed(FIXER_RATES.Rates.UGX*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.USD = toFixed(FIXER_RATES.Rates.USD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.UYU = toFixed(FIXER_RATES.Rates.UYU*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.UZS = toFixed(FIXER_RATES.Rates.UZS*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.VEF = toFixed(FIXER_RATES.Rates.VEF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.VND = toFixed(FIXER_RATES.Rates.VND*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.VUV = toFixed(FIXER_RATES.Rates.VUV*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.WST = toFixed(FIXER_RATES.Rates.WST*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XAF = toFixed(FIXER_RATES.Rates.XAF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XAG = toFixed(FIXER_RATES.Rates.XAG*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XAU = toFixed(FIXER_RATES.Rates.XAU*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XCD = toFixed(FIXER_RATES.Rates.XCD*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XDR = toFixed(FIXER_RATES.Rates.XDR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XOF = toFixed(FIXER_RATES.Rates.XOF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.XPF = toFixed(FIXER_RATES.Rates.XPF*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.YER = toFixed(FIXER_RATES.Rates.YER*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ZAR = toFixed(FIXER_RATES.Rates.ZAR*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ZMK = toFixed(FIXER_RATES.Rates.ZMK*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ZMW = toFixed(FIXER_RATES.Rates.ZMW*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)
		btc.Rates.ZWL = toFixed(FIXER_RATES.Rates.ZWL*result.([]interface{})[0].(map[string]interface{})["current_price"].(float64), 6)

		BTC_RATES = btc

		// b, _ := json.Marshal(btc)
		// fmt.Println(string(b))

		sleepSeconds := 30
		fmt.Printf("Updated Bitcoin Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return
	}
}
