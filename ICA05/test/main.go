package main

import (
"html/template"
"net/http"
"path"
	"is105/is105/ICA05/test/functions"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	btcUSD := functions.GetBTCUSD()
	fp := path.Join("templates", "index.tmpl")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, btcUSD); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}