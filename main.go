package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

var mutex = sync.Mutex{}
var priceMap = make(map[string]float64)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func getExchangeInfo() (ExchangeInfo, error) {
	res, err := http.Get("https://api.binance.com/api/v3/exchangeInfo")
	//TODO: Fill request body
	if err != nil {
		return ExchangeInfo{}, err
	}
	out := json.NewDecoder(res.Body)
	var exInfo ExchangeInfo
	err = out.Decode(&exInfo)
	if err != nil {
		return ExchangeInfo{}, err
	}
	return exInfo, err
}

func getCurrPrices(info ExchangeInfo) {
	defer timeTrack(time.Now(), "getCurrPrices")
	var wg sync.WaitGroup
	for _, name := range info.Symbols {
		//fmt.Printf("%s %f\n",name.Symbol, getPrice(name.Symbol))
		wg.Add(1)
		go getPrice(name.Symbol, &wg)
	}

	wg.Wait()
}

func getPrice(coinPair string, wg *sync.WaitGroup) {
	currLink := "https://api.binance.com/api/v3/avgPrice?symbol=" + coinPair
	res, err := http.Get(currLink)
	defer wg.Done()
	if err != nil {
		println(err.Error())
		priceMap[coinPair] = -420
	}
	out := json.NewDecoder(res.Body)
	var coinStats CoinInfo
	err = out.Decode(&coinStats)
	if err != nil {
		println(err.Error())
		priceMap[coinPair] = -420

	}
	mutex.Lock()
	priceMap[coinPair] = coinStats.Price
	mutex.Unlock()
}

type CoinInfo struct {
	Symbol string
	Price  float64 `json:",string"`
}

type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []RateLimit   `json:"rateLimits"`
	ExchangeFilters []interface{} `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

// RateLimit struct
type RateLimit struct {
	RateLimitType string `json:"rateLimitType"`
	Interval      string `json:"interval"`
	Limit         int64  `json:"limit"`
}

// Symbol market symbol
type Symbol struct {
	Symbol                 string                   `json:"symbol"`
	Status                 string                   `json:"status"`
	BaseAsset              string                   `json:"baseAsset"`
	BaseAssetPrecision     int                      `json:"baseAssetPrecision"`
	QuoteAsset             string                   `json:"quoteAsset"`
	QuotePrecision         int                      `json:"quotePrecision"`
	OrderTypes             []string                 `json:"orderTypes"`
	IcebergAllowed         bool                     `json:"icebergAllowed"`
	OcoAllowed             bool                     `json:"ocoAllowed"`
	IsSpotTradingAllowed   bool                     `json:"isSpotTradingAllowed"`
	IsMarginTradingAllowed bool                     `json:"isMarginTradingAllowed"`
	Filters                []map[string]interface{} `json:"filters"`
}

func main() {
	/*var (
		apiKey = "a9aKL2kh01NABxvpRCYp2D7CpspMMLimbT0g7gPflj6XyQfgBHnxRWxzBVopRt8o"
		secret = "8qCnG7WPKYBSPqFdVZDIlKhkEsP7WFhQLbCnoGnDpcDLtnu54VvhGyutD05tCJ4o"
	)*/
	//res, _ := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=LTCBTC")
	/*exInfo, err := getExchangeInfo()
	//TODO: maybe use decoder
	if err != nil {
		print(err.Error())
		print("Run")
		for _, val := range exInfo.Symbols {
			println(val.Symbol)
		}
	}*/
	//fmt.Printf("%v",getPrice("ETHBTC"))
	exInfo, err := getExchangeInfo()
	if err != nil {
		println(err.Error())

	}
	getCurrPrices(exInfo)
	//fmt.Printf("%f",priceMap["LTCBTC"])
	for key, value := range priceMap {
		fmt.Printf("%v, %f\n", key, value)
	}
	/*for _,symbol := range exInfo.timezone{
		print(symbol)
		//print(string(i.baseAsset),"\n")
	}*/

	/*if err != nil{
		fmt.Printf("", err,"\n")
	}*/
	/*buf := new(bytes.Buffer)
	_ ,err := buf.ReadFrom(res.Body)
	if err != nil{
		print(err)
	}

	outStr := buf.String()
	fmt.Printf(outStr)*/

}