package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"time"
)

// APIQuery holds the URL and params which is sent to APICall to query rates
type APIQuery struct {
	Url    string `json:"method"`
	Params string `json:"params"`
}

func APICall(q *APIQuery) (string, error) {

	// url := "http://data.fixer.io/api/latest?access_key=YOUR_FIXER_API_KEY&format=1"
	// url := "https://tradeogre.com/api/v1/ticker/BTC-ARRR"
	url := q.Url + `?` + q.Params
	// fmt.Println(q.Url)
	// fmt.Println(url)
	method := "GET"

	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	// fmt.Println(string(body))

	s := string(body)
	return s, nil
}

// RPCResultMap using golang's own http package
func RPCResultMap(url string, params interface{}) (interface{}, error) {

	// fmt.Printf("params -- %+v\n", params)

	paramsStr := fmt.Sprintf("%v", params)

	query := APIQuery{
		Url:    url,
		Params: string(paramsStr),
	}
	// fmt.Printf("%+v\n", query)

	getJSON, err := APICall(&query)
	if err != nil {
		fmt.Println(err)
		return "error", err
	}
	// fmt.Printf("%T\n", getJSON)

	var result interface{}
	json.Unmarshal([]byte(getJSON), &result)
	// fmt.Println("Error:", result["error"])
	// fmt.Println("success:", result["success"])
	// if result.(map[string]interface{})["error"] != nil {
	// 	// fmt.Printf("error is not nil: %v\n", result["error"])
	// 	return nil, errors.New(result.(map[string]interface{})["error"].(string))
	// }
	return result, nil
	// return "", nil
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
