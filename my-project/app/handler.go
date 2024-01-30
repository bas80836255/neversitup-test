package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// routes configures route handler functions
func routes() *mux.Router {
	r := mux.NewRouter()

	r.StrictSlash(true)
	r.SkipClean(true)
	r.HandleFunc("/", root).Methods(http.MethodGet)

	subRUser := r.PathPrefix("/users").Subrouter()
	subRUser.HandleFunc("/{id}", userHandler.Get).Methods(http.MethodGet)
	return r
}

// root HTTP handler function for service health check
// return header status code 200
func root(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	res := map[string]string{"status": "up"}
	marshal, _ := json.Marshal(res)
	w.Write(marshal)
}
