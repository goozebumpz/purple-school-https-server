package hello

import "net/http"

type Handler struct{}

func NewHelloHandler(router *http.ServeMux) {
	handler := &Handler{}
	router.HandleFunc("/hello", handler.Hello())
}

func (handler *Handler) Hello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello dick"))
	}
}
