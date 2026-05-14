package link

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"purple-school/configs"
	"purple-school/pkg/middleware"
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
	router.Handle("GET /link", middleware.Token(handler.GetLinks(), deps.Config))
	router.Handle("POST /link", middleware.Token(handler.Create(), deps.Config))
	router.Handle("PATCH /link/{id}", middleware.Token(handler.Update(), deps.Config))
	router.Handle("DELETE /link/{id}", middleware.Token(handler.Delete(), deps.Config))
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

		if parsedEmail, ok := r.Context().Value(middleware.ContextEmailKey).(string); ok {
			fmt.Println(parsedEmail)
		}

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

		exists, err := h.Repository.ExistLink(uint(id))

		if err != nil {
			res.JSON(w, err.Error(), http.StatusNotFound)
			return
		}

		if exists == false {
			res.JSON(w, "Link not exists", http.StatusNotFound)
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

func (h *Handler) GetLinks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queries := r.URL.Query()
		limit, err := strconv.Atoi(queries.Get("limit"))

		if err != nil {
			res.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		offset, err := strconv.Atoi(queries.Get("offset"))

		if err != nil {
			res.JSON(w, err.Error(), http.StatusBadRequest)
			return
		}

		links := h.Repository.GetLinks(uint(limit), uint(offset))
		res.JSON(w, links, http.StatusOK)
	}
}
