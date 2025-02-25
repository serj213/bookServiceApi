package services

import (
	"context"

	bsv1 "github.com/serj213/bookService-contract/gen/go/bookService"
	"github.com/serj213/bookServiceApi/internal/domain"
	"go.uber.org/zap"
)

type BookService struct {
	log *zap.SugaredLogger
	grpc bsv1.BookClient
}

func New(log *zap.SugaredLogger, client bsv1.BookClient) *BookService{
	return &BookService{
		log: log,
		grpc: client,
	}	
}

func (s BookService) Create(ctx context.Context, title string, author string, categoryId int)(domain.Book, error) {

	req := &bsv1.BookCreateRequest{
		Title: title,
		Author: author,
		CategoryId: int64(categoryId),
	}

	// Где то здесь можно использовать кафку

	book, err := s.grpc.Create(ctx, req)
	if err != nil {
		s.log.Errorf("failed grpc create: %w", err.Error())
		return domain.Book{}, err
	}

	return domain.Book{
		ID: int(book.Id),
		Title: book.Title,
		Author: book.Author,
		CategoryId: int(book.CategoryId),
	}, nil

}