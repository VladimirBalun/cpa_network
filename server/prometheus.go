package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

const ()

var counter = prometheus.NewCounter(
	prometheus.CounterOpts{
		Name: "Total",
		Help: "Total number of requests",
	},
)

func prometheusMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		val := os.Getenv(prometheus_enable)
		enable, err := strconv.ParseBool(val)
		if err == nil && enable {
			counter.Inc()
		}

		next.ServeHTTP(w, r)
	})
}

func init() {
	prometheus.MustRegister(counter)
}
