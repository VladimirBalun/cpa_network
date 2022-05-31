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
	// TODO: implement CPA endpoints routes
}
