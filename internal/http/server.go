package http

import (
	"context"

	"github.com/serj213/bookServiceApi/internal/domain"
	"go.uber.org/zap"
)

type BookService interface {
	Create(ctx context.Context, title string, author string, categoryId int)(domain.Book, error)
}


type HTTPServer struct {
	log *zap.SugaredLogger
	BookService BookService
}


func New(log *zap.SugaredLogger, bookService BookService) *HTTPServer {
	return &HTTPServer{
		log: log,
		BookService: bookService,
	}
}