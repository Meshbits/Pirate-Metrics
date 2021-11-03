## Pirate Metrics
 
:warning: This code is a work in progress and is subject to change.

If you have any suggestions, requests, improvement ideas, please open an issue on github repo. 

## Objective

To have an independent price aggregator API which doesn't depend on CoinMarketCap or CoinGecko type services, and instead gets the prices directly from exchanges to process and render for the application use.

## Currently supported markets and tickers

- Fixer
    - USD
- Binance
    - BTC/USDT
    - KMD/BTC
- KuCoin
    - ARRR/BTC
    - ARRR/USDT
- CoinGecko
    - BTC/USD
- SafeTrade
    - VRSC/BTC
- TradeOgre
    - ARRR/BTC

## Build Requirements

- [Go v1.17+](https://golang.org/doc/install)
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git)
- [Fixer API](https://fixer.io/)

### Install Go on your system

On Ubuntu can follow this guide: https://github.com/golang/go/wiki/Ubuntu

```shell
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go
```

On Mac can install Go using `brew`.

```shell
brew update
brew install go
```

## Installing Pirate Metrics

In Linux/Mac you must have a `go` directory under your `$HOME` directory in your OS.
If you don't find this directory then create the following:

```shell
mkdir -p $HOME/go/{bin,src,pkg}
```

#### Method 1

You can install it with `go` command and after that you can expect the `pirate-metrics` binary under `$HOME/go/bin`.

```
go get -u github.com/Meshbits/Pirate-Metrics
```

#### Method 2

You can make a clone of this repository and then build the executable binary.

```
git clone https://github.com/Meshbits/Pirate-Metrics.git
cd Pirate-Metrics
go build
```

## Get Fixer.io API key

Fixer.io provides foreign exchange for more than 160+ curencies of world. And with the free sign-up you can use maximum 250 API calls per month.
Since foreign exchange rates don't fluctuate too wildely we can utilise 250 API calls approximatly every 4 hours and should be fine keeping up to date with the national currencies conversion rates.

Get a free API by signing up here: https://fixer.io/signup/free

Even in case you feel 250 is insufficient, you can sign up for a [Basic plan](https://fixer.io/signup/basic) for $10 per month and get 10,000 API calls limit.


## API Documentation

### Using Help command

```shell
./pirate-metrics -h 

or 

./pirate-metrics -help
```

Help command will output:

```shell
Usage of ./pirate-metrics:
  -btcPriceSource string
    	You can specify any single string for this option: Binance, CoinGecko (default "Binance")
  -fixerAPIToken string
    	Get the API token from https://fixer.io. API Token from a free sign-up is sufficient to use with this application.
```

### Starting Pirate Metrics Service

Assuming you have Fixer.io API use the following command to start Pirate-Metrics Service.

Example your API key is: `63b7cc4e0f4aea391a65906a67508dda7fbf80`

```shell
./pirate-metrics -fixerAPIToken 63b7cc4e0f4aea391a65906a67508dda7fbf80
```

This will output something like this:

```shell
fixerAPIToken: 63b7cc4e0f4aea391a65906a67508dda7fbf80
btcPriceSource: Binance
2021/11/03 20:35:45 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
2021/11/03 20:35:45 BTC Price Source is: Binance
2021/11/03 20:35:45 BTC USD: 0.000000
2021/11/03 20:35:45 BTC NZD: 0.000000
2021/11/03 20:35:45 TradeOgre-ARRR/BTC: ARRR (BTC): 0.00000000
2021/11/03 20:35:45 TradeOgre-ARRR/BTC: ARRR (USD): 0.000000
2021/11/03 20:35:45 KuCoin-ARRR/BTC: ARRR (BTC): 0.00000000
2021/11/03 20:35:45 KuCoin-ARRR/BTC: ARRR (USD): 0.000000
2021/11/03 20:35:45 KuCoin-ARRR/USDT: ARRR (BTC): 0.00000000
2021/11/03 20:35:45 KuCoin-ARRR/USDT: ARRR (USD): 0.000000
2021/11/03 20:35:45 Binance-KMD/BTC: KMD (BTC) 0.00000000
2021/11/03 20:35:45 Binance-KMD/BTC: KMD (USD) 0.000000
2021/11/03 20:35:45 SafeTrade-VRSC/BTC: VRSC (BTC) 0.00000000
2021/11/03 20:35:45 SafeTrade-VRSC/BTC: VRSC (USD) 0.000000
2021/11/03 20:35:45 ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
2021/11/03 20:35:45 Will Display Rates every other 30 seconds...

2021/11/03 20:35:45 ARRR Price (USDT): 1.9568
2021/11/03 20:35:45 Updated Pirate Rates from KuCoin for ARRR/USDT pair. Will update again in 120 seconds...
2021/11/03 20:35:45 Updated KMD Rates from Binance for KMD/BTC pair. Will update again in 120 seconds...
2021/11/03 20:35:45 BTC Price (USDT): 62996.49
2021/11/03 20:35:45 Updated BTC Rates from Binance for BTC/USDT pair. Will update again in 120 seconds...
2021/11/03 20:35:45 ARRR Price (BTC): 0.00003109
2021/11/03 20:35:45 Updated Pirate Rates from KuCoin for ARRR/BTC pair. Will update again in 120 seconds...
2021/11/03 20:35:46 Updated Fixer Rates. Will update again in 14400 seconds...
2021/11/03 20:35:46 ARRR Price (BTC): 0.00003088
2021/11/03 20:35:46 Updated Pirate Rates from TradeOgre. Will update again in 60 seconds...
2021/11/03 20:35:46 Updated VRSC Rates from SafeTrade for VRSC/BTC pair. Will update again in 60 seconds...
```

The console log will keep printing output.

It will take aproximately 5-10 minutes to display all sample prices in console prints to display accurate prices.

And if you see some price output like the following, that's temporary:

```
2021/11/03 20:37:45 KuCoin-ARRR/USDT: ARRR (BTC): 92233720368.54776001
```

It just means that it's missing some USD prices data and as soon as it's calculated it will start refelecting correct price output.

Once the service is active it is accessible on port 8000.

URL: http://localhost:8000 or http://127.0.0.1:8000


### Stopping Pirate Metrics serivce

To stop the service just press `CTRL+C` to exit the console ouptut log.


### Switching between Binance and CoinGecko for BTC/USD price

At the begining I used CoinGecko for BTC/USD price. But thought I should try getting the price directly from an exchange cutting CoinGecko service dependency.
So, it is possible to switch between either Binance or CoinGecko for getting BTC/USD price data.

By default Price-Metrics code selects `Binanac` as BTC/USD price data, but you can use the following commands to switch:

```shell
# To switch to CoinGecko
./pirate-metrics -fixerAPIToken 63b7cc4e0f4aea391a65906a67508dda7fbf80 -btcPriceSource coingecko
```

```shell
# To switch to Binance
./pirate-metrics -fixerAPIToken 63b7cc4e0f4aea391a65906a67508dda7fbf80 -btcPriceSource binance
```

### Get list of active Markets and Tickers

I have not yet made it optional to enable/disable specific Market or Tickers, but probably will do in the future.


```shell
curl -s http://localhost:8000/v1/prices/markets
```

Output format:

| Key  | Data Type |
| ------------- | ------------- |
| Market name  | string  |
| Tickers  | srings array  |


JSON output:

```json
{
  "binance": [
    "KMD-BTC",
    "BTC-USDT"
  ],
  "fixer": [
    "USD"
  ],
  "kucoin": [
    "ARRR-BTC",
    "ARRR-USDT"
  ],
  "safetrade": [
    "VRSC-BTC"
  ],
  "tradeogre": [
    "ARRR-BTC"
  ]
}
```

### Getting data for selected market and ticker

```shell
curl -s http://localhost:8000/v1/prices/tradeogre/arrr-btc
```

Output format:

| Key  | Data Type | Info
| ------------- | ------------- | ------------- |
| Success  | bool  |   |
| Timestamp  | int64  | Unix Timestamp  |
| Base  | string  | Trading pair ticker  |
| Market  | string  | Exchange website/name  |
| Rates  | objects  |   |
| Rates -> key  | string  | National currency ticker  |
| Rates -> value  | float64  | National currency conversion rate  |


JSON output:

```json
{
  "success": true,
  "timestamp": 1635926217,
  "base": "ARRR/BTC",
  "market": "TradeOgre",
  "rates": {
    "AED": 7.194631,
    "AFN": 178.002009,
    "ALL": 207.747929,
    "AMD": 938.087878,
    "ANG": 3.536366,
    "AOA": 1169.332331,
    "ARS": 195.517951,
    "AUD": 2.632231,
    "AWG": 3.526606,
    "AZN": 3.334618,
    "BAM": 3.307488,
    "BBD": 3.931399,
    "BDT": 168.13037,
    "BGN": 3.307146,
    "BHD": 0.738564,
    "BIF": 3910.260759,
    "BMD": 1.958681,
    "BND": 2.644496,
    "BOB": 13.529242,
    "BRL": 11.125834,
    "BSD": 1.962148,
    "BTC": 3.09e-05,
    "BTN": 146.590901,
    "BWP": 22.50222,
    "BYN": 4.820184,
    "BYR": 38390.159883,
    "BZD": 3.934104,
    "CAD": 2.432173,
    "CDF": 3940.867523,
    "CHF": 1.789849,
    "CLF": 0.057738,
    "CLP": 1593.132573,
    "CNY": 12.533803,
    "COP": 7441.031499,
    "CRC": 1252.118832,
    "CUC": 1.958681,
    "CUP": 51.905063,
    "CVE": 186.462594,
    "CZK": 43.244951,
    "DJF": 349.358116,
    "DKK": 12.580709,
    "DOP": 110.558763,
    "DZD": 267.998792,
    "EGP": 30.777555,
    "ERN": 29.382651,
    "ETB": 93.359399,
    "EUR": 1.691037,
    "FJD": 4.075636,
    "FKP": 1.436148,
    "GBP": 1.436987,
    "GEL": 6.189633,
    "GGP": 1.436148,
    "GHS": 11.988633,
    "GIP": 1.436148,
    "GMD": 101.851333,
    "GNF": 18878.427117,
    "GTQ": 15.180978,
    "GYD": 410.753768,
    "HKD": 15.242079,
    "HNL": 47.40576,
    "HRK": 12.721562,
    "HTG": 192.633146,
    "HUF": 607.916657,
    "IDR": 28040.877905,
    "ILS": 6.140917,
    "IMP": 1.436148,
    "INR": 146.016976,
    "IQD": 2863.471626,
    "IRR": 82734.711958,
    "ISK": 254.001855,
    "JEP": 1.436148,
    "JMD": 303.35087,
    "JOD": 1.388787,
    "JPY": 222.955753,
    "KES": 217.995881,
    "KGS": 166.099129,
    "KHR": 7994.701713,
    "KMF": 832.880456,
    "KPW": 1762.812877,
    "KRW": 2314.808257,
    "KWD": 0.590934,
    "KYD": 1.635096,
    "KZT": 840.292598,
    "LAK": 20267.471872,
    "LBP": 2969.154767,
    "LKR": 396.357342,
    "LRD": 291.59864,
    "LSL": 28.165469,
    "LTL": 5.783477,
    "LVL": 1.184787,
    "LYD": 8.94948,
    "MAD": 17.79316,
    "MDL": 34.366767,
    "MGA": 7797.100203,
    "MKD": 104.195267,
    "MMK": 3537.785889,
    "MNT": 5584.053334,
    "MOP": 15.7305,
    "MRO": 699.249005,
    "MUR": 84.262264,
    "MVR": 30.163366,
    "MWK": 1600.293199,
    "MXN": 40.71323,
    "MYR": 8.137354,
    "MZN": 125.022503,
    "NAD": 28.165125,
    "NGN": 806.447716,
    "NIO": 69.323717,
    "NOK": 16.713033,
    "NPR": 234.549694,
    "NZD": 2.748412,
    "OMR": 0.754114,
    "PAB": 1.962174,
    "PEN": 7.833813,
    "PGK": 6.889588,
    "PHP": 99.014287,
    "PKR": 333.725021,
    "PLN": 7.787377,
    "PYG": 13541.918065,
    "QAR": 7.463068,
    "RON": 8.371017,
    "RSD": 198.869811,
    "RUB": 140.453921,
    "RWF": 2001.212032,
    "SAR": 7.347153,
    "SBD": 15.713514,
    "SCR": 26.002121,
    "SDG": 862.797401,
    "SEK": 16.793717,
    "SGD": 2.642888,
    "SHP": 2.697888,
    "SLL": 21320.250028,
    "SOS": 1143.869799,
    "SRD": 42.308501,
    "STD": 40540.75511,
    "SVC": 17.169432,
    "SYP": 2462.008008,
    "SZL": 30.278773,
    "THB": 65.341324,
    "TJS": 22.034767,
    "TMT": 6.855385,
    "TND": 5.562303,
    "TOP": 4.396257,
    "TRY": 18.850161,
    "TTD": 13.304322,
    "TWD": 54.62215,
    "TZS": 4510.843331,
    "UAH": 51.60401,
    "UGX": 6972.329609,
    "USD": 1.958681,
    "UYU": 86.716564,
    "UZS": 21009.293472,
    "VEF": -285002195.938813,
    "VND": 44556.089643,
    "VUV": 219.952961,
    "WST": 5.047508,
    "XAF": 1109.26909,
    "XAG": 0.083626,
    "XAU": 0.001099,
    "XCD": 5.293435,
    "XDR": 1.390736,
    "XOF": 1109.26909,
    "XPF": 202.430259,
    "YER": 490.160601,
    "ZAR": 30.178485,
    "ZMK": 17630.488288,
    "ZMW": 33.974946,
    "ZWL": 630.694685
  }
}
```

#### More Examples:

For Fixer USD API command would be:

```shell
curl -s http://localhost:8000/v1/prices/fixer/usd
```

For VRSC/BTC from SafeTrade:

```shell
curl -s http://localhost:8000/v1/prices/safetrade/vrsc-btc
```

### Query only specific prices

You can use `symbol` API parameter to query just the specific rates from API.

For example, if you'd like to only get exchange rates for USD, EUR, NZD, INR and BTC, you will need to supply the national currency symbols as following (with no spaces)

```shell
curl -s http://localhost:8000/v1/prices/tradeogre/arrr-btc?symbols=USD,EUR,NZD,INR,BTC
```

It will output like this:

```json
{
  "timestamp": 1635926637,
  "base": "ARRR/BTC",
  "market": "TradeOgre",
  "rates": [
    {
      "USD": 1.962988
    },
    {
      "EUR": 1.694755
    },
    {
      "NZD": 2.754454
    },
    {
      "INR": 146.33799
    },
    {
      "BTC": 3.1e-05
    }
  ]
}
```

For single price query like EUR it will be like this:

```shell
curl -s curl -s http://localhost:8000/v1/prices/safetrade/vrsc-btc\?symbols\=EUR
```

```json
{
  "timestamp": 1635926431,
  "base": "VRSC",
  "market": "SafeTrade",
  "rates": [
    {
      "EUR": 0.930603
    }
  ]
}
```