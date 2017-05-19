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
		done := make(chan bool, 7)
		btcUSD := functions.GetBTCUSD(done)
		btcGBP := functions.GetBTCGBP(done)
		btcEUR := functions.GetBTCEUR(done)
		btcNOK := functions.GetBTCNOK(done)
		btcDKK := functions.GetBTCDKK(done)
		bitstamp := functions.Getbitstamp(done)
		block := functions.GetLatestBlock(done)

		<- done



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




	m.RunOnAddr(":8001")
	m.Run()
}



