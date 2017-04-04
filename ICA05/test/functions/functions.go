
package functions

import (
"encoding/json"
"net/http"
)

type btcUSD struct {
	Two4HAvg float64 `json:"24h_avg"`
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
func GetBTCUSD() *btcUSD {
	url := "https://api.bitcoinaverage.com/ticker/global/USD/"
	bitUSD := new(btcUSD)
	GetJSON(url, bitUSD)
	return bitUSD
}
