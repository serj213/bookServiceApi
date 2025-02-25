package http

import (
	"encoding/json"
	"net/http"
)

func (h HTTPServer) Create(w http.ResponseWriter, r *http.Request) {
	var bookReq BookRequest

	if err := json.NewDecoder(r.Body).Decode(&bookReq); err != nil {
		ErrResponse("invalid request", w, r, http.StatusBadRequest)
		return
	}

	if err := bookReq.Validate(); err != nil {
		ErrResponse("invalid request", w, r, http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	book, err := h.BookService.Create(r.Context(), bookReq.Title, bookReq.Author, bookReq.CategoryId)
	if err != nil {
		h.log.Errorf("failed create serivce: %w", err)
		ErrResponse("server error", w, r, http.StatusInternalServerError)
		return
	}

	ResponseOk(BookResponse{
		Id: book.ID,
		Title: book.Title,
		Author: book.Author,
		CategoryId: book.CategoryId,
	}, w)
}