package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/serj213/bookServiceApi/internal/domain"
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
		if errors.Is(err, domain.ErrBookExist){
			ErrResponse("book is exists", w, r, http.StatusInternalServerError)
			return
		}
		ErrResponse("server error", w, r, http.StatusInternalServerError)
		return
	}

	ResponseOk(BookResponse{
		Id: book.ID,
		Title: book.Title,
		Author: book.Author,
		CategoryId: book.CategoryId,
		CreatedAt: book.CreatedAt,
	}, w)
}

func (h HTTPServer) GetBooks(w http.ResponseWriter, r *http.Request) {
	
	books, err := h.BookService.GetBooks(r.Context())
	if err != nil {
		ErrResponse("internal", w, r, http.StatusInternalServerError)
		return 
	}

	resBooks := make([]BookResponse, len(books))

	for i, book := range books {

		if book.UpdatedAt != nil {
			resBooks[i].UpdatedAt = book.UpdatedAt
		}
		
		resBooks[i] = BookResponse{
			Id: book.ID,
			Title: book.Title,
			Author: book.Author,
			CategoryId: book.CategoryId,
			CreatedAt: book.CreatedAt,
		}
	}

	resOk := GetBooksResponseOk{
		Status: "success",
		Books: resBooks,
	}

	ResponseOk(resOk, w)

}