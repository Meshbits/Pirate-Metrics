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
		url := "http://data.fixer.io/api/latest"
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
		fx.Rates.AED = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AED"].(float64))
		fx.Rates.AFN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AFN"].(float64))
		fx.Rates.ALL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ALL"].(float64))
		fx.Rates.AMD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AMD"].(float64))
		fx.Rates.ANG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ANG"].(float64))
		fx.Rates.AOA = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AOA"].(float64))
		fx.Rates.ARS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ARS"].(float64))
		fx.Rates.AUD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AUD"].(float64))
		fx.Rates.AWG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AWG"].(float64))
		fx.Rates.AZN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["AZN"].(float64))
		fx.Rates.BAM = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BAM"].(float64))
		fx.Rates.BBD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BBD"].(float64))
		fx.Rates.BDT = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BDT"].(float64))
		fx.Rates.BGN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BGN"].(float64))
		fx.Rates.BHD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BHD"].(float64))
		fx.Rates.BIF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BIF"].(float64))
		fx.Rates.BMD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BMD"].(float64))
		fx.Rates.BND = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BND"].(float64))
		fx.Rates.BOB = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BOB"].(float64))
		fx.Rates.BRL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BRL"].(float64))
		fx.Rates.BSD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BSD"].(float64))
		fx.Rates.BTC = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BTC"].(float64))
		fx.Rates.BTN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BTN"].(float64))
		fx.Rates.BWP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BWP"].(float64))
		fx.Rates.BYN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BYN"].(float64))
		fx.Rates.BYR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BYR"].(float64))
		fx.Rates.BZD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["BZD"].(float64))
		fx.Rates.CAD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CAD"].(float64))
		fx.Rates.CDF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CDF"].(float64))
		fx.Rates.CHF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CHF"].(float64))
		fx.Rates.CLF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CLF"].(float64))
		fx.Rates.CLP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CLP"].(float64))
		fx.Rates.CNY = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CNY"].(float64))
		fx.Rates.COP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["COP"].(float64))
		fx.Rates.CRC = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CRC"].(float64))
		fx.Rates.CUC = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CUC"].(float64))
		fx.Rates.CUP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CUP"].(float64))
		fx.Rates.CVE = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CVE"].(float64))
		fx.Rates.CZK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["CZK"].(float64))
		fx.Rates.DJF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["DJF"].(float64))
		fx.Rates.DKK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["DKK"].(float64))
		fx.Rates.DOP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["DOP"].(float64))
		fx.Rates.DZD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["DZD"].(float64))
		fx.Rates.EGP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["EGP"].(float64))
		fx.Rates.ERN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ERN"].(float64))
		fx.Rates.ETB = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ETB"].(float64))
		fx.Rates.EUR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["EUR"].(float64))
		fx.Rates.FJD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["FJD"].(float64))
		fx.Rates.FKP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["FKP"].(float64))
		fx.Rates.GBP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GBP"].(float64))
		fx.Rates.GEL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GEL"].(float64))
		fx.Rates.GGP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GGP"].(float64))
		fx.Rates.GHS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GHS"].(float64))
		fx.Rates.GIP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GIP"].(float64))
		fx.Rates.GMD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GMD"].(float64))
		fx.Rates.GNF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GNF"].(float64))
		fx.Rates.GTQ = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GTQ"].(float64))
		fx.Rates.GYD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["GYD"].(float64))
		fx.Rates.HKD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["HKD"].(float64))
		fx.Rates.HNL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["HNL"].(float64))
		fx.Rates.HRK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["HRK"].(float64))
		fx.Rates.HTG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["HTG"].(float64))
		fx.Rates.HUF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["HUF"].(float64))
		fx.Rates.IDR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["IDR"].(float64))
		fx.Rates.ILS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ILS"].(float64))
		fx.Rates.IMP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["IMP"].(float64))
		fx.Rates.INR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["INR"].(float64))
		fx.Rates.IQD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["IQD"].(float64))
		fx.Rates.IRR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["IRR"].(float64))
		fx.Rates.ISK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ISK"].(float64))
		fx.Rates.JEP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["JEP"].(float64))
		fx.Rates.JMD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["JMD"].(float64))
		fx.Rates.JOD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["JOD"].(float64))
		fx.Rates.JPY = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["JPY"].(float64))
		fx.Rates.KES = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KES"].(float64))
		fx.Rates.KGS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KGS"].(float64))
		fx.Rates.KHR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KHR"].(float64))
		fx.Rates.KMF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KMF"].(float64))
		fx.Rates.KPW = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KPW"].(float64))
		fx.Rates.KRW = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KRW"].(float64))
		fx.Rates.KWD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KWD"].(float64))
		fx.Rates.KYD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KYD"].(float64))
		fx.Rates.KZT = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["KZT"].(float64))
		fx.Rates.LAK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LAK"].(float64))
		fx.Rates.LBP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LBP"].(float64))
		fx.Rates.LKR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LKR"].(float64))
		fx.Rates.LRD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LRD"].(float64))
		fx.Rates.LSL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LSL"].(float64))
		fx.Rates.LTL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LTL"].(float64))
		fx.Rates.LVL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LVL"].(float64))
		fx.Rates.LYD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["LYD"].(float64))
		fx.Rates.MAD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MAD"].(float64))
		fx.Rates.MDL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MDL"].(float64))
		fx.Rates.MGA = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MGA"].(float64))
		fx.Rates.MKD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MKD"].(float64))
		fx.Rates.MMK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MMK"].(float64))
		fx.Rates.MNT = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MNT"].(float64))
		fx.Rates.MOP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MOP"].(float64))
		fx.Rates.MRO = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MRO"].(float64))
		fx.Rates.MUR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MUR"].(float64))
		fx.Rates.MVR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MVR"].(float64))
		fx.Rates.MWK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MWK"].(float64))
		fx.Rates.MXN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MXN"].(float64))
		fx.Rates.MYR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MYR"].(float64))
		fx.Rates.MZN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["MZN"].(float64))
		fx.Rates.NAD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NAD"].(float64))
		fx.Rates.NGN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NGN"].(float64))
		fx.Rates.NIO = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NIO"].(float64))
		fx.Rates.NOK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NOK"].(float64))
		fx.Rates.NPR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NPR"].(float64))
		fx.Rates.NZD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["NZD"].(float64))
		fx.Rates.OMR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["OMR"].(float64))
		fx.Rates.PAB = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PAB"].(float64))
		fx.Rates.PEN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PEN"].(float64))
		fx.Rates.PGK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PGK"].(float64))
		fx.Rates.PHP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PHP"].(float64))
		fx.Rates.PKR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PKR"].(float64))
		fx.Rates.PLN = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PLN"].(float64))
		fx.Rates.PYG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["PYG"].(float64))
		fx.Rates.QAR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["QAR"].(float64))
		fx.Rates.RON = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["RON"].(float64))
		fx.Rates.RSD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["RSD"].(float64))
		fx.Rates.RUB = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["RUB"].(float64))
		fx.Rates.RWF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["RWF"].(float64))
		fx.Rates.SAR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SAR"].(float64))
		fx.Rates.SBD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SBD"].(float64))
		fx.Rates.SCR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SCR"].(float64))
		fx.Rates.SDG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SDG"].(float64))
		fx.Rates.SEK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SEK"].(float64))
		fx.Rates.SGD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SGD"].(float64))
		fx.Rates.SHP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SHP"].(float64))
		fx.Rates.SLL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SLL"].(float64))
		fx.Rates.SOS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SOS"].(float64))
		fx.Rates.SRD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SRD"].(float64))
		fx.Rates.STD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["STD"].(float64))
		fx.Rates.SVC = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SVC"].(float64))
		fx.Rates.SYP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SYP"].(float64))
		fx.Rates.SZL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["SZL"].(float64))
		fx.Rates.THB = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["THB"].(float64))
		fx.Rates.TJS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TJS"].(float64))
		fx.Rates.TMT = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TMT"].(float64))
		fx.Rates.TND = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TND"].(float64))
		fx.Rates.TOP = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TOP"].(float64))
		fx.Rates.TRY = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TRY"].(float64))
		fx.Rates.TTD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TTD"].(float64))
		fx.Rates.TWD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TWD"].(float64))
		fx.Rates.TZS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["TZS"].(float64))
		fx.Rates.UAH = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["UAH"].(float64))
		fx.Rates.UGX = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["UGX"].(float64))
		fx.Rates.USD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["USD"].(float64))
		fx.Rates.UYU = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["UYU"].(float64))
		fx.Rates.UZS = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["UZS"].(float64))
		fx.Rates.VEF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["VEF"].(float64))
		fx.Rates.VND = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["VND"].(float64))
		fx.Rates.VUV = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["VUV"].(float64))
		fx.Rates.WST = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["WST"].(float64))
		fx.Rates.XAF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XAF"].(float64))
		fx.Rates.XAG = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XAG"].(float64))
		fx.Rates.XAU = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XAU"].(float64))
		fx.Rates.XCD = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XCD"].(float64))
		fx.Rates.XDR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XDR"].(float64))
		fx.Rates.XOF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XOF"].(float64))
		fx.Rates.XPF = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["XPF"].(float64))
		fx.Rates.YER = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["YER"].(float64))
		fx.Rates.ZAR = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ZAR"].(float64))
		fx.Rates.ZMK = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ZMK"].(float64))
		fx.Rates.ZMW = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ZMW"].(float64))
		fx.Rates.ZWL = rateEUR / (1 / result.(map[string]interface{})["rates"].(map[string]interface{})["ZWL"].(float64))

		// fmt.Printf("base: %v\n", result.(map[string]interface{})["base"].(string))
		// fmt.Printf("rates USD: %v\n", result.(map[string]interface{})["rates"].(map[string]interface{})["USD"])
		// fmt.Printf("rates INR: %v\n", result.(map[string]interface{})["rates"].(map[string]interface{})["INR"])

		// b, _ := json.Marshal(fx)
		// fmt.Println(string(b))

		FIXER_RATES = fx

		sleepSeconds := 60 * 15
		fmt.Printf("Updated Fixer Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return fx, nil
	}
}
