package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

const (
	getFunc = "user.handler.Get"
)

// Handler interface
type Handler interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// handler data structure
type handler struct {
	service Service
}

const pathUserID = "id"

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if rec := recover(); rec != nil {
			h.writeResponse(getFunc, w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		}
	}()

	vars := mux.Vars(r)
	id, ok := vars[pathUserID]
	id = strings.Trim(id, " ")
	if !ok || id == "" {
		h.writeResponse(getFunc, w, http.StatusBadRequest, map[string]string{"error": "bad request"})
		fmt.Printf("%s: id is missing in parameters\n", getFunc)
		return
	}

	if id == "999" {
		panic("id cannot 999")
	}

	user, err := h.service.GetUser(r.Context(), id)
	if err != nil {
		h.writeResponse(getFunc, w, http.StatusNotFound, map[string]string{"error": "not found user"})
		return
	}
	h.writeResponse(getFunc, w, http.StatusOK, user)
}

func (h *handler) writeResponse(function string, w http.ResponseWriter, statusCode int, data any) {
	w.Header().Add("Content-Type", "application/json")

	marshal, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("%s: unmarshall error\n", function)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(marshal)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(marshal)
}

func NewHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}
