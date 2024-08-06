package bookservice

import (
	"context"
	bookentity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/book"
	bookusecase "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/usecase/book"
	"github.com/google/uuid"
	"log/slog"
)

type Book struct {
	logger *slog.Logger
	repo   *bookusecase.RepoBook
}

func NewBook(logger *slog.Logger, repo *bookusecase.RepoBook) *Book {
	return &Book{logger: logger, repo: repo}
}

func (b *Book) CreateBook(ctx context.Context, req *bookentity.CreateBookRequest) (*bookentity.Book, error) {
	const op = "Service.Book.CreateBook"
	log := b.logger.With(
		slog.String("method", op),
	)

	log.Info("start book")
	defer log.Info("end book")

	book := &bookentity.Book{
		Title:  req.Title,
		Author: req.Author,
		ID:     uuid.NewString(),
	}

	log.Info("call SaveBook")
	err := b.repo.SaveBook(ctx, book)
	if err != nil {
		log.Error("call SaveBook fail")
		return nil, err
	}
	return book, nil
}

func (b *Book) UpdateBook(ctx context.Context, req *bookentity.UpdateBookRequest) (*bookentity.Book, error) {
	const op = "Service.Book.UpdateBook"
	log := b.logger.With(
		slog.String("method", op))

	log.Info("start book")
	defer log.Info("end book")

	log.Info("Call UpdateBook")
	book, err := b.repo.UpdateBook(ctx, req)
	if err != nil {
		log.Error("call UpdateBook fail")
		return nil, err
	}
	return book, nil
}

func (b *Book) GetBook(ctx context.Context, req *bookentity.FieldValueReq) (*bookentity.Book, error) {
	const op = "Service.Book.GetBook"
	log := b.logger.With(
		slog.String("method", op))

	log.Info("start book")
	defer log.Info("end book")

	log.Info("call GetBook")
	book, err := b.repo.GetBook(ctx, req)
	if err != nil {
		log.Error("call GetBook fail")
		return nil, err
	}

	return book, nil
}

func (b *Book) DeleteBook(ctx context.Context, req *bookentity.FieldValueReq) error {
	const op = "Service.Book.DeleteBook"
	log := b.logger.With(
		slog.String("method", op))

	log.Info("start book")
	defer log.Info("end book")

	log.Info("Call DeleteBook")
	err := b.repo.DeleteBook(ctx, req)
	if err != nil {
		log.Error("call DeleteBook fail")
		return err
	}
	return nil
}

func (b *Book) GetBookAll(ctx context.Context, req *bookentity.GetAllBookReq) ([]*bookentity.Book, error) {
	const op = "Service.Book.GetBookAll"
	log := b.logger.With(
		slog.String("method", op))

	log.Info("start book")
	defer log.Info("end book")

	log.Info("call GetBookAll")
	books, err := b.repo.GetBooks(ctx, req)
	if err != nil {
		log.Error("call GetBookAll fail")
		return nil, err
	}
	return books, err
}
