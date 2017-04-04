package main
import (
	"net/http"
	"path"
	"./functions"
	"html/template"
)

func BTCUSD(w http.ResponseWriter, r *http.Request) {
	btc := functions.GetBTCUSD()
	fp := path.Join("templates", "index.tmpl")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, btc); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", BTCUSD)
	http.ListenAndServe(":8001", nil)
}