package main

import (
	"cpa_network/internal/config"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	var cfg *config.Config
	var err error

	if cfg, err = config.ReadConfig("resources/config.json"); err != nil {
		panic("failed to read config file: " + err.Error())
	}

	address := fmt.Sprintf("%s:%d", cfg.Server.Network.Host, cfg.Server.Network.Port)
	router := mux.NewRouter()
	if err = http.ListenAndServe(address, router); err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}
