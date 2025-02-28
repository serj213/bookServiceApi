package services

import (
	"context"
	"errors"
	"fmt"

	bsv1 "github.com/serj213/bookService/pb/grpc"
	"github.com/serj213/bookServiceApi/internal/domain"
	"github.com/serj213/bookServiceApi/internal/lib"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/emptypb"
)

var ErrBookExist = errors.New("book is exist")

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
		s.log.Errorf("failed grpc create: %s", err.Error())
		if lib.GetDescGrpcErr(err) != ""{
			return domain.Book{}, domain.ErrBookExist
		}
		return domain.Book{}, fmt.Errorf("server error")
	}

	s.log.Info("book create successful")

	return domain.Book{
		ID: int(book.Id),
		Title: book.Title,
		Author: book.Author,
		CategoryId: int(book.CategoryId),
		CreatedAt: book.CreatedAt.AsTime(),
	}, nil
}


func (s BookService) GetBooks(ctx context.Context) ([]domain.Book, error) {
	log := s.log.With(zap.String("method", "GetBooks"))

	log.Info("get books start")
	res, err := s.grpc.GetBooks(ctx, &emptypb.Empty{})
	if err != nil {
		log.Error(err)
		return []domain.Book{}, fmt.Errorf("server error")
	}

	books := make([]domain.Book, len(res.Books))

	for i, book := range res.Books{
		books[i] = bookToDomain(book)
	}

	log.Info("get books finish")
	return books, nil
}

func (s BookService) UpdateBook(ctx context.Context, book domain.Book) (domain.Book, error) {

	log := s.log.With(zap.String("method", "UpdateBook"))

	log.Info("update book active")

	req := &bsv1.BookRequest{
		Id: int64(book.ID),
		Title: book.Title,
		Author: book.Author,
		CategoryId: int64(book.CategoryId),
	}

	pbBook, err := s.grpc.UpdateBook(ctx, req)
	if err != nil {
		return domain.Book{}, err
	}

	respBook := bookToDomain(pbBook)
	

	return respBook, nil

}