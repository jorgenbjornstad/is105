package main
import (
	"net/http"
	"./functions"
	//"strconv"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

type btcMap map[string]interface{}

func main() {
	m := martini.Classic()

	m.Use( render.Renderer(render.Options{
		IndentJSON: true, // so we can read it..
	}))

	m.Get("/", func(r render.Render, x *http.Request) {
		btcUSD := functions.GetBTCUSD()
		btcGBP := functions.GetBTCGBP()
		btcEUR := functions.GetBTCEUR()
		btcNOK := functions.GetBTCNOK()
		btcDKK := functions.GetBTCDKK()
		bitstamp := functions.Getbitstamp()
		block := functions.GetLatestBlock()



		r.HTML(200, "index", btcMap{
			"btcUSD": btcUSD,
			"btcGBP": btcGBP,
			"btcEUR": btcEUR,
			"btcNOK": btcNOK,
			"btcDKK": btcDKK,
			"bitstamp": bitstamp,
			"block": block,


		})
	})




	m.RunOnAddr(":8002")
	m.Run()
}



