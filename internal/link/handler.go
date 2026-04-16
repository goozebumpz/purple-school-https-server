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
	router.HandleFunc("PATCH /link/{hash}", handler.Update())
	router.HandleFunc("DELETE /link/{hash}", handler.Delete())
}

func (h *Handler) GoTo() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		link, err := h.Repository.GetByHash(hash)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Redirect(w, r, link.Url, http.StatusPermanentRedirect)
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, err := req.HandleBody[CreateRequest](w, r)

		if err != nil {
			return
		}

		link := NewLink(body.Url)
		isUnique := h.Repository.CheckUniqueHash(link.Hash)

		for !isUnique {
			link.GenerateNewHash()
			isUnique = h.Repository.CheckUniqueHash(link.Hash)
		}

		createdLink, err := h.Repository.CreateLink(link)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JSON(w, createdLink, 201)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		body, err := req.HandleBody[UpdateRequest](w, r)

		if err != nil {
			res.JSON(w, err.Error(), 400)
			return
		}

		link, err := h.Repository.UpdateLink(hash, body)

		if err != nil {
			res.JSON(w, err.Error(), 400)
			return
		}

		res.JSON(w, link, 200)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Println(id)
	}
}
