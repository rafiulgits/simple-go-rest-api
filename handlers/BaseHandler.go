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

//NotFound :
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message" : "requested data is not found"}`))
}

//JSONResponse :
func JSONResponse(w http.ResponseWriter, d interface{}) {
	_ = json.NewEncoder(w).Encode(d)
}
