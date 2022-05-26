package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func simpleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	response := fmt.Sprintf("Test %s", id)
	fmt.Fprint(w, response)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test/{id:[0-9]+}", simpleHandler)
	log.Println("Server started. Listening...")
	err := http.ListenAndServe(":8181", router)
	if err != nil {
		log.Fatalf("failed to listen and serve: %s", err.Error())
		panic("failed to listen and serve: " + err.Error())
	}
}
