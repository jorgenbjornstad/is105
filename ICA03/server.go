// Filen for å eksperimentere med webserver i ICA03 (IS-105)
// Starte server med
// 			go run server.go
// Du kan ha tilgang til server fra nettleser med
// 			http://localhost:3000
//
package main

import (
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", foo)
	http.ListenAndServe(":3000", nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	// Her kan man bl. a. skrive data som skal sendes til nettleser
	// Neste linje kan brukes for å endre koding av sekvensen som sendes til nettleser
	// Standardinstilling er UTF-8
	//w.Header().Set("Content-Type", "text/html;charset=ISO-8859-1")

	// Her skriver man data som er respons til brukeren som har skrevet
	// http://localhost:3000 i sin nettleser
	i := time.Now().Format(time.RFC1123Z)
	w.Write([]byte("<font color=\"green\">Hvordan  g&#229;r det, <b>\u16a6</b> ?\u23f0</font><br/>"))
	w.Write([]byte("\u16a6 - Thurs<br/>"))
	w.Write([]byte(i))
	w.Write([]byte("<br/>git checkout this clock yo &#128526;&#128512;&#128521;&#128521;&#128521;&#128521;&#128521;&#128521;&#128521;&#128521;<br/>"))


}
