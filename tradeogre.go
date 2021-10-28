package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

var ARRR_RATES pirateRates

type pirateRates struct {
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

func TOgreARRRAPI() {
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "https://tradeogre.com/api/v1/ticker/BTC-ARRR"
		params := ``
		result, err := RPCResultMap(url, params)
		if err != nil {
			fmt.Printf("Got error fetching Fixer rates: %v\n", err)
		}
		// fmt.Println(result)
		fmt.Printf("ARRR Price (BTC): %v\n", result.(map[string]interface{})["price"].(string))

		var arrr pirateRates
		arrrBTC, _ := strconv.ParseFloat(result.(map[string]interface{})["price"].(string), 64)
		arrr.Base = "ARRR"
		arrr.Rates.AED = toFixed(arrrBTC*BTC_RATES.Rates.AED, 6)
		arrr.Rates.AFN = toFixed(arrrBTC*BTC_RATES.Rates.AFN, 6)
		arrr.Rates.ALL = toFixed(arrrBTC*BTC_RATES.Rates.ALL, 6)
		arrr.Rates.AMD = toFixed(arrrBTC*BTC_RATES.Rates.AMD, 6)
		arrr.Rates.ANG = toFixed(arrrBTC*BTC_RATES.Rates.ANG, 6)
		arrr.Rates.AOA = toFixed(arrrBTC*BTC_RATES.Rates.AOA, 6)
		arrr.Rates.ARS = toFixed(arrrBTC*BTC_RATES.Rates.ARS, 6)
		arrr.Rates.AUD = toFixed(arrrBTC*BTC_RATES.Rates.AUD, 6)
		arrr.Rates.AWG = toFixed(arrrBTC*BTC_RATES.Rates.AWG, 6)
		arrr.Rates.AZN = toFixed(arrrBTC*BTC_RATES.Rates.AZN, 6)
		arrr.Rates.BAM = toFixed(arrrBTC*BTC_RATES.Rates.BAM, 6)
		arrr.Rates.BBD = toFixed(arrrBTC*BTC_RATES.Rates.BBD, 6)
		arrr.Rates.BDT = toFixed(arrrBTC*BTC_RATES.Rates.BDT, 6)
		arrr.Rates.BGN = toFixed(arrrBTC*BTC_RATES.Rates.BGN, 6)
		arrr.Rates.BHD = toFixed(arrrBTC*BTC_RATES.Rates.BHD, 6)
		arrr.Rates.BIF = toFixed(arrrBTC*BTC_RATES.Rates.BIF, 6)
		arrr.Rates.BMD = toFixed(arrrBTC*BTC_RATES.Rates.BMD, 6)
		arrr.Rates.BND = toFixed(arrrBTC*BTC_RATES.Rates.BND, 6)
		arrr.Rates.BOB = toFixed(arrrBTC*BTC_RATES.Rates.BOB, 6)
		arrr.Rates.BRL = toFixed(arrrBTC*BTC_RATES.Rates.BRL, 6)
		arrr.Rates.BSD = toFixed(arrrBTC*BTC_RATES.Rates.BSD, 6)
		arrr.Rates.BTC = toFixed(arrrBTC*1, 8)
		arrr.Rates.BTN = toFixed(arrrBTC*BTC_RATES.Rates.BTN, 6)
		arrr.Rates.BWP = toFixed(arrrBTC*BTC_RATES.Rates.BWP, 6)
		arrr.Rates.BYN = toFixed(arrrBTC*BTC_RATES.Rates.BYN, 6)
		arrr.Rates.BYR = toFixed(arrrBTC*BTC_RATES.Rates.BYR, 6)
		arrr.Rates.BZD = toFixed(arrrBTC*BTC_RATES.Rates.BZD, 6)
		arrr.Rates.CAD = toFixed(arrrBTC*BTC_RATES.Rates.CAD, 6)
		arrr.Rates.CDF = toFixed(arrrBTC*BTC_RATES.Rates.CDF, 6)
		arrr.Rates.CHF = toFixed(arrrBTC*BTC_RATES.Rates.CHF, 6)
		arrr.Rates.CLF = toFixed(arrrBTC*BTC_RATES.Rates.CLF, 6)
		arrr.Rates.CLP = toFixed(arrrBTC*BTC_RATES.Rates.CLP, 6)
		arrr.Rates.CNY = toFixed(arrrBTC*BTC_RATES.Rates.CNY, 6)
		arrr.Rates.COP = toFixed(arrrBTC*BTC_RATES.Rates.COP, 6)
		arrr.Rates.CRC = toFixed(arrrBTC*BTC_RATES.Rates.CRC, 6)
		arrr.Rates.CUC = toFixed(arrrBTC*BTC_RATES.Rates.CUC, 6)
		arrr.Rates.CUP = toFixed(arrrBTC*BTC_RATES.Rates.CUP, 6)
		arrr.Rates.CVE = toFixed(arrrBTC*BTC_RATES.Rates.CVE, 6)
		arrr.Rates.CZK = toFixed(arrrBTC*BTC_RATES.Rates.CZK, 6)
		arrr.Rates.DJF = toFixed(arrrBTC*BTC_RATES.Rates.DJF, 6)
		arrr.Rates.DKK = toFixed(arrrBTC*BTC_RATES.Rates.DKK, 6)
		arrr.Rates.DOP = toFixed(arrrBTC*BTC_RATES.Rates.DOP, 6)
		arrr.Rates.DZD = toFixed(arrrBTC*BTC_RATES.Rates.DZD, 6)
		arrr.Rates.EGP = toFixed(arrrBTC*BTC_RATES.Rates.EGP, 6)
		arrr.Rates.ERN = toFixed(arrrBTC*BTC_RATES.Rates.ERN, 6)
		arrr.Rates.ETB = toFixed(arrrBTC*BTC_RATES.Rates.ETB, 6)
		arrr.Rates.EUR = toFixed(arrrBTC*BTC_RATES.Rates.EUR, 6)
		arrr.Rates.FJD = toFixed(arrrBTC*BTC_RATES.Rates.FJD, 6)
		arrr.Rates.FKP = toFixed(arrrBTC*BTC_RATES.Rates.FKP, 6)
		arrr.Rates.GBP = toFixed(arrrBTC*BTC_RATES.Rates.GBP, 6)
		arrr.Rates.GEL = toFixed(arrrBTC*BTC_RATES.Rates.GEL, 6)
		arrr.Rates.GGP = toFixed(arrrBTC*BTC_RATES.Rates.GGP, 6)
		arrr.Rates.GHS = toFixed(arrrBTC*BTC_RATES.Rates.GHS, 6)
		arrr.Rates.GIP = toFixed(arrrBTC*BTC_RATES.Rates.GIP, 6)
		arrr.Rates.GMD = toFixed(arrrBTC*BTC_RATES.Rates.GMD, 6)
		arrr.Rates.GNF = toFixed(arrrBTC*BTC_RATES.Rates.GNF, 6)
		arrr.Rates.GTQ = toFixed(arrrBTC*BTC_RATES.Rates.GTQ, 6)
		arrr.Rates.GYD = toFixed(arrrBTC*BTC_RATES.Rates.GYD, 6)
		arrr.Rates.HKD = toFixed(arrrBTC*BTC_RATES.Rates.HKD, 6)
		arrr.Rates.HNL = toFixed(arrrBTC*BTC_RATES.Rates.HNL, 6)
		arrr.Rates.HRK = toFixed(arrrBTC*BTC_RATES.Rates.HRK, 6)
		arrr.Rates.HTG = toFixed(arrrBTC*BTC_RATES.Rates.HTG, 6)
		arrr.Rates.HUF = toFixed(arrrBTC*BTC_RATES.Rates.HUF, 6)
		arrr.Rates.IDR = toFixed(arrrBTC*BTC_RATES.Rates.IDR, 6)
		arrr.Rates.ILS = toFixed(arrrBTC*BTC_RATES.Rates.ILS, 6)
		arrr.Rates.IMP = toFixed(arrrBTC*BTC_RATES.Rates.IMP, 6)
		arrr.Rates.INR = toFixed(arrrBTC*BTC_RATES.Rates.INR, 6)
		arrr.Rates.IQD = toFixed(arrrBTC*BTC_RATES.Rates.IQD, 6)
		arrr.Rates.IRR = toFixed(arrrBTC*BTC_RATES.Rates.IRR, 6)
		arrr.Rates.ISK = toFixed(arrrBTC*BTC_RATES.Rates.ISK, 6)
		arrr.Rates.JEP = toFixed(arrrBTC*BTC_RATES.Rates.JEP, 6)
		arrr.Rates.JMD = toFixed(arrrBTC*BTC_RATES.Rates.JMD, 6)
		arrr.Rates.JOD = toFixed(arrrBTC*BTC_RATES.Rates.JOD, 6)
		arrr.Rates.JPY = toFixed(arrrBTC*BTC_RATES.Rates.JPY, 6)
		arrr.Rates.KES = toFixed(arrrBTC*BTC_RATES.Rates.KES, 6)
		arrr.Rates.KGS = toFixed(arrrBTC*BTC_RATES.Rates.KGS, 6)
		arrr.Rates.KHR = toFixed(arrrBTC*BTC_RATES.Rates.KHR, 6)
		arrr.Rates.KMF = toFixed(arrrBTC*BTC_RATES.Rates.KMF, 6)
		arrr.Rates.KPW = toFixed(arrrBTC*BTC_RATES.Rates.KPW, 6)
		arrr.Rates.KRW = toFixed(arrrBTC*BTC_RATES.Rates.KRW, 6)
		arrr.Rates.KWD = toFixed(arrrBTC*BTC_RATES.Rates.KWD, 6)
		arrr.Rates.KYD = toFixed(arrrBTC*BTC_RATES.Rates.KYD, 6)
		arrr.Rates.KZT = toFixed(arrrBTC*BTC_RATES.Rates.KZT, 6)
		arrr.Rates.LAK = toFixed(arrrBTC*BTC_RATES.Rates.LAK, 6)
		arrr.Rates.LBP = toFixed(arrrBTC*BTC_RATES.Rates.LBP, 6)
		arrr.Rates.LKR = toFixed(arrrBTC*BTC_RATES.Rates.LKR, 6)
		arrr.Rates.LRD = toFixed(arrrBTC*BTC_RATES.Rates.LRD, 6)
		arrr.Rates.LSL = toFixed(arrrBTC*BTC_RATES.Rates.LSL, 6)
		arrr.Rates.LTL = toFixed(arrrBTC*BTC_RATES.Rates.LTL, 6)
		arrr.Rates.LVL = toFixed(arrrBTC*BTC_RATES.Rates.LVL, 6)
		arrr.Rates.LYD = toFixed(arrrBTC*BTC_RATES.Rates.LYD, 6)
		arrr.Rates.MAD = toFixed(arrrBTC*BTC_RATES.Rates.MAD, 6)
		arrr.Rates.MDL = toFixed(arrrBTC*BTC_RATES.Rates.MDL, 6)
		arrr.Rates.MGA = toFixed(arrrBTC*BTC_RATES.Rates.MGA, 6)
		arrr.Rates.MKD = toFixed(arrrBTC*BTC_RATES.Rates.MKD, 6)
		arrr.Rates.MMK = toFixed(arrrBTC*BTC_RATES.Rates.MMK, 6)
		arrr.Rates.MNT = toFixed(arrrBTC*BTC_RATES.Rates.MNT, 6)
		arrr.Rates.MOP = toFixed(arrrBTC*BTC_RATES.Rates.MOP, 6)
		arrr.Rates.MRO = toFixed(arrrBTC*BTC_RATES.Rates.MRO, 6)
		arrr.Rates.MUR = toFixed(arrrBTC*BTC_RATES.Rates.MUR, 6)
		arrr.Rates.MVR = toFixed(arrrBTC*BTC_RATES.Rates.MVR, 6)
		arrr.Rates.MWK = toFixed(arrrBTC*BTC_RATES.Rates.MWK, 6)
		arrr.Rates.MXN = toFixed(arrrBTC*BTC_RATES.Rates.MXN, 6)
		arrr.Rates.MYR = toFixed(arrrBTC*BTC_RATES.Rates.MYR, 6)
		arrr.Rates.MZN = toFixed(arrrBTC*BTC_RATES.Rates.MZN, 6)
		arrr.Rates.NAD = toFixed(arrrBTC*BTC_RATES.Rates.NAD, 6)
		arrr.Rates.NGN = toFixed(arrrBTC*BTC_RATES.Rates.NGN, 6)
		arrr.Rates.NIO = toFixed(arrrBTC*BTC_RATES.Rates.NIO, 6)
		arrr.Rates.NOK = toFixed(arrrBTC*BTC_RATES.Rates.NOK, 6)
		arrr.Rates.NPR = toFixed(arrrBTC*BTC_RATES.Rates.NPR, 6)
		arrr.Rates.NZD = toFixed(arrrBTC*BTC_RATES.Rates.NZD, 6)
		arrr.Rates.OMR = toFixed(arrrBTC*BTC_RATES.Rates.OMR, 6)
		arrr.Rates.PAB = toFixed(arrrBTC*BTC_RATES.Rates.PAB, 6)
		arrr.Rates.PEN = toFixed(arrrBTC*BTC_RATES.Rates.PEN, 6)
		arrr.Rates.PGK = toFixed(arrrBTC*BTC_RATES.Rates.PGK, 6)
		arrr.Rates.PHP = toFixed(arrrBTC*BTC_RATES.Rates.PHP, 6)
		arrr.Rates.PKR = toFixed(arrrBTC*BTC_RATES.Rates.PKR, 6)
		arrr.Rates.PLN = toFixed(arrrBTC*BTC_RATES.Rates.PLN, 6)
		arrr.Rates.PYG = toFixed(arrrBTC*BTC_RATES.Rates.PYG, 6)
		arrr.Rates.QAR = toFixed(arrrBTC*BTC_RATES.Rates.QAR, 6)
		arrr.Rates.RON = toFixed(arrrBTC*BTC_RATES.Rates.RON, 6)
		arrr.Rates.RSD = toFixed(arrrBTC*BTC_RATES.Rates.RSD, 6)
		arrr.Rates.RUB = toFixed(arrrBTC*BTC_RATES.Rates.RUB, 6)
		arrr.Rates.RWF = toFixed(arrrBTC*BTC_RATES.Rates.RWF, 6)
		arrr.Rates.SAR = toFixed(arrrBTC*BTC_RATES.Rates.SAR, 6)
		arrr.Rates.SBD = toFixed(arrrBTC*BTC_RATES.Rates.SBD, 6)
		arrr.Rates.SCR = toFixed(arrrBTC*BTC_RATES.Rates.SCR, 6)
		arrr.Rates.SDG = toFixed(arrrBTC*BTC_RATES.Rates.SDG, 6)
		arrr.Rates.SEK = toFixed(arrrBTC*BTC_RATES.Rates.SEK, 6)
		arrr.Rates.SGD = toFixed(arrrBTC*BTC_RATES.Rates.SGD, 6)
		arrr.Rates.SHP = toFixed(arrrBTC*BTC_RATES.Rates.SHP, 6)
		arrr.Rates.SLL = toFixed(arrrBTC*BTC_RATES.Rates.SLL, 6)
		arrr.Rates.SOS = toFixed(arrrBTC*BTC_RATES.Rates.SOS, 6)
		arrr.Rates.SRD = toFixed(arrrBTC*BTC_RATES.Rates.SRD, 6)
		arrr.Rates.STD = toFixed(arrrBTC*BTC_RATES.Rates.STD, 6)
		arrr.Rates.SVC = toFixed(arrrBTC*BTC_RATES.Rates.SVC, 6)
		arrr.Rates.SYP = toFixed(arrrBTC*BTC_RATES.Rates.SYP, 6)
		arrr.Rates.SZL = toFixed(arrrBTC*BTC_RATES.Rates.SZL, 6)
		arrr.Rates.THB = toFixed(arrrBTC*BTC_RATES.Rates.THB, 6)
		arrr.Rates.TJS = toFixed(arrrBTC*BTC_RATES.Rates.TJS, 6)
		arrr.Rates.TMT = toFixed(arrrBTC*BTC_RATES.Rates.TMT, 6)
		arrr.Rates.TND = toFixed(arrrBTC*BTC_RATES.Rates.TND, 6)
		arrr.Rates.TOP = toFixed(arrrBTC*BTC_RATES.Rates.TOP, 6)
		arrr.Rates.TRY = toFixed(arrrBTC*BTC_RATES.Rates.TRY, 6)
		arrr.Rates.TTD = toFixed(arrrBTC*BTC_RATES.Rates.TTD, 6)
		arrr.Rates.TWD = toFixed(arrrBTC*BTC_RATES.Rates.TWD, 6)
		arrr.Rates.TZS = toFixed(arrrBTC*BTC_RATES.Rates.TZS, 6)
		arrr.Rates.UAH = toFixed(arrrBTC*BTC_RATES.Rates.UAH, 6)
		arrr.Rates.UGX = toFixed(arrrBTC*BTC_RATES.Rates.UGX, 6)
		arrr.Rates.USD = toFixed(arrrBTC*BTC_RATES.Rates.USD, 6)
		arrr.Rates.UYU = toFixed(arrrBTC*BTC_RATES.Rates.UYU, 6)
		arrr.Rates.UZS = toFixed(arrrBTC*BTC_RATES.Rates.UZS, 6)
		arrr.Rates.VEF = toFixed(arrrBTC*BTC_RATES.Rates.VEF, 6)
		arrr.Rates.VND = toFixed(arrrBTC*BTC_RATES.Rates.VND, 6)
		arrr.Rates.VUV = toFixed(arrrBTC*BTC_RATES.Rates.VUV, 6)
		arrr.Rates.WST = toFixed(arrrBTC*BTC_RATES.Rates.WST, 6)
		arrr.Rates.XAF = toFixed(arrrBTC*BTC_RATES.Rates.XAF, 6)
		arrr.Rates.XAG = toFixed(arrrBTC*BTC_RATES.Rates.XAG, 6)
		arrr.Rates.XAU = toFixed(arrrBTC*BTC_RATES.Rates.XAU, 6)
		arrr.Rates.XCD = toFixed(arrrBTC*BTC_RATES.Rates.XCD, 6)
		arrr.Rates.XDR = toFixed(arrrBTC*BTC_RATES.Rates.XDR, 6)
		arrr.Rates.XOF = toFixed(arrrBTC*BTC_RATES.Rates.XOF, 6)
		arrr.Rates.XPF = toFixed(arrrBTC*BTC_RATES.Rates.XPF, 6)
		arrr.Rates.YER = toFixed(arrrBTC*BTC_RATES.Rates.YER, 6)
		arrr.Rates.ZAR = toFixed(arrrBTC*BTC_RATES.Rates.ZAR, 6)
		arrr.Rates.ZMK = toFixed(arrrBTC*BTC_RATES.Rates.ZMK, 6)
		arrr.Rates.ZMW = toFixed(arrrBTC*BTC_RATES.Rates.ZMW, 6)
		arrr.Rates.ZWL = toFixed(arrrBTC*BTC_RATES.Rates.ZWL, 6)

		ARRR_RATES = arrr

		b, _ := json.Marshal(arrr)
		fmt.Println(string(b))

		sleepSeconds := 30
		fmt.Printf("Updated Pirate Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return
	}
}
