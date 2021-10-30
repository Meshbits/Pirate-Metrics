package main

import (
	"fmt"
	"strconv"
	"time"
)

var BTC_BINANCE_RATES_USDT ConversionRates

func BtcUsdtBinanceAPI() {
	MARKETS_AVAILABLE["binance"] = append(MARKETS_AVAILABLE["binance"], "BTC-USDT")
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "https://api.binance.com/api/v3/trades?symbol=BTCUSDT"
		params := ``
		result, err := RPCResultMap(url, params)
		if err != nil {
			fmt.Printf("Got error fetching Binance BTC/USDT rates: %v\n", err)
		}
		// fmt.Println(result)
		price, _ := strconv.ParseFloat(result.([]interface{})[0].(map[string]interface{})["price"].(string), 64)
		fmt.Printf("BTC Price (USDT): %v\n", price)
		fmt.Printf("BTC Price Time (USDT): %d\n", int64(result.([]interface{})[0].(map[string]interface{})["time"].(float64)))

		var btc ConversionRates
		btc.Timestamp = time.Now().Unix()
		btc.Base = "USD"
		btc.Market = "Binance"
		btc.Rates.AED = toFixed(FIXER_RATES.Rates.AED*price, 6)
		btc.Rates.AFN = toFixed(FIXER_RATES.Rates.AFN*price, 6)
		btc.Rates.ALL = toFixed(FIXER_RATES.Rates.ALL*price, 6)
		btc.Rates.AMD = toFixed(FIXER_RATES.Rates.AMD*price, 6)
		btc.Rates.ANG = toFixed(FIXER_RATES.Rates.ANG*price, 6)
		btc.Rates.AOA = toFixed(FIXER_RATES.Rates.AOA*price, 6)
		btc.Rates.ARS = toFixed(FIXER_RATES.Rates.ARS*price, 6)
		btc.Rates.AUD = toFixed(FIXER_RATES.Rates.AUD*price, 6)
		btc.Rates.AWG = toFixed(FIXER_RATES.Rates.AWG*price, 6)
		btc.Rates.AZN = toFixed(FIXER_RATES.Rates.AZN*price, 6)
		btc.Rates.BAM = toFixed(FIXER_RATES.Rates.BAM*price, 6)
		btc.Rates.BBD = toFixed(FIXER_RATES.Rates.BBD*price, 6)
		btc.Rates.BDT = toFixed(FIXER_RATES.Rates.BDT*price, 6)
		btc.Rates.BGN = toFixed(FIXER_RATES.Rates.BGN*price, 6)
		btc.Rates.BHD = toFixed(FIXER_RATES.Rates.BHD*price, 6)
		btc.Rates.BIF = toFixed(FIXER_RATES.Rates.BIF*price, 6)
		btc.Rates.BMD = toFixed(FIXER_RATES.Rates.BMD*price, 6)
		btc.Rates.BND = toFixed(FIXER_RATES.Rates.BND*price, 6)
		btc.Rates.BOB = toFixed(FIXER_RATES.Rates.BOB*price, 6)
		btc.Rates.BRL = toFixed(FIXER_RATES.Rates.BRL*price, 6)
		btc.Rates.BSD = toFixed(FIXER_RATES.Rates.BSD*price, 6)
		btc.Rates.BTC = toFixed(FIXER_RATES.Rates.BTC*price, 8)
		btc.Rates.BTN = toFixed(FIXER_RATES.Rates.BTN*price, 6)
		btc.Rates.BWP = toFixed(FIXER_RATES.Rates.BWP*price, 6)
		btc.Rates.BYN = toFixed(FIXER_RATES.Rates.BYN*price, 6)
		btc.Rates.BYR = toFixed(FIXER_RATES.Rates.BYR*price, 6)
		btc.Rates.BZD = toFixed(FIXER_RATES.Rates.BZD*price, 6)
		btc.Rates.CAD = toFixed(FIXER_RATES.Rates.CAD*price, 6)
		btc.Rates.CDF = toFixed(FIXER_RATES.Rates.CDF*price, 6)
		btc.Rates.CHF = toFixed(FIXER_RATES.Rates.CHF*price, 6)
		btc.Rates.CLF = toFixed(FIXER_RATES.Rates.CLF*price, 6)
		btc.Rates.CLP = toFixed(FIXER_RATES.Rates.CLP*price, 6)
		btc.Rates.CNY = toFixed(FIXER_RATES.Rates.CNY*price, 6)
		btc.Rates.COP = toFixed(FIXER_RATES.Rates.COP*price, 6)
		btc.Rates.CRC = toFixed(FIXER_RATES.Rates.CRC*price, 6)
		btc.Rates.CUC = toFixed(FIXER_RATES.Rates.CUC*price, 6)
		btc.Rates.CUP = toFixed(FIXER_RATES.Rates.CUP*price, 6)
		btc.Rates.CVE = toFixed(FIXER_RATES.Rates.CVE*price, 6)
		btc.Rates.CZK = toFixed(FIXER_RATES.Rates.CZK*price, 6)
		btc.Rates.DJF = toFixed(FIXER_RATES.Rates.DJF*price, 6)
		btc.Rates.DKK = toFixed(FIXER_RATES.Rates.DKK*price, 6)
		btc.Rates.DOP = toFixed(FIXER_RATES.Rates.DOP*price, 6)
		btc.Rates.DZD = toFixed(FIXER_RATES.Rates.DZD*price, 6)
		btc.Rates.EGP = toFixed(FIXER_RATES.Rates.EGP*price, 6)
		btc.Rates.ERN = toFixed(FIXER_RATES.Rates.ERN*price, 6)
		btc.Rates.ETB = toFixed(FIXER_RATES.Rates.ETB*price, 6)
		btc.Rates.EUR = toFixed(FIXER_RATES.Rates.EUR*price, 6)
		btc.Rates.FJD = toFixed(FIXER_RATES.Rates.FJD*price, 6)
		btc.Rates.FKP = toFixed(FIXER_RATES.Rates.FKP*price, 6)
		btc.Rates.GBP = toFixed(FIXER_RATES.Rates.GBP*price, 6)
		btc.Rates.GEL = toFixed(FIXER_RATES.Rates.GEL*price, 6)
		btc.Rates.GGP = toFixed(FIXER_RATES.Rates.GGP*price, 6)
		btc.Rates.GHS = toFixed(FIXER_RATES.Rates.GHS*price, 6)
		btc.Rates.GIP = toFixed(FIXER_RATES.Rates.GIP*price, 6)
		btc.Rates.GMD = toFixed(FIXER_RATES.Rates.GMD*price, 6)
		btc.Rates.GNF = toFixed(FIXER_RATES.Rates.GNF*price, 6)
		btc.Rates.GTQ = toFixed(FIXER_RATES.Rates.GTQ*price, 6)
		btc.Rates.GYD = toFixed(FIXER_RATES.Rates.GYD*price, 6)
		btc.Rates.HKD = toFixed(FIXER_RATES.Rates.HKD*price, 6)
		btc.Rates.HNL = toFixed(FIXER_RATES.Rates.HNL*price, 6)
		btc.Rates.HRK = toFixed(FIXER_RATES.Rates.HRK*price, 6)
		btc.Rates.HTG = toFixed(FIXER_RATES.Rates.HTG*price, 6)
		btc.Rates.HUF = toFixed(FIXER_RATES.Rates.HUF*price, 6)
		btc.Rates.IDR = toFixed(FIXER_RATES.Rates.IDR*price, 6)
		btc.Rates.ILS = toFixed(FIXER_RATES.Rates.ILS*price, 6)
		btc.Rates.IMP = toFixed(FIXER_RATES.Rates.IMP*price, 6)
		btc.Rates.INR = toFixed(FIXER_RATES.Rates.INR*price, 6)
		btc.Rates.IQD = toFixed(FIXER_RATES.Rates.IQD*price, 6)
		btc.Rates.IRR = toFixed(FIXER_RATES.Rates.IRR*price, 6)
		btc.Rates.ISK = toFixed(FIXER_RATES.Rates.ISK*price, 6)
		btc.Rates.JEP = toFixed(FIXER_RATES.Rates.JEP*price, 6)
		btc.Rates.JMD = toFixed(FIXER_RATES.Rates.JMD*price, 6)
		btc.Rates.JOD = toFixed(FIXER_RATES.Rates.JOD*price, 6)
		btc.Rates.JPY = toFixed(FIXER_RATES.Rates.JPY*price, 6)
		btc.Rates.KES = toFixed(FIXER_RATES.Rates.KES*price, 6)
		btc.Rates.KGS = toFixed(FIXER_RATES.Rates.KGS*price, 6)
		btc.Rates.KHR = toFixed(FIXER_RATES.Rates.KHR*price, 6)
		btc.Rates.KMF = toFixed(FIXER_RATES.Rates.KMF*price, 6)
		btc.Rates.KPW = toFixed(FIXER_RATES.Rates.KPW*price, 6)
		btc.Rates.KRW = toFixed(FIXER_RATES.Rates.KRW*price, 6)
		btc.Rates.KWD = toFixed(FIXER_RATES.Rates.KWD*price, 6)
		btc.Rates.KYD = toFixed(FIXER_RATES.Rates.KYD*price, 6)
		btc.Rates.KZT = toFixed(FIXER_RATES.Rates.KZT*price, 6)
		btc.Rates.LAK = toFixed(FIXER_RATES.Rates.LAK*price, 6)
		btc.Rates.LBP = toFixed(FIXER_RATES.Rates.LBP*price, 6)
		btc.Rates.LKR = toFixed(FIXER_RATES.Rates.LKR*price, 6)
		btc.Rates.LRD = toFixed(FIXER_RATES.Rates.LRD*price, 6)
		btc.Rates.LSL = toFixed(FIXER_RATES.Rates.LSL*price, 6)
		btc.Rates.LTL = toFixed(FIXER_RATES.Rates.LTL*price, 6)
		btc.Rates.LVL = toFixed(FIXER_RATES.Rates.LVL*price, 6)
		btc.Rates.LYD = toFixed(FIXER_RATES.Rates.LYD*price, 6)
		btc.Rates.MAD = toFixed(FIXER_RATES.Rates.MAD*price, 6)
		btc.Rates.MDL = toFixed(FIXER_RATES.Rates.MDL*price, 6)
		btc.Rates.MGA = toFixed(FIXER_RATES.Rates.MGA*price, 6)
		btc.Rates.MKD = toFixed(FIXER_RATES.Rates.MKD*price, 6)
		btc.Rates.MMK = toFixed(FIXER_RATES.Rates.MMK*price, 6)
		btc.Rates.MNT = toFixed(FIXER_RATES.Rates.MNT*price, 6)
		btc.Rates.MOP = toFixed(FIXER_RATES.Rates.MOP*price, 6)
		btc.Rates.MRO = toFixed(FIXER_RATES.Rates.MRO*price, 6)
		btc.Rates.MUR = toFixed(FIXER_RATES.Rates.MUR*price, 6)
		btc.Rates.MVR = toFixed(FIXER_RATES.Rates.MVR*price, 6)
		btc.Rates.MWK = toFixed(FIXER_RATES.Rates.MWK*price, 6)
		btc.Rates.MXN = toFixed(FIXER_RATES.Rates.MXN*price, 6)
		btc.Rates.MYR = toFixed(FIXER_RATES.Rates.MYR*price, 6)
		btc.Rates.MZN = toFixed(FIXER_RATES.Rates.MZN*price, 6)
		btc.Rates.NAD = toFixed(FIXER_RATES.Rates.NAD*price, 6)
		btc.Rates.NGN = toFixed(FIXER_RATES.Rates.NGN*price, 6)
		btc.Rates.NIO = toFixed(FIXER_RATES.Rates.NIO*price, 6)
		btc.Rates.NOK = toFixed(FIXER_RATES.Rates.NOK*price, 6)
		btc.Rates.NPR = toFixed(FIXER_RATES.Rates.NPR*price, 6)
		btc.Rates.NZD = toFixed(FIXER_RATES.Rates.NZD*price, 6)
		btc.Rates.OMR = toFixed(FIXER_RATES.Rates.OMR*price, 6)
		btc.Rates.PAB = toFixed(FIXER_RATES.Rates.PAB*price, 6)
		btc.Rates.PEN = toFixed(FIXER_RATES.Rates.PEN*price, 6)
		btc.Rates.PGK = toFixed(FIXER_RATES.Rates.PGK*price, 6)
		btc.Rates.PHP = toFixed(FIXER_RATES.Rates.PHP*price, 6)
		btc.Rates.PKR = toFixed(FIXER_RATES.Rates.PKR*price, 6)
		btc.Rates.PLN = toFixed(FIXER_RATES.Rates.PLN*price, 6)
		btc.Rates.PYG = toFixed(FIXER_RATES.Rates.PYG*price, 6)
		btc.Rates.QAR = toFixed(FIXER_RATES.Rates.QAR*price, 6)
		btc.Rates.RON = toFixed(FIXER_RATES.Rates.RON*price, 6)
		btc.Rates.RSD = toFixed(FIXER_RATES.Rates.RSD*price, 6)
		btc.Rates.RUB = toFixed(FIXER_RATES.Rates.RUB*price, 6)
		btc.Rates.RWF = toFixed(FIXER_RATES.Rates.RWF*price, 6)
		btc.Rates.SAR = toFixed(FIXER_RATES.Rates.SAR*price, 6)
		btc.Rates.SBD = toFixed(FIXER_RATES.Rates.SBD*price, 6)
		btc.Rates.SCR = toFixed(FIXER_RATES.Rates.SCR*price, 6)
		btc.Rates.SDG = toFixed(FIXER_RATES.Rates.SDG*price, 6)
		btc.Rates.SEK = toFixed(FIXER_RATES.Rates.SEK*price, 6)
		btc.Rates.SGD = toFixed(FIXER_RATES.Rates.SGD*price, 6)
		btc.Rates.SHP = toFixed(FIXER_RATES.Rates.SHP*price, 6)
		btc.Rates.SLL = toFixed(FIXER_RATES.Rates.SLL*price, 6)
		btc.Rates.SOS = toFixed(FIXER_RATES.Rates.SOS*price, 6)
		btc.Rates.SRD = toFixed(FIXER_RATES.Rates.SRD*price, 6)
		btc.Rates.STD = toFixed(FIXER_RATES.Rates.STD*price, 6)
		btc.Rates.SVC = toFixed(FIXER_RATES.Rates.SVC*price, 6)
		btc.Rates.SYP = toFixed(FIXER_RATES.Rates.SYP*price, 6)
		btc.Rates.SZL = toFixed(FIXER_RATES.Rates.SZL*price, 6)
		btc.Rates.THB = toFixed(FIXER_RATES.Rates.THB*price, 6)
		btc.Rates.TJS = toFixed(FIXER_RATES.Rates.TJS*price, 6)
		btc.Rates.TMT = toFixed(FIXER_RATES.Rates.TMT*price, 6)
		btc.Rates.TND = toFixed(FIXER_RATES.Rates.TND*price, 6)
		btc.Rates.TOP = toFixed(FIXER_RATES.Rates.TOP*price, 6)
		btc.Rates.TRY = toFixed(FIXER_RATES.Rates.TRY*price, 6)
		btc.Rates.TTD = toFixed(FIXER_RATES.Rates.TTD*price, 6)
		btc.Rates.TWD = toFixed(FIXER_RATES.Rates.TWD*price, 6)
		btc.Rates.TZS = toFixed(FIXER_RATES.Rates.TZS*price, 6)
		btc.Rates.UAH = toFixed(FIXER_RATES.Rates.UAH*price, 6)
		btc.Rates.UGX = toFixed(FIXER_RATES.Rates.UGX*price, 6)
		btc.Rates.USD = toFixed(FIXER_RATES.Rates.USD*price, 6)
		btc.Rates.UYU = toFixed(FIXER_RATES.Rates.UYU*price, 6)
		btc.Rates.UZS = toFixed(FIXER_RATES.Rates.UZS*price, 6)
		btc.Rates.VEF = toFixed(FIXER_RATES.Rates.VEF*price, 6)
		btc.Rates.VND = toFixed(FIXER_RATES.Rates.VND*price, 6)
		btc.Rates.VUV = toFixed(FIXER_RATES.Rates.VUV*price, 6)
		btc.Rates.WST = toFixed(FIXER_RATES.Rates.WST*price, 6)
		btc.Rates.XAF = toFixed(FIXER_RATES.Rates.XAF*price, 6)
		btc.Rates.XAG = toFixed(FIXER_RATES.Rates.XAG*price, 6)
		btc.Rates.XAU = toFixed(FIXER_RATES.Rates.XAU*price, 6)
		btc.Rates.XCD = toFixed(FIXER_RATES.Rates.XCD*price, 6)
		btc.Rates.XDR = toFixed(FIXER_RATES.Rates.XDR*price, 6)
		btc.Rates.XOF = toFixed(FIXER_RATES.Rates.XOF*price, 6)
		btc.Rates.XPF = toFixed(FIXER_RATES.Rates.XPF*price, 6)
		btc.Rates.YER = toFixed(FIXER_RATES.Rates.YER*price, 6)
		btc.Rates.ZAR = toFixed(FIXER_RATES.Rates.ZAR*price, 6)
		btc.Rates.ZMK = toFixed(FIXER_RATES.Rates.ZMK*price, 6)
		btc.Rates.ZMW = toFixed(FIXER_RATES.Rates.ZMW*price, 6)
		btc.Rates.ZWL = toFixed(FIXER_RATES.Rates.ZWL*price, 6)

		BTC_BINANCE_RATES_USDT = btc
		if BTC_PRICE_SOURCE == "Binance" {
			BTC_RATES = btc
		}

		// b, _ := json.Marshal(arrr)
		// fmt.Println(string(b))

		sleepSeconds := 60 * 3
		fmt.Printf("Updated BTC Rates from Binance for BTC/USDT pair. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)

		// return
	}
}