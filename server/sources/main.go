package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	err := http.ListenAndServe("127.0.0.1:8080", router)
	if err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}
