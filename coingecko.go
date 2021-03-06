package main

import (
	"fmt"
	"log"
	"pirate-metrics/utils"
	"sync"
	"time"
)

var BTC_COINGECKO_RATES ConversionRates

func BtcUsdCoinGeckoAPI(wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			utils.Log.Println("Recovered in BtcUsdCoinGeckoAPI", r)
		}
	}()
	updateMarketsAvailable("coingecko", "BTC-USD", wg)
	// Forever loop to keep fetching rates every N seconds
	for {
		url := "https://api.coingecko.com/api/v3/coins/markets?"
		params := `vs_currency=usd&ids=bitcoin&order=market_cap_desc&per_page=100&page=1&sparkline=false&price_change_percentage=1h%2C%2024h%2C%207d%2C%2030d`
		result, err := RPCResultMap(url, params)
		if err != nil {
			log.Printf("Got error fetching Coingecko BTC/USD rates: %v\n", err)
		}
		// fmt.Println(result)

		resultType := fmt.Sprintf("%T", result)
		if resultType != "[]interface {}" {
			fmt.Println("type is not []interface {}")
			fmt.Println("resultType value:", resultType)
			fmt.Printf("xTyle type: %T\n", resultType)
		} else {
			log.Printf("BTC Price (USD): %v\n", result.([]interface{})[0].(map[string]interface{})["current_price"])
			var btc ConversionRates
			btc.Success = true
			btc.Timestamp = time.Now().Unix()
			btc.Base = "USD"
			btc.Market = "CoinGecko"
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

			BTC_COINGECKO_RATES = btc
			if BTC_PRICE_SOURCE == "CoinGecko" {
				BTC_RATES = btc
			}

			// b, _ := json.Marshal(btc)
			// fmt.Println(string(b))
		}

		sleepSeconds := COINGECKO_SECONDS
		log.Printf("Updated Bitcoin Rates. Will update again in %v seconds...\n", sleepSeconds)
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
}
