package main

import (
	"fmt"
	"time"
)

var FIXER_RATES fixerRates

type fixerRates struct {
	Success   bool    `json:"success"`
	Timestamp float64 `json:"timestamp"`
	Base      string  `json:"base"`
	Date      string  `json:"date"`
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

func fixer(APIToken string) /*(fixerRates, error)*/ {
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "http://data.fixer.io/api/latest?"
		params := `access_key=` + APIToken + `&format=1`
		result, err := RPCResultMap(url, params)
		if err != nil {
			fmt.Printf("Got error fetching Fixer rates: %v\n", err)
			// return fixerRates{Success: false}, errors.New(err.Error())
		}
		// fmt.Println(result)

		var fx fixerRates
		var baseUSD float64 = 1
		rateEUR := baseUSD / result.(map[string]interface{})["rates"].(map[string]interface{})["USD"].(float64)
		fx.Success = result.(map[string]interface{})["success"].(bool)
		fx.Timestamp = result.(map[string]interface{})["timestamp"].(float64)
		fx.Date = result.(map[string]interface{})["date"].(string)
		fx.Base = "USD"
		fx.Rates.AED = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AED"].(float64)), 6)
		fx.Rates.AFN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AFN"].(float64)), 6)
		fx.Rates.ALL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ALL"].(float64)), 6)
		fx.Rates.AMD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AMD"].(float64)), 6)
		fx.Rates.ANG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ANG"].(float64)), 6)
		fx.Rates.AOA = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AOA"].(float64)), 6)
		fx.Rates.ARS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ARS"].(float64)), 6)
		fx.Rates.AUD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AUD"].(float64)), 6)
		fx.Rates.AWG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AWG"].(float64)), 6)
		fx.Rates.AZN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["AZN"].(float64)), 6)
		fx.Rates.BAM = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BAM"].(float64)), 6)
		fx.Rates.BBD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BBD"].(float64)), 6)
		fx.Rates.BDT = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BDT"].(float64)), 6)
		fx.Rates.BGN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BGN"].(float64)), 6)
		fx.Rates.BHD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BHD"].(float64)), 6)
		fx.Rates.BIF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BIF"].(float64)), 6)
		fx.Rates.BMD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BMD"].(float64)), 6)
		fx.Rates.BND = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BND"].(float64)), 6)
		fx.Rates.BOB = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BOB"].(float64)), 6)
		fx.Rates.BRL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BRL"].(float64)), 6)
		fx.Rates.BSD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BSD"].(float64)), 6)
		fx.Rates.BTC = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BTC"].(float64)), 8)
		fx.Rates.BTN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BTN"].(float64)), 6)
		fx.Rates.BWP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BWP"].(float64)), 6)
		fx.Rates.BYN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BYN"].(float64)), 6)
		fx.Rates.BYR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BYR"].(float64)), 6)
		fx.Rates.BZD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["BZD"].(float64)), 6)
		fx.Rates.CAD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CAD"].(float64)), 6)
		fx.Rates.CDF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CDF"].(float64)), 6)
		fx.Rates.CHF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CHF"].(float64)), 6)
		fx.Rates.CLF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CLF"].(float64)), 6)
		fx.Rates.CLP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CLP"].(float64)), 6)
		fx.Rates.CNY = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CNY"].(float64)), 6)
		fx.Rates.COP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["COP"].(float64)), 6)
		fx.Rates.CRC = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CRC"].(float64)), 6)
		fx.Rates.CUC = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CUC"].(float64)), 6)
		fx.Rates.CUP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CUP"].(float64)), 6)
		fx.Rates.CVE = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CVE"].(float64)), 6)
		fx.Rates.CZK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["CZK"].(float64)), 6)
		fx.Rates.DJF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["DJF"].(float64)), 6)
		fx.Rates.DKK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["DKK"].(float64)), 6)
		fx.Rates.DOP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["DOP"].(float64)), 6)
		fx.Rates.DZD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["DZD"].(float64)), 6)
		fx.Rates.EGP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["EGP"].(float64)), 6)
		fx.Rates.ERN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ERN"].(float64)), 6)
		fx.Rates.ETB = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ETB"].(float64)), 6)
		fx.Rates.EUR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["EUR"].(float64)), 6)
		fx.Rates.FJD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["FJD"].(float64)), 6)
		fx.Rates.FKP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["FKP"].(float64)), 6)
		fx.Rates.GBP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GBP"].(float64)), 6)
		fx.Rates.GEL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GEL"].(float64)), 6)
		fx.Rates.GGP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GGP"].(float64)), 6)
		fx.Rates.GHS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GHS"].(float64)), 6)
		fx.Rates.GIP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GIP"].(float64)), 6)
		fx.Rates.GMD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GMD"].(float64)), 6)
		fx.Rates.GNF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GNF"].(float64)), 6)
		fx.Rates.GTQ = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GTQ"].(float64)), 6)
		fx.Rates.GYD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["GYD"].(float64)), 6)
		fx.Rates.HKD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["HKD"].(float64)), 6)
		fx.Rates.HNL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["HNL"].(float64)), 6)
		fx.Rates.HRK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["HRK"].(float64)), 6)
		fx.Rates.HTG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["HTG"].(float64)), 6)
		fx.Rates.HUF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["HUF"].(float64)), 6)
		fx.Rates.IDR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["IDR"].(float64)), 6)
		fx.Rates.ILS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ILS"].(float64)), 6)
		fx.Rates.IMP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["IMP"].(float64)), 6)
		fx.Rates.INR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["INR"].(float64)), 6)
		fx.Rates.IQD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["IQD"].(float64)), 6)
		fx.Rates.IRR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["IRR"].(float64)), 6)
		fx.Rates.ISK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ISK"].(float64)), 6)
		fx.Rates.JEP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["JEP"].(float64)), 6)
		fx.Rates.JMD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["JMD"].(float64)), 6)
		fx.Rates.JOD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["JOD"].(float64)), 6)
		fx.Rates.JPY = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["JPY"].(float64)), 6)
		fx.Rates.KES = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KES"].(float64)), 6)
		fx.Rates.KGS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KGS"].(float64)), 6)
		fx.Rates.KHR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KHR"].(float64)), 6)
		fx.Rates.KMF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KMF"].(float64)), 6)
		fx.Rates.KPW = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KPW"].(float64)), 6)
		fx.Rates.KRW = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KRW"].(float64)), 6)
		fx.Rates.KWD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KWD"].(float64)), 6)
		fx.Rates.KYD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KYD"].(float64)), 6)
		fx.Rates.KZT = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["KZT"].(float64)), 6)
		fx.Rates.LAK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LAK"].(float64)), 6)
		fx.Rates.LBP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LBP"].(float64)), 6)
		fx.Rates.LKR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LKR"].(float64)), 6)
		fx.Rates.LRD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LRD"].(float64)), 6)
		fx.Rates.LSL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LSL"].(float64)), 6)
		fx.Rates.LTL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LTL"].(float64)), 6)
		fx.Rates.LVL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LVL"].(float64)), 6)
		fx.Rates.LYD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["LYD"].(float64)), 6)
		fx.Rates.MAD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MAD"].(float64)), 6)
		fx.Rates.MDL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MDL"].(float64)), 6)
		fx.Rates.MGA = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MGA"].(float64)), 6)
		fx.Rates.MKD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MKD"].(float64)), 6)
		fx.Rates.MMK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MMK"].(float64)), 6)
		fx.Rates.MNT = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MNT"].(float64)), 6)
		fx.Rates.MOP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MOP"].(float64)), 6)
		fx.Rates.MRO = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MRO"].(float64)), 6)
		fx.Rates.MUR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MUR"].(float64)), 6)
		fx.Rates.MVR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MVR"].(float64)), 6)
		fx.Rates.MWK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MWK"].(float64)), 6)
		fx.Rates.MXN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MXN"].(float64)), 6)
		fx.Rates.MYR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MYR"].(float64)), 6)
		fx.Rates.MZN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["MZN"].(float64)), 6)
		fx.Rates.NAD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NAD"].(float64)), 6)
		fx.Rates.NGN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NGN"].(float64)), 6)
		fx.Rates.NIO = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NIO"].(float64)), 6)
		fx.Rates.NOK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NOK"].(float64)), 6)
		fx.Rates.NPR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NPR"].(float64)), 6)
		fx.Rates.NZD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["NZD"].(float64)), 6)
		fx.Rates.OMR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["OMR"].(float64)), 6)
		fx.Rates.PAB = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PAB"].(float64)), 6)
		fx.Rates.PEN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PEN"].(float64)), 6)
		fx.Rates.PGK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PGK"].(float64)), 6)
		fx.Rates.PHP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PHP"].(float64)), 6)
		fx.Rates.PKR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PKR"].(float64)), 6)
		fx.Rates.PLN = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PLN"].(float64)), 6)
		fx.Rates.PYG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["PYG"].(float64)), 6)
		fx.Rates.QAR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["QAR"].(float64)), 6)
		fx.Rates.RON = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["RON"].(float64)), 6)
		fx.Rates.RSD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["RSD"].(float64)), 6)
		fx.Rates.RUB = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["RUB"].(float64)), 6)
		fx.Rates.RWF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["RWF"].(float64)), 6)
		fx.Rates.SAR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SAR"].(float64)), 6)
		fx.Rates.SBD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SBD"].(float64)), 6)
		fx.Rates.SCR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SCR"].(float64)), 6)
		fx.Rates.SDG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SDG"].(float64)), 6)
		fx.Rates.SEK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SEK"].(float64)), 6)
		fx.Rates.SGD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SGD"].(float64)), 6)
		fx.Rates.SHP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SHP"].(float64)), 6)
		fx.Rates.SLL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SLL"].(float64)), 6)
		fx.Rates.SOS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SOS"].(float64)), 6)
		fx.Rates.SRD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SRD"].(float64)), 6)
		fx.Rates.STD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["STD"].(float64)), 6)
		fx.Rates.SVC = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SVC"].(float64)), 6)
		fx.Rates.SYP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SYP"].(float64)), 6)
		fx.Rates.SZL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["SZL"].(float64)), 6)
		fx.Rates.THB = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["THB"].(float64)), 6)
		fx.Rates.TJS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TJS"].(float64)), 6)
		fx.Rates.TMT = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TMT"].(float64)), 6)
		fx.Rates.TND = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TND"].(float64)), 6)
		fx.Rates.TOP = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TOP"].(float64)), 6)
		fx.Rates.TRY = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TRY"].(float64)), 6)
		fx.Rates.TTD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TTD"].(float64)), 6)
		fx.Rates.TWD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TWD"].(float64)), 6)
		fx.Rates.TZS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["TZS"].(float64)), 6)
		fx.Rates.UAH = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["UAH"].(float64)), 6)
		fx.Rates.UGX = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["UGX"].(float64)), 6)
		fx.Rates.USD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["USD"].(float64)), 6)
		fx.Rates.UYU = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["UYU"].(float64)), 6)
		fx.Rates.UZS = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["UZS"].(float64)), 6)
		fx.Rates.VEF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["VEF"].(float64)), 6)
		fx.Rates.VND = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["VND"].(float64)), 6)
		fx.Rates.VUV = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["VUV"].(float64)), 6)
		fx.Rates.WST = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["WST"].(float64)), 6)
		fx.Rates.XAF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XAF"].(float64)), 6)
		fx.Rates.XAG = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XAG"].(float64)), 6)
		fx.Rates.XAU = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XAU"].(float64)), 6)
		fx.Rates.XCD = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XCD"].(float64)), 6)
		fx.Rates.XDR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XDR"].(float64)), 6)
		fx.Rates.XOF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XOF"].(float64)), 6)
		fx.Rates.XPF = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["XPF"].(float64)), 6)
		fx.Rates.YER = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["YER"].(float64)), 6)
		fx.Rates.ZAR = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ZAR"].(float64)), 6)
		fx.Rates.ZMK = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ZMK"].(float64)), 6)
		fx.Rates.ZMW = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ZMW"].(float64)), 6)
		fx.Rates.ZWL = toFixed(rateEUR/(1/result.(map[string]interface{})["rates"].(map[string]interface{})["ZWL"].(float64)), 6)

		// fmt.Printf("base: %v\n", result.(map[string]interface{})["base"].(string))
		// fmt.Printf("rates USD: %v\n", result.(map[string]interface{})["rates"].(map[string]interface{})["USD"])
		// fmt.Printf("rates INR: %v\n", result.(map[string]interface{})["rates"].(map[string]interface{})["INR"])

		// b, _ := json.Marshal(fx)
		// fmt.Println(string(b))

		FIXER_RATES = fx

		sleepSeconds := 60 * 60 * 4
		fmt.Printf("Updated Fixer Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return fx, nil
	}
}