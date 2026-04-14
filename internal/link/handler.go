package link

import (
	"fmt"
	"net/http"
	"purple-school/configs"
	"purple-school/pkg/req"
	"purple-school/pkg/res"
)

type HandlerDeps struct {
	Config     *configs.Config
	Repository *Repository
}

type Handler struct {
	Config     *configs.Config
	Repository *Repository
}

func NewLinksHandler(router *http.ServeMux, deps HandlerDeps) {
	handler := &Handler{
		Config:     deps.Config,
		Repository: deps.Repository,
	}

	router.HandleFunc("GET /{hash}", handler.GoTo())
	router.HandleFunc("POST /link", handler.Create())
	router.HandleFunc("PATCH /link/{id}", handler.Update())
	router.HandleFunc("DELETE /link/{id}", handler.Delete())
}

func (h *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("get many"))
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateRequest](w, r)
		if err != nil {
			return
		}

		link := NewLink(body.Url)

		createdLink, err := h.Repository.CreateLink(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res.JSON(w, createdLink, 201)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("update"))
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
