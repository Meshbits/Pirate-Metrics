package main

import (
	"fmt"
	"log"
	"time"
)

var VRSC_BTC_SAFETRADE_RATES ConversionRates

func VrscBtcSafeTradeAPI() {
	if _, ok := MARKETS_AVAILABLE["safetrade"]; ok {
		for _, ma := range MARKETS_AVAILABLE["safetrade"] {
			if ma != "VRSC-BTC" {
				MARKETS_AVAILABLE["safetrade"] = append(MARKETS_AVAILABLE["safetrade"], "VRSC-BTC")
			}
		}
	} else {
		MARKETS_AVAILABLE["safetrade"] = append(MARKETS_AVAILABLE["safetrade"], "VRSC-BTC")
	}
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "https://safe.trade/api/v2/peatio/public/markets/vrscbtc/trades"
		params := ``
		result, err := RPCResultMap(url, params)
		if err != nil {
			log.Printf("Got error fetching SafeTrade VRSC/BTC rates: %v\n", err)
		}
		// fmt.Println(result)
		resultType := fmt.Sprintf("%T", result)
		if resultType == "<nil>" {
			// fmt.Println("type is <nil>")
			// fmt.Println("resultType value:", resultType)
			// fmt.Printf("xTyle type: %T\n", resultType)
		} else {
			vrscBTCPrice := result.([]interface{})[0].(map[string]interface{})["price"].(float64)
			vrscUSDPrice := vrscBTCPrice * BTC_RATES.Rates.USD
			// fmt.Println("==========================")
			// fmt.Printf("VRSC Price (BTC): %.8f\n", vrscBTCPrice)
			// fmt.Printf("VRSC Price (USD): %.6f\n", vrscUSDPrice)
			// fmt.Println("==========================")

			var vrsc ConversionRates
			vrsc.Timestamp = int64(result.([]interface{})[0].(map[string]interface{})["created_at"].(float64))
			vrsc.Base = "VRSC"
			vrsc.Market = "SafeTrade"
			vrsc.Rates.AED = toFixed(BTC_RATES.Rates.AED*vrscBTCPrice, 6)
			vrsc.Rates.AFN = toFixed(BTC_RATES.Rates.AFN*vrscBTCPrice, 6)
			vrsc.Rates.ALL = toFixed(BTC_RATES.Rates.ALL*vrscBTCPrice, 6)
			vrsc.Rates.AMD = toFixed(BTC_RATES.Rates.AMD*vrscBTCPrice, 6)
			vrsc.Rates.ANG = toFixed(BTC_RATES.Rates.ANG*vrscBTCPrice, 6)
			vrsc.Rates.AOA = toFixed(BTC_RATES.Rates.AOA*vrscBTCPrice, 6)
			vrsc.Rates.ARS = toFixed(BTC_RATES.Rates.ARS*vrscBTCPrice, 6)
			vrsc.Rates.AUD = toFixed(BTC_RATES.Rates.AUD*vrscBTCPrice, 6)
			vrsc.Rates.AWG = toFixed(BTC_RATES.Rates.AWG*vrscBTCPrice, 6)
			vrsc.Rates.AZN = toFixed(BTC_RATES.Rates.AZN*vrscBTCPrice, 6)
			vrsc.Rates.BAM = toFixed(BTC_RATES.Rates.BAM*vrscBTCPrice, 6)
			vrsc.Rates.BBD = toFixed(BTC_RATES.Rates.BBD*vrscBTCPrice, 6)
			vrsc.Rates.BDT = toFixed(BTC_RATES.Rates.BDT*vrscBTCPrice, 6)
			vrsc.Rates.BGN = toFixed(BTC_RATES.Rates.BGN*vrscBTCPrice, 6)
			vrsc.Rates.BHD = toFixed(BTC_RATES.Rates.BHD*vrscBTCPrice, 6)
			vrsc.Rates.BIF = toFixed(BTC_RATES.Rates.BIF*vrscBTCPrice, 6)
			vrsc.Rates.BMD = toFixed(vrscUSDPrice, 6)
			vrsc.Rates.BND = toFixed(BTC_RATES.Rates.BND*vrscBTCPrice, 6)
			vrsc.Rates.BOB = toFixed(BTC_RATES.Rates.BOB*vrscBTCPrice, 6)
			vrsc.Rates.BRL = toFixed(BTC_RATES.Rates.BRL*vrscBTCPrice, 6)
			vrsc.Rates.BSD = toFixed(BTC_RATES.Rates.BSD*vrscBTCPrice, 6)
			vrsc.Rates.BTC = toFixed(vrscBTCPrice, 8)
			vrsc.Rates.BTN = toFixed(BTC_RATES.Rates.BTN*vrscBTCPrice, 6)
			vrsc.Rates.BWP = toFixed(BTC_RATES.Rates.BWP*vrscBTCPrice, 6)
			vrsc.Rates.BYN = toFixed(BTC_RATES.Rates.BYN*vrscBTCPrice, 6)
			vrsc.Rates.BYR = toFixed(BTC_RATES.Rates.BYR*vrscBTCPrice, 6)
			vrsc.Rates.BZD = toFixed(BTC_RATES.Rates.BZD*vrscBTCPrice, 6)
			vrsc.Rates.CAD = toFixed(BTC_RATES.Rates.CAD*vrscBTCPrice, 6)
			vrsc.Rates.CDF = toFixed(BTC_RATES.Rates.CDF*vrscBTCPrice, 6)
			vrsc.Rates.CHF = toFixed(BTC_RATES.Rates.CHF*vrscBTCPrice, 6)
			vrsc.Rates.CLF = toFixed(BTC_RATES.Rates.CLF*vrscBTCPrice, 6)
			vrsc.Rates.CLP = toFixed(BTC_RATES.Rates.CLP*vrscBTCPrice, 6)
			vrsc.Rates.CNY = toFixed(BTC_RATES.Rates.CNY*vrscBTCPrice, 6)
			vrsc.Rates.COP = toFixed(BTC_RATES.Rates.COP*vrscBTCPrice, 6)
			vrsc.Rates.CRC = toFixed(BTC_RATES.Rates.CRC*vrscBTCPrice, 6)
			vrsc.Rates.CUC = toFixed(vrscUSDPrice, 6)
			vrsc.Rates.CUP = toFixed(BTC_RATES.Rates.CUP*vrscBTCPrice, 6)
			vrsc.Rates.CVE = toFixed(BTC_RATES.Rates.CVE*vrscBTCPrice, 6)
			vrsc.Rates.CZK = toFixed(BTC_RATES.Rates.CZK*vrscBTCPrice, 6)
			vrsc.Rates.DJF = toFixed(BTC_RATES.Rates.DJF*vrscBTCPrice, 6)
			vrsc.Rates.DKK = toFixed(BTC_RATES.Rates.DKK*vrscBTCPrice, 6)
			vrsc.Rates.DOP = toFixed(BTC_RATES.Rates.DOP*vrscBTCPrice, 6)
			vrsc.Rates.DZD = toFixed(BTC_RATES.Rates.DZD*vrscBTCPrice, 6)
			vrsc.Rates.EGP = toFixed(BTC_RATES.Rates.EGP*vrscBTCPrice, 6)
			vrsc.Rates.ERN = toFixed(BTC_RATES.Rates.ERN*vrscBTCPrice, 6)
			vrsc.Rates.ETB = toFixed(BTC_RATES.Rates.ETB*vrscBTCPrice, 6)
			vrsc.Rates.EUR = toFixed(BTC_RATES.Rates.EUR*vrscBTCPrice, 6)
			vrsc.Rates.FJD = toFixed(BTC_RATES.Rates.FJD*vrscBTCPrice, 6)
			vrsc.Rates.FKP = toFixed(BTC_RATES.Rates.FKP*vrscBTCPrice, 6)
			vrsc.Rates.GBP = toFixed(BTC_RATES.Rates.GBP*vrscBTCPrice, 6)
			vrsc.Rates.GEL = toFixed(BTC_RATES.Rates.GEL*vrscBTCPrice, 6)
			vrsc.Rates.GGP = toFixed(BTC_RATES.Rates.GGP*vrscBTCPrice, 6)
			vrsc.Rates.GHS = toFixed(BTC_RATES.Rates.GHS*vrscBTCPrice, 6)
			vrsc.Rates.GIP = toFixed(BTC_RATES.Rates.GIP*vrscBTCPrice, 6)
			vrsc.Rates.GMD = toFixed(BTC_RATES.Rates.GMD*vrscBTCPrice, 6)
			vrsc.Rates.GNF = toFixed(BTC_RATES.Rates.GNF*vrscBTCPrice, 6)
			vrsc.Rates.GTQ = toFixed(BTC_RATES.Rates.GTQ*vrscBTCPrice, 6)
			vrsc.Rates.GYD = toFixed(BTC_RATES.Rates.GYD*vrscBTCPrice, 6)
			vrsc.Rates.HKD = toFixed(BTC_RATES.Rates.HKD*vrscBTCPrice, 6)
			vrsc.Rates.HNL = toFixed(BTC_RATES.Rates.HNL*vrscBTCPrice, 6)
			vrsc.Rates.HRK = toFixed(BTC_RATES.Rates.HRK*vrscBTCPrice, 6)
			vrsc.Rates.HTG = toFixed(BTC_RATES.Rates.HTG*vrscBTCPrice, 6)
			vrsc.Rates.HUF = toFixed(BTC_RATES.Rates.HUF*vrscBTCPrice, 6)
			vrsc.Rates.IDR = toFixed(BTC_RATES.Rates.IDR*vrscBTCPrice, 6)
			vrsc.Rates.ILS = toFixed(BTC_RATES.Rates.ILS*vrscBTCPrice, 6)
			vrsc.Rates.IMP = toFixed(BTC_RATES.Rates.IMP*vrscBTCPrice, 6)
			vrsc.Rates.INR = toFixed(BTC_RATES.Rates.INR*vrscBTCPrice, 6)
			vrsc.Rates.IQD = toFixed(BTC_RATES.Rates.IQD*vrscBTCPrice, 6)
			vrsc.Rates.IRR = toFixed(BTC_RATES.Rates.IRR*vrscBTCPrice, 6)
			vrsc.Rates.ISK = toFixed(BTC_RATES.Rates.ISK*vrscBTCPrice, 6)
			vrsc.Rates.JEP = toFixed(BTC_RATES.Rates.JEP*vrscBTCPrice, 6)
			vrsc.Rates.JMD = toFixed(BTC_RATES.Rates.JMD*vrscBTCPrice, 6)
			vrsc.Rates.JOD = toFixed(BTC_RATES.Rates.JOD*vrscBTCPrice, 6)
			vrsc.Rates.JPY = toFixed(BTC_RATES.Rates.JPY*vrscBTCPrice, 6)
			vrsc.Rates.KES = toFixed(BTC_RATES.Rates.KES*vrscBTCPrice, 6)
			vrsc.Rates.KGS = toFixed(BTC_RATES.Rates.KGS*vrscBTCPrice, 6)
			vrsc.Rates.KHR = toFixed(BTC_RATES.Rates.KHR*vrscBTCPrice, 6)
			vrsc.Rates.KMF = toFixed(BTC_RATES.Rates.KMF*vrscBTCPrice, 6)
			vrsc.Rates.KPW = toFixed(BTC_RATES.Rates.KPW*vrscBTCPrice, 6)
			vrsc.Rates.KRW = toFixed(BTC_RATES.Rates.KRW*vrscBTCPrice, 6)
			vrsc.Rates.KWD = toFixed(BTC_RATES.Rates.KWD*vrscBTCPrice, 6)
			vrsc.Rates.KYD = toFixed(BTC_RATES.Rates.KYD*vrscBTCPrice, 6)
			vrsc.Rates.KZT = toFixed(BTC_RATES.Rates.KZT*vrscBTCPrice, 6)
			vrsc.Rates.LAK = toFixed(BTC_RATES.Rates.LAK*vrscBTCPrice, 6)
			vrsc.Rates.LBP = toFixed(BTC_RATES.Rates.LBP*vrscBTCPrice, 6)
			vrsc.Rates.LKR = toFixed(BTC_RATES.Rates.LKR*vrscBTCPrice, 6)
			vrsc.Rates.LRD = toFixed(BTC_RATES.Rates.LRD*vrscBTCPrice, 6)
			vrsc.Rates.LSL = toFixed(BTC_RATES.Rates.LSL*vrscBTCPrice, 6)
			vrsc.Rates.LTL = toFixed(BTC_RATES.Rates.LTL*vrscBTCPrice, 6)
			vrsc.Rates.LVL = toFixed(BTC_RATES.Rates.LVL*vrscBTCPrice, 6)
			vrsc.Rates.LYD = toFixed(BTC_RATES.Rates.LYD*vrscBTCPrice, 6)
			vrsc.Rates.MAD = toFixed(BTC_RATES.Rates.MAD*vrscBTCPrice, 6)
			vrsc.Rates.MDL = toFixed(BTC_RATES.Rates.MDL*vrscBTCPrice, 6)
			vrsc.Rates.MGA = toFixed(BTC_RATES.Rates.MGA*vrscBTCPrice, 6)
			vrsc.Rates.MKD = toFixed(BTC_RATES.Rates.MKD*vrscBTCPrice, 6)
			vrsc.Rates.MMK = toFixed(BTC_RATES.Rates.MMK*vrscBTCPrice, 6)
			vrsc.Rates.MNT = toFixed(BTC_RATES.Rates.MNT*vrscBTCPrice, 6)
			vrsc.Rates.MOP = toFixed(BTC_RATES.Rates.MOP*vrscBTCPrice, 6)
			vrsc.Rates.MRO = toFixed(BTC_RATES.Rates.MRO*vrscBTCPrice, 6)
			vrsc.Rates.MUR = toFixed(BTC_RATES.Rates.MUR*vrscBTCPrice, 6)
			vrsc.Rates.MVR = toFixed(BTC_RATES.Rates.MVR*vrscBTCPrice, 6)
			vrsc.Rates.MWK = toFixed(BTC_RATES.Rates.MWK*vrscBTCPrice, 6)
			vrsc.Rates.MXN = toFixed(BTC_RATES.Rates.MXN*vrscBTCPrice, 6)
			vrsc.Rates.MYR = toFixed(BTC_RATES.Rates.MYR*vrscBTCPrice, 6)
			vrsc.Rates.MZN = toFixed(BTC_RATES.Rates.MZN*vrscBTCPrice, 6)
			vrsc.Rates.NAD = toFixed(BTC_RATES.Rates.NAD*vrscBTCPrice, 6)
			vrsc.Rates.NGN = toFixed(BTC_RATES.Rates.NGN*vrscBTCPrice, 6)
			vrsc.Rates.NIO = toFixed(BTC_RATES.Rates.NIO*vrscBTCPrice, 6)
			vrsc.Rates.NOK = toFixed(BTC_RATES.Rates.NOK*vrscBTCPrice, 6)
			vrsc.Rates.NPR = toFixed(BTC_RATES.Rates.NPR*vrscBTCPrice, 6)
			vrsc.Rates.NZD = toFixed(BTC_RATES.Rates.NZD*vrscBTCPrice, 6)
			vrsc.Rates.OMR = toFixed(BTC_RATES.Rates.OMR*vrscBTCPrice, 6)
			vrsc.Rates.PAB = toFixed(BTC_RATES.Rates.PAB*vrscBTCPrice, 6)
			vrsc.Rates.PEN = toFixed(BTC_RATES.Rates.PEN*vrscBTCPrice, 6)
			vrsc.Rates.PGK = toFixed(BTC_RATES.Rates.PGK*vrscBTCPrice, 6)
			vrsc.Rates.PHP = toFixed(BTC_RATES.Rates.PHP*vrscBTCPrice, 6)
			vrsc.Rates.PKR = toFixed(BTC_RATES.Rates.PKR*vrscBTCPrice, 6)
			vrsc.Rates.PLN = toFixed(BTC_RATES.Rates.PLN*vrscBTCPrice, 6)
			vrsc.Rates.PYG = toFixed(BTC_RATES.Rates.PYG*vrscBTCPrice, 6)
			vrsc.Rates.QAR = toFixed(BTC_RATES.Rates.QAR*vrscBTCPrice, 6)
			vrsc.Rates.RON = toFixed(BTC_RATES.Rates.RON*vrscBTCPrice, 6)
			vrsc.Rates.RSD = toFixed(BTC_RATES.Rates.RSD*vrscBTCPrice, 6)
			vrsc.Rates.RUB = toFixed(BTC_RATES.Rates.RUB*vrscBTCPrice, 6)
			vrsc.Rates.RWF = toFixed(BTC_RATES.Rates.RWF*vrscBTCPrice, 6)
			vrsc.Rates.SAR = toFixed(BTC_RATES.Rates.SAR*vrscBTCPrice, 6)
			vrsc.Rates.SBD = toFixed(BTC_RATES.Rates.SBD*vrscBTCPrice, 6)
			vrsc.Rates.SCR = toFixed(BTC_RATES.Rates.SCR*vrscBTCPrice, 6)
			vrsc.Rates.SDG = toFixed(BTC_RATES.Rates.SDG*vrscBTCPrice, 6)
			vrsc.Rates.SEK = toFixed(BTC_RATES.Rates.SEK*vrscBTCPrice, 6)
			vrsc.Rates.SGD = toFixed(BTC_RATES.Rates.SGD*vrscBTCPrice, 6)
			vrsc.Rates.SHP = toFixed(BTC_RATES.Rates.SHP*vrscBTCPrice, 6)
			vrsc.Rates.SLL = toFixed(BTC_RATES.Rates.SLL*vrscBTCPrice, 6)
			vrsc.Rates.SOS = toFixed(BTC_RATES.Rates.SOS*vrscBTCPrice, 6)
			vrsc.Rates.SRD = toFixed(BTC_RATES.Rates.SRD*vrscBTCPrice, 6)
			vrsc.Rates.STD = toFixed(BTC_RATES.Rates.STD*vrscBTCPrice, 6)
			vrsc.Rates.SVC = toFixed(BTC_RATES.Rates.SVC*vrscBTCPrice, 6)
			vrsc.Rates.SYP = toFixed(BTC_RATES.Rates.SYP*vrscBTCPrice, 6)
			vrsc.Rates.SZL = toFixed(BTC_RATES.Rates.SZL*vrscBTCPrice, 6)
			vrsc.Rates.THB = toFixed(BTC_RATES.Rates.THB*vrscBTCPrice, 6)
			vrsc.Rates.TJS = toFixed(BTC_RATES.Rates.TJS*vrscBTCPrice, 6)
			vrsc.Rates.TMT = toFixed(BTC_RATES.Rates.TMT*vrscBTCPrice, 6)
			vrsc.Rates.TND = toFixed(BTC_RATES.Rates.TND*vrscBTCPrice, 6)
			vrsc.Rates.TOP = toFixed(BTC_RATES.Rates.TOP*vrscBTCPrice, 6)
			vrsc.Rates.TRY = toFixed(BTC_RATES.Rates.TRY*vrscBTCPrice, 6)
			vrsc.Rates.TTD = toFixed(BTC_RATES.Rates.TTD*vrscBTCPrice, 6)
			vrsc.Rates.TWD = toFixed(BTC_RATES.Rates.TWD*vrscBTCPrice, 6)
			vrsc.Rates.TZS = toFixed(BTC_RATES.Rates.TZS*vrscBTCPrice, 6)
			vrsc.Rates.UAH = toFixed(BTC_RATES.Rates.UAH*vrscBTCPrice, 6)
			vrsc.Rates.UGX = toFixed(BTC_RATES.Rates.UGX*vrscBTCPrice, 6)
			vrsc.Rates.USD = toFixed(vrscUSDPrice, 6)
			vrsc.Rates.UYU = toFixed(BTC_RATES.Rates.UYU*vrscBTCPrice, 6)
			vrsc.Rates.UZS = toFixed(BTC_RATES.Rates.UZS*vrscBTCPrice, 6)
			vrsc.Rates.VEF = toFixed(BTC_RATES.Rates.VEF*vrscBTCPrice, 6)
			vrsc.Rates.VND = toFixed(BTC_RATES.Rates.VND*vrscBTCPrice, 6)
			vrsc.Rates.VUV = toFixed(BTC_RATES.Rates.VUV*vrscBTCPrice, 6)
			vrsc.Rates.WST = toFixed(BTC_RATES.Rates.WST*vrscBTCPrice, 6)
			vrsc.Rates.XAF = toFixed(BTC_RATES.Rates.XAF*vrscBTCPrice, 6)
			vrsc.Rates.XAG = toFixed(BTC_RATES.Rates.XAG*vrscBTCPrice, 6)
			vrsc.Rates.XAU = toFixed(BTC_RATES.Rates.XAU*vrscBTCPrice, 6)
			vrsc.Rates.XCD = toFixed(BTC_RATES.Rates.XCD*vrscBTCPrice, 6)
			vrsc.Rates.XDR = toFixed(BTC_RATES.Rates.XDR*vrscBTCPrice, 6)
			vrsc.Rates.XOF = toFixed(BTC_RATES.Rates.XOF*vrscBTCPrice, 6)
			vrsc.Rates.XPF = toFixed(BTC_RATES.Rates.XPF*vrscBTCPrice, 6)
			vrsc.Rates.YER = toFixed(BTC_RATES.Rates.YER*vrscBTCPrice, 6)
			vrsc.Rates.ZAR = toFixed(BTC_RATES.Rates.ZAR*vrscBTCPrice, 6)
			vrsc.Rates.ZMK = toFixed(BTC_RATES.Rates.ZMK*vrscBTCPrice, 6)
			vrsc.Rates.ZMW = toFixed(BTC_RATES.Rates.ZMW*vrscBTCPrice, 6)
			vrsc.Rates.ZWL = toFixed(BTC_RATES.Rates.ZWL*vrscBTCPrice, 6)

			VRSC_BTC_SAFETRADE_RATES = vrsc

			// b, _ := json.Marshal(arrr)
			// fmt.Println(string(b))

		}
		sleepSeconds := SAFETRADE_SECONDS
		log.Printf("Updated VRSC Rates from SafeTrade for VRSC/BTC pair. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return
	}
}
