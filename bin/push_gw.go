package main

import (
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	completionTime := prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "tz_time",
			Help: "TZ timestamp",
		},
		[]string{"identifier"},
	)

	completionTime.With(prometheus.Labels{"identifier": "tz_worktime"}).SetToCurrentTime()
	if err := push.New("http://pushgateway:9091", "prometheus").Collector(completionTime).Push(); err != nil {
		fmt.Println("Could not push new gauge.", err)
	}

	fmt.Println("Done.")

}
