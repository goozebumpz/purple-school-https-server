package random

import (
	"net/http"
)

type Handler struct {
	Service
}

func NewRandomHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/random", handler.Random())
}

func (handler *Handler) Random() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		randomNumber := handler.Service.Random()
		w.Write([]byte(randomNumber))
	}
}
