package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

//IHandler :
type IHandler interface {
	Handle(router chi.Router)
}

//NotFound : HTTP 404 response
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "requested data is not found"}`))
}

//Ok : HTTP 200 response
func Ok(w http.ResponseWriter, d interface{}) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(d)
}
