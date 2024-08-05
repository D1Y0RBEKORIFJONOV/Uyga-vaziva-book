package bookusecase

import (
	"context"
	bookEntity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/book"
)

type (
	Saver interface {
		SaveBook(ctx context.Context, req *bookEntity.Book) error
	}
	Updater interface {
		UpdateBook(ctx context.Context, req *bookEntity.UpdateBookRequest) (*bookEntity.Book, error)
	}
	Provider interface {
		GetBook(ctx context.Context, req *bookEntity.FieldValueReq) (*bookEntity.Book, error)
		GetBooks(ctx context.Context, req *bookEntity.GetAllBookReq) ([]*bookEntity.Book, error)
	}
	Deleter interface {
		DeleteBook(ctx context.Context, req *bookEntity.FieldValueReq) error
	}
)

type Repository interface {
	Saver
	Updater
	Provider
	Deleter
}
type RepoBook struct {
	repo Repository
}

func NewRepoBook(repo Repository) *RepoBook {
	return &RepoBook{repo: repo}
}
func (repo *RepoBook) SaveBook(ctx context.Context, req *bookEntity.Book) error {
	return repo.repo.SaveBook(ctx, req)
}
func (repo *RepoBook) UpdateBook(ctx context.Context, req *bookEntity.UpdateBookRequest) (*bookEntity.Book, error) {
	return repo.repo.UpdateBook(ctx, req)
}
func (repo *RepoBook) GetBook(ctx context.Context, req *bookEntity.FieldValueReq) (*bookEntity.Book, error) {
	return repo.repo.GetBook(ctx, req)
}
func (repo *RepoBook) GetBooks(ctx context.Context, req *bookEntity.GetAllBookReq) ([]*bookEntity.Book, error) {
	return repo.repo.GetBooks(ctx, req)
}
func (repo *RepoBook) DeleteBook(ctx context.Context, req *bookEntity.FieldValueReq) error {
	return repo.repo.DeleteBook(ctx, req)
}
