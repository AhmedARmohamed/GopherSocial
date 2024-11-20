package main

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/AhmedARMohamed/social/internal/store"
	"github.com/go-chi/chi/v5"
)

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostHandler(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload
	if err := readJSON(w, r, &payload); err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: Change after auth
		UserID: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusCreated, post); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idParam := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	posts, err := app.store.Posts.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			writeJSONError(w, http.StatusNotFound, err.Error())
		default:
			writeJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := writeJSON(w, http.StatusOK, posts); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (app *application) listPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	posts, err := app.store.Posts.List(ctx)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJSON(w, http.StatusOK, posts); err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
