package functions

import (
	"encoding/json"
	"net/http"
)

type btc struct {
	Two4HAvg float64 `json:"24h_avg"`
	Ask float64 `json:"ask"`
	Bid float64 `json:"bid"`
	Last float64 `json:"last"`
	Timestamp string `json:"timestamp"`
	VolumeBtc float64 `json:"volume_btc"`
	VolumePercent float64 `json:"volume_percent"`
}

type bitstampEx struct {
	Name string `json:"name"`
	DisplayName string `json:"display_name"`
	URL string `json:"url"`
	Timestamp int `json:"timestamp"`
	DataSource string `json:"data_source"`
	Symbols struct {
		BTCEUR struct {
			Last float64 `json:"last"`
			Volume float64 `json:"volume"`
			Ask float64 `json:"ask"`
			Bid float64 `json:"bid"`
		} `json:"BTCEUR"`
		BTCUSD struct {
			Last float64 `json:"last"`
			Volume float64 `json:"volume"`
			Ask float64 `json:"ask"`
			Bid float64 `json:"bid"`
		} `json:"BTCUSD"`
	} `json:"symbols"`
}

type latestBlock struct {
	Hash string `json:"hash"`
	Time int `json:"time"`
	BlockIndex int `json:"block_index"`
	Height int `json:"height"`
	TxIndexes []int `json:"txIndexes"`
}


func GetJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func (b *bitstampEx) LastEUR()  float64 {
	return b.Symbols.BTCEUR.Last

}

func (b *bitstampEx) VolumeEUR()  float64 {
	return b.Symbols.BTCEUR.Volume
}

func (b *bitstampEx) AskEUR()  float64 {
	return b.Symbols.BTCEUR.Ask
}

func (b *bitstampEx) BidEUR() float64 {
	return b.Symbols.BTCEUR.Bid
}

func GetBTCUSD(done chan bool) *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/USD/"
	bitUSD := new(btc)
	GetJSON(url, bitUSD)
	done <- true
	return bitUSD
}

func GetBTCEUR(done chan bool) *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/EUR/"
	bitEUR := new(btc)
	GetJSON(url, bitEUR)
	done <- true
	return bitEUR
}

func GetBTCNOK(done chan bool) *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/NOK/"
	bitNOK := new(btc)
	GetJSON(url, bitNOK)
	done <- true
	return bitNOK
}

func GetBTCGBP(done chan bool) *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/GBP/"
	bitGBP := new(btc)
	GetJSON(url, bitGBP)
	done <- true
	return bitGBP
}

func GetBTCDKK(done chan bool) *btc {
	url := "https://api.bitcoinaverage.com/ticker/global/DKK/"
	bitDKK := new(btc)
	GetJSON(url, bitDKK)
	done <- true
	return bitDKK
}


func Getbitstamp(done chan bool) *bitstampEx {
	url := "https://apiv2.bitcoinaverage.com/exchanges/bitstamp"
	bitstamp := new(bitstampEx)
	GetJSON(url, bitstamp)
	done <- true
	return bitstamp
}



func GetLatestBlock(done chan bool) *latestBlock {
	url := "https://blockchain.info/latestblock?format=json"
	block := new(latestBlock)
	GetJSON(url, block)
	done <- true
	return block
}


