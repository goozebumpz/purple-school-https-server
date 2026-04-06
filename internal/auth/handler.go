package auth

import "net/http"

type Handler struct{}

func NewAuthHandler(router *http.ServeMux) {
	handler := &Handler{}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (a *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	}
}

func (a *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("register"))
	}
}
