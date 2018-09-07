package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var url string = "http://192.168.0.13:9090/"
var gauge prometheus.Gauge

func main() {
}

func startServer() {
	fmt.Println("Starting listener")
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)

	gauge.Add(30)
}

func test() {
	opts := prometheus.GaugeOpts{
		Name: "gauge_test",
		Help: "Gauge for testing.",
	}
	gauge = prometheus.NewGauge(opts)

	gauge.Set(100)
	prometheus.MustRegister(gauge)

	col := newCollector()
	col.Collect()

	startServer()
}
