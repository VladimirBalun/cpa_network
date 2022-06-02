package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter()
	router.Use(prometheusMiddleware)
	router.Path("/metrics").Handler(promhttp.Handler())
	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}
