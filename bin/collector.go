package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type localCollector struct {
	gaugeTest *prometheus.Desc
}

func NewCollector() *localCollector {
	return &localCollector{
		gaugeTest: prometheus.NewDesc("gauge_test", "A gauge test", nil, nil),
	}
}

func (col *localCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- col.gaugeTest
}

func (col *localCollector) Collect(ch chan<- prometheus.Metric) {
	var metric float64 = 1

	ch <- prometheus.MustNewConstMetric(col.gaugeTest, prometheus.GaugeValue, metric)
}
