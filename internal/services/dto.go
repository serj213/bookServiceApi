package services

import (
	bsv1 "github.com/serj213/bookService/pb/grpc/grpc"
	"github.com/serj213/bookServiceApi/internal/domain"
)


func bookToDomain(book *bsv1.BookResponse ) domain.Book {

	return domain.Book{
		ID: int(book.Id),
		Title: book.Title,
		Author: book.Author,
		CategoryId: int(book.CategoryId),
	}

}