package bookusecase

import (
	"context"
	bookeentity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/book"
)

type BookUseCase interface {
	CreateBook(ctx context.Context, req *bookeentity.CreateBookRequest) (*bookeentity.Book, error)
	UpdateBook(ctx context.Context, req *bookeentity.UpdateBookRequest) (*bookeentity.Book, error)
	GetBook(ctx context.Context, req *bookeentity.FieldValueReq) (*bookeentity.Book, error)
	GetBookAll(ctx context.Context, req *bookeentity.GetAllBookReq) (*[]bookeentity.Book, error)
	DeleteBook(ctx context.Context, req *bookeentity.FieldValueReq) error
}

type BookUseCaseImpl struct {
	book BookUseCase
}

func NewBookUseCase(book BookUseCase) *BookUseCaseImpl {
	return &BookUseCaseImpl{
		book: book,
	}
}

func (b *BookUseCaseImpl) CreateBook(ctx context.Context, req *bookeentity.CreateBookRequest) (*bookeentity.Book, error) {
	return b.book.CreateBook(ctx, req)
}

func (b *BookUseCaseImpl) UpdateBook(ctx context.Context, req *bookeentity.UpdateBookRequest) (*bookeentity.Book, error) {
	return b.book.UpdateBook(ctx, req)
}
func (b *BookUseCaseImpl) DeleteBook(ctx context.Context, req *bookeentity.FieldValueReq) error {
	return b.book.DeleteBook(ctx, req)
}
func (b *BookUseCaseImpl) GetBook(ctx context.Context, req *bookeentity.FieldValueReq) (*bookeentity.Book, error) {
	return b.book.GetBook(ctx, req)
}
func (b *BookUseCaseImpl) GetBookAll(ctx context.Context, req *bookeentity.GetAllBookReq) (*[]bookeentity.Book, error) {
	return b.book.GetBookAll(ctx, req)
}
