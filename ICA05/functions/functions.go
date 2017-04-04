package functions

import (
	"encoding/json"
	"net/http"
)

type btc struct {
	Two4HAvg float64 `json:"avg"`
	Ask float64 `json:"ask"`
	Bid float64 `json:"bid"`
	Last float64 `json:"last"`
	Timestamp string `json:"timestamp"`
	VolumeBtc float64 `json:"volume_btc"`
	VolumePercent float64 `json:"volume_percent"`
}

func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func GetBTCUSD() *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/USD/"
	bitUSD := new(btc)
	GetJSON(url, bitUSD)
	return bitUSD
}

func GetBTCEUR() *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/EUR/"
	bitEUR := new(btc)
	GetJSON(url, bitEUR)
	return bitEUR
}

func GetBTCNOK() *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/NOK/"
	bitNOK := new(btc)
	GetJSON(url, bitNOK)
	return bitNOK
}

func GetBTCGBP() *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/GBP/"
	bitGBP := new(btc)
	GetJSON(url, bitGBP)
	return bitGBP
}

func GetBTCDKK() *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/DKK/"
	bitDKK := new(btc)
	GetJSON(url, bitDKK)
	return bitDKK
}