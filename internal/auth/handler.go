package auth

import (
	"net/http"
	"purple-school/configs"
	"purple-school/pkg/req"
	"purple-school/pkg/res"
)

type HandlerDeps struct {
	Config *configs.Config
}

type Handler struct {
	Config *configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", handler.Login())
	router.HandleFunc("POST /auth/register", handler.Register())
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)

		if err != nil {
			return
		}

		res.JSON(w, body, 200)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)

		if err != nil {
			return
		}

		res.JSON(w, body, 200)
	}
}
