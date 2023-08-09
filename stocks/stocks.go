package stocks

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const PoligonPath = "https://api.polygon.io"
const TickerPath = PoligonPath + "/v3/reference/tickers"

type Stock struct {
	Ticker string `json:"ticker"`
	Name   string `json:"name"`
}

type SearchResult struct {
	Results []Stock `json:"results"`
}

func Config(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	return os.Getenv(key)
}

func SearchTicker(ticker string) []Stock {
	ApiKey := Config("ApiKey")
	GetPath := TickerPath + "?apiKey=" + ApiKey + "&ticker=" + strings.ToUpper(ticker)
	fmt.Printf("%+v\n", GetPath)
	res, err := http.Get(GetPath)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", res)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	data := SearchResult{}
	json.Unmarshal([]byte(string(body)), &data)
	return data.Results
}

const DailyValuesPath = PoligonPath + "/v1/open-close"

type Values struct {
	Open float64 `json:"open"`
	High float64 `json:"high"`
	Low  float64 `json:"low"`
}

func GetDailyValues(ticker string) Values {
	ApiKey := Config("ApiKey")
	GetPath := DailyValuesPath + "/" + strings.ToUpper(ticker) +
		"/2023-08-01?adjusted=true&apiKey=" + ApiKey
	fmt.Printf("%+v\n", GetPath)
	res, err := http.Get(GetPath)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("%+v\n", res)
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}
	data := Values{}
	json.Unmarshal([]byte(string(body)), &data)
	return data
}
