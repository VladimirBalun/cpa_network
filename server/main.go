package main

import (
	transportHTTP "cpa_network/internal/transport/http"
	"log"
	"net/http"
)

func main() {
	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	err := http.ListenAndServe("127.0.0.1:8080", handler.Router)
	if err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}
