package main

import (
	"encoding/json"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Response struct {
	Message string `json:"message"`
}

var (
	helloCounter = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "hello_requests_total",
			Help: "Number of hello endpoint calls",
		},
	)
)

func hello(w http.ResponseWriter, r *http.Request) {

	helloCounter.Inc()

	w.Header().Set("Content-Type", "application/json")

	resp := Response{
		Message: "Hello World",
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {

	prometheus.MustRegister(helloCounter)

	http.HandleFunc("/hello", hello)

	// Prometheus metrics endpoint
	http.Handle("/metrics", promhttp.Handler())

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
