package main

import (
	"net/http"
	// add prometheus
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello MBA Full Cycle!!!!!!!</h1>"))
	})
	// add promhttp.Handler()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
