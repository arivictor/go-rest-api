package comment

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Handler struct {
	service *Service
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

func (h Handler) Get(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id := parameters["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	c, err := h.service.Get(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&c)
}

func (h Handler) Update(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id := parameters["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	var comment Comment
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err = h.service.Update(request.Context(), id, comment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&comment)
}

func (h Handler) Create(writer http.ResponseWriter, request *http.Request) {
	var comment Comment
	err := json.NewDecoder(request.Body).Decode(&comment)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}
	comment, err = h.service.Create(request.Context(), comment)
	if err != nil {

	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&comment)
}

func (h Handler) List(writer http.ResponseWriter, request *http.Request) {
	comments, err := h.service.List(request.Context())
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&comments)
}

func (h Handler) Delete(writer http.ResponseWriter, request *http.Request) {
	parameters := mux.Vars(request)
	id := parameters["id"]
	if id == "" {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	comment, err := h.service.Delete(request.Context(), id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(&comment)
}
