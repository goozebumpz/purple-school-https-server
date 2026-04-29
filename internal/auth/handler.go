package auth

import (
	"net/http"
	"purple-school/configs"
	"purple-school/pkg/jwt"
	"purple-school/pkg/middleware"
	"purple-school/pkg/req"
	"purple-school/pkg/res"
)

type HandlerDeps struct {
	Config      *configs.Config
	AuthService *Service
}

type Handler struct {
	Config      *configs.Config
	AuthService *Service
}

func NewAuthHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:      deps.Config,
		AuthService: deps.AuthService,
	}

	middlewares := middleware.ChainMiddlewares(middleware.CORS, middleware.Logging)

	router.Handle("POST /auth/login", middlewares(handler.Login()))
	router.Handle("POST /auth/register", middlewares(handler.Register()))
}

func (h *Handler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[LoginRequest](w, r)

		if err != nil {
			res.JSON(w, err.Error(), 400)
			return
		}

		user, err := h.AuthService.Login(body.Email, body.Password)

		if err != nil {
			res.JSON(w, err.Error(), 400)
			return
		}

		token, err := jwt.NewJWT(h.Config.AuthConfig.Secret).Create(user.Email)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSON(w, LoginResponse{Token: token}, 200)
	}
}

func (h *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[RegisterRequest](w, r)

		if err != nil {
			return
		}

		user, err := h.AuthService.Register(body.Email, body.Name, body.Password)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		token, err := jwt.NewJWT(h.Config.AuthConfig.Secret).Create(user.Email)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSON(w, RegisterResponse{Token: token}, 200)
	}
}
