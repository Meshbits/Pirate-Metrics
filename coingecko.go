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
		fmt.Println(result.([]interface{})[0].(map[string]interface{})["current_price"])

		var btc bitcoinRates
		btc.Base = "BTC"
		// btc.Rates.USD = result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AED = FIXER_RATES.Rates.AED * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AFN = FIXER_RATES.Rates.AFN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ALL = FIXER_RATES.Rates.ALL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AMD = FIXER_RATES.Rates.AMD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ANG = FIXER_RATES.Rates.ANG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AOA = FIXER_RATES.Rates.AOA * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ARS = FIXER_RATES.Rates.ARS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AUD = FIXER_RATES.Rates.AUD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AWG = FIXER_RATES.Rates.AWG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.AZN = FIXER_RATES.Rates.AZN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BAM = FIXER_RATES.Rates.BAM * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BBD = FIXER_RATES.Rates.BBD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BDT = FIXER_RATES.Rates.BDT * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BGN = FIXER_RATES.Rates.BGN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BHD = FIXER_RATES.Rates.BHD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BIF = FIXER_RATES.Rates.BIF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BMD = FIXER_RATES.Rates.BMD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BND = FIXER_RATES.Rates.BND * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BOB = FIXER_RATES.Rates.BOB * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BRL = FIXER_RATES.Rates.BRL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BSD = FIXER_RATES.Rates.BSD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BTC = FIXER_RATES.Rates.BTC * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BTN = FIXER_RATES.Rates.BTN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BWP = FIXER_RATES.Rates.BWP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BYN = FIXER_RATES.Rates.BYN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BYR = FIXER_RATES.Rates.BYR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.BZD = FIXER_RATES.Rates.BZD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CAD = FIXER_RATES.Rates.CAD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CDF = FIXER_RATES.Rates.CDF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CHF = FIXER_RATES.Rates.CHF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CLF = FIXER_RATES.Rates.CLF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CLP = FIXER_RATES.Rates.CLP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CNY = FIXER_RATES.Rates.CNY * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.COP = FIXER_RATES.Rates.COP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CRC = FIXER_RATES.Rates.CRC * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CUC = FIXER_RATES.Rates.CUC * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CUP = FIXER_RATES.Rates.CUP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CVE = FIXER_RATES.Rates.CVE * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.CZK = FIXER_RATES.Rates.CZK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.DJF = FIXER_RATES.Rates.DJF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.DKK = FIXER_RATES.Rates.DKK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.DOP = FIXER_RATES.Rates.DOP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.DZD = FIXER_RATES.Rates.DZD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.EGP = FIXER_RATES.Rates.EGP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ERN = FIXER_RATES.Rates.ERN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ETB = FIXER_RATES.Rates.ETB * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.EUR = FIXER_RATES.Rates.EUR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.FJD = FIXER_RATES.Rates.FJD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.FKP = FIXER_RATES.Rates.FKP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GBP = FIXER_RATES.Rates.GBP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GEL = FIXER_RATES.Rates.GEL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GGP = FIXER_RATES.Rates.GGP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GHS = FIXER_RATES.Rates.GHS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GIP = FIXER_RATES.Rates.GIP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GMD = FIXER_RATES.Rates.GMD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GNF = FIXER_RATES.Rates.GNF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GTQ = FIXER_RATES.Rates.GTQ * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.GYD = FIXER_RATES.Rates.GYD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.HKD = FIXER_RATES.Rates.HKD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.HNL = FIXER_RATES.Rates.HNL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.HRK = FIXER_RATES.Rates.HRK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.HTG = FIXER_RATES.Rates.HTG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.HUF = FIXER_RATES.Rates.HUF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.IDR = FIXER_RATES.Rates.IDR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ILS = FIXER_RATES.Rates.ILS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.IMP = FIXER_RATES.Rates.IMP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.INR = FIXER_RATES.Rates.INR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.IQD = FIXER_RATES.Rates.IQD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.IRR = FIXER_RATES.Rates.IRR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ISK = FIXER_RATES.Rates.ISK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.JEP = FIXER_RATES.Rates.JEP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.JMD = FIXER_RATES.Rates.JMD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.JOD = FIXER_RATES.Rates.JOD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.JPY = FIXER_RATES.Rates.JPY * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KES = FIXER_RATES.Rates.KES * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KGS = FIXER_RATES.Rates.KGS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KHR = FIXER_RATES.Rates.KHR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KMF = FIXER_RATES.Rates.KMF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KPW = FIXER_RATES.Rates.KPW * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KRW = FIXER_RATES.Rates.KRW * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KWD = FIXER_RATES.Rates.KWD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KYD = FIXER_RATES.Rates.KYD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.KZT = FIXER_RATES.Rates.KZT * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LAK = FIXER_RATES.Rates.LAK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LBP = FIXER_RATES.Rates.LBP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LKR = FIXER_RATES.Rates.LKR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LRD = FIXER_RATES.Rates.LRD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LSL = FIXER_RATES.Rates.LSL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LTL = FIXER_RATES.Rates.LTL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LVL = FIXER_RATES.Rates.LVL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.LYD = FIXER_RATES.Rates.LYD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MAD = FIXER_RATES.Rates.MAD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MDL = FIXER_RATES.Rates.MDL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MGA = FIXER_RATES.Rates.MGA * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MKD = FIXER_RATES.Rates.MKD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MMK = FIXER_RATES.Rates.MMK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MNT = FIXER_RATES.Rates.MNT * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MOP = FIXER_RATES.Rates.MOP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MRO = FIXER_RATES.Rates.MRO * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MUR = FIXER_RATES.Rates.MUR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MVR = FIXER_RATES.Rates.MVR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MWK = FIXER_RATES.Rates.MWK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MXN = FIXER_RATES.Rates.MXN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MYR = FIXER_RATES.Rates.MYR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.MZN = FIXER_RATES.Rates.MZN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NAD = FIXER_RATES.Rates.NAD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NGN = FIXER_RATES.Rates.NGN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NIO = FIXER_RATES.Rates.NIO * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NOK = FIXER_RATES.Rates.NOK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NPR = FIXER_RATES.Rates.NPR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.NZD = FIXER_RATES.Rates.NZD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.OMR = FIXER_RATES.Rates.OMR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PAB = FIXER_RATES.Rates.PAB * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PEN = FIXER_RATES.Rates.PEN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PGK = FIXER_RATES.Rates.PGK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PHP = FIXER_RATES.Rates.PHP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PKR = FIXER_RATES.Rates.PKR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PLN = FIXER_RATES.Rates.PLN * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.PYG = FIXER_RATES.Rates.PYG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.QAR = FIXER_RATES.Rates.QAR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.RON = FIXER_RATES.Rates.RON * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.RSD = FIXER_RATES.Rates.RSD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.RUB = FIXER_RATES.Rates.RUB * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.RWF = FIXER_RATES.Rates.RWF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SAR = FIXER_RATES.Rates.SAR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SBD = FIXER_RATES.Rates.SBD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SCR = FIXER_RATES.Rates.SCR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SDG = FIXER_RATES.Rates.SDG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SEK = FIXER_RATES.Rates.SEK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SGD = FIXER_RATES.Rates.SGD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SHP = FIXER_RATES.Rates.SHP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SLL = FIXER_RATES.Rates.SLL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SOS = FIXER_RATES.Rates.SOS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SRD = FIXER_RATES.Rates.SRD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.STD = FIXER_RATES.Rates.STD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SVC = FIXER_RATES.Rates.SVC * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SYP = FIXER_RATES.Rates.SYP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.SZL = FIXER_RATES.Rates.SZL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.THB = FIXER_RATES.Rates.THB * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TJS = FIXER_RATES.Rates.TJS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TMT = FIXER_RATES.Rates.TMT * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TND = FIXER_RATES.Rates.TND * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TOP = FIXER_RATES.Rates.TOP * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TRY = FIXER_RATES.Rates.TRY * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TTD = FIXER_RATES.Rates.TTD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TWD = FIXER_RATES.Rates.TWD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.TZS = FIXER_RATES.Rates.TZS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.UAH = FIXER_RATES.Rates.UAH * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.UGX = FIXER_RATES.Rates.UGX * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.USD = FIXER_RATES.Rates.USD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.UYU = FIXER_RATES.Rates.UYU * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.UZS = FIXER_RATES.Rates.UZS * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.VEF = FIXER_RATES.Rates.VEF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.VND = FIXER_RATES.Rates.VND * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.VUV = FIXER_RATES.Rates.VUV * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.WST = FIXER_RATES.Rates.WST * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XAF = FIXER_RATES.Rates.XAF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XAG = FIXER_RATES.Rates.XAG * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XAU = FIXER_RATES.Rates.XAU * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XCD = FIXER_RATES.Rates.XCD * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XDR = FIXER_RATES.Rates.XDR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XOF = FIXER_RATES.Rates.XOF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.XPF = FIXER_RATES.Rates.XPF * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.YER = FIXER_RATES.Rates.YER * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ZAR = FIXER_RATES.Rates.ZAR * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ZMK = FIXER_RATES.Rates.ZMK * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ZMW = FIXER_RATES.Rates.ZMW * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)
		btc.Rates.ZWL = FIXER_RATES.Rates.ZWL * result.([]interface{})[0].(map[string]interface{})["current_price"].(float64)

		BTC_RATES = btc

		// b, _ := json.Marshal(btc)
		// fmt.Println(string(b))

		sleepSeconds := 30
		fmt.Printf("Updated Bitcoin Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return
	}
}
