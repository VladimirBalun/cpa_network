package http

import "github.com/gorilla/mux"

// Handler - stores pointers to CPA service and Router
type Handler struct {
	Router *mux.Router
	// TODO : add CPA service pointer
}

// NewHandler - Handler constructor
func NewHandler() *Handler {
	return &Handler{}
}

// SetupRoutes - sets up all endpoint routes for CPA application
func (h *Handler) SetupRoutes() {
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("/cpa/main/{id}", h.GetMain).Methods("GET")
	h.Router.HandleFunc("/cpa/offers/{id}", h.GetOffers).Methods("GET")
	h.Router.HandleFunc("/cpa/threads/{id}", h.GetThreads).Methods("GET")

	// instruments
	h.Router.HandleFunc("/cpa/creatives/{id}", h.GetCreatives).Methods("GET")
	h.Router.HandleFunc("/cpa/domains/{id}", h.GetDomains).Methods("GET")
	h.Router.HandleFunc("/cpa/finances/{id}", h.GetFinances).Methods("GET")
	h.Router.HandleFunc("/cpa/statistics/{id}", h.GetStatistics).Methods("GET")
	h.Router.HandleFunc("/cpa/instruments/{id}", h.GetInstruments).Methods("GET")

	// additionally
	h.Router.HandleFunc("/cpa/help", h.GetHelp).Methods("GET")
	h.Router.HandleFunc("/cpa/api", h.GetApi).Methods("GET")
	h.Router.HandleFunc("/cpa/postbacks", h.GetPostbacks).Methods("GET")
	h.Router.HandleFunc("/cpa/useful_services", h.GetUsefulServices).Methods("GET")
}
