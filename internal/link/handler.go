package link

import (
	"gorm.io/gorm"
	"net/http"
	"purple-school/configs"
	"purple-school/pkg/req"
	"purple-school/pkg/res"
	"strconv"
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

		res.JSON(w, createdLink, http.StatusCreated)
	}
}

func (h *Handler) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		body, err := req.HandleBody[UpdateRequest](w, r)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		linkFrom := Link{
			Model: gorm.Model{ID: uint(id)},
			Url:   body.Url,
			Hash:  body.Hash,
		}

		link, err := h.Repository.UpdateLink(&linkFrom)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		res.JSON(w, link, http.StatusOK)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.ParseUint(idStr, 10, 64)

		if err != nil {
			res.JSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = h.Repository.DeleteLink(uint(id))

		if err != nil {
			res.JSON(w, err.Error(), http.StatusInternalServerError)
			return
		}

		res.JSON(w, nil, http.StatusOK)

	}
}
