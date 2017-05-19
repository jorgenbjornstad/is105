package main
import (
	"net/http"

	//"strconv"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"encoding/json"
)

type btcMap map[string]interface{}


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

var c = make(chan int)
var done = make(chan bool)
var bitUSD *btc
var bitEUR *btc
var bitGBP *btc
var bitNOK *btc
var bitDKK *btc
var bitstamp *bitstampEx
var block *latestBlock
func main() {
	m := martini.Classic()

	m.Use( render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))

	m.Get("/", func(r render.Render, x *http.Request) {


		for i := 0; i < 1; i++{
			go GetBTCUSD()
		}
		for i := 0; i < 1; i++{
			go GetBTCGBP()
		}

		for i := 0; i < 1; i++{
			go GetBTCEUR()
		}
		for i := 0; i < 1; i++{
			go GetBTCNOK()
		}
		for i := 0; i < 1; i++{
			go GetBTCDKK()
		}
		for i := 0; i < 1; i++{
			go Getbitstamp()
		}
		for i := 0; i < 1; i++{
			go GetLatestBlock()

		}

		<- done


		r.HTML(200, "index", btcMap{
			"btcUSD": bitUSD,
			"btcGBP": bitGBP,
			"btcEUR": bitEUR,
			"btcNOK": bitNOK,
			"btcDKK": bitDKK,
			"bitstamp": bitstamp,
			"block": block,


		})
	})




	m.RunOnAddr(":8001")
	m.Run()
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

func GetBTCUSD()  {

	url := "https://api.bitcoinaverage.com/ticker/global/USD/"
	bitUSD = new(btc)
	GetJSON(url, bitUSD)
	done <- true
	<-c

}

func GetBTCEUR()  {

	url := "https://api.bitcoinaverage.com/ticker/global/EUR/"
	bitEUR = new(btc)
	GetJSON(url, bitEUR)
	done <- true
	<-c
}

func GetBTCNOK () {
	url := "https://api.bitcoinaverage.com/ticker/global/NOK/"
	bitNOK = new(btc)
	GetJSON(url, bitNOK)
	done <- true
	<-c
}

func GetBTCGBP()  {
	url := "https://api.bitcoinaverage.com/ticker/global/GBP/"
	bitGBP = new(btc)
	GetJSON(url, bitGBP)
	done <- true
	<-c

}

func GetBTCDKK()  {
	url := "https://api.bitcoinaverage.com/ticker/global/DKK/"
	bitDKK = new(btc)
	GetJSON(url, bitDKK)
	done <- true
	<-c

}


func Getbitstamp()  {
	url := "https://apiv2.bitcoinaverage.com/exchanges/bitstamp"
	bitstamp = new(bitstampEx)
	GetJSON(url, bitstamp)
	done <- true
	<-c

}



func GetLatestBlock()  {
	url := "https://blockchain.info/latestblock?format=json"
	block = new(latestBlock)
	GetJSON(url, block)
	done <- true
	<-c

}


