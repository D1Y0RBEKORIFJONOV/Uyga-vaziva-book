package postgres

import (
	"context"
	bookEntity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/book"
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/postgres"
	"log/slog"
)

type BookRepository struct {
	db        *postgres.PostgresDB
	tableName string
	log       *slog.Logger
}

func NewSolderRepository(db *postgres.PostgresDB, log *slog.Logger) *BookRepository {
	return &BookRepository{
		db:        db,
		tableName: "users",
		log:       log,
	}
}

//type (
//	Saver interface {
//		SaveBook(ctx context.Context, req *bookEntity.Book) error
//	}
//	Updater interface {
//		UpdateBook(ctx context.Context, req *bookEntity.UpdateBookRequest) (*bookEntity.Book, error)
//	}
//	Provider interface {
//		GetBook(ctx context.Context, req *bookEntity.FieldValueReq) (*bookEntity.Book, error)
//		GetBooks(ctx context.Context, req *bookEntity.GetAllBookReq) ([]*bookEntity.Book, error)
//	}
//	Deleter interface {
//		DeleteBook(ctx context.Context, req *bookEntity.FieldValueReq) error
//	}
//)

func (r *BookRepository) selectQuery() string {
	return `
	id,
	author_name,
	publisher_year,
	title
`
}

func (r *BookRepository) SaveBook(ctx context.Context, req *bookEntity.Book) error {
	const op = "postgres.SoldiersRepository.SaveBook"
	log := r.log.With(slog.String("method", op))

	log.Info("saving book start")
	defer log.Info("saving book end")

	data := map[string]interface{}{
		"id":             req.ID,
		"author_name":    req.Author,
		"publisher_year": req.PublisherDay,
		"title":          req.Title,
	}

	query, args, err := r.db.Sq.Builder.Insert(r.tableName).SetMap(data).ToSql()
	if err != nil {
		log.Error(err.Error())
		return err
	}

	log.Info("executing query:", query)
	_, err = r.db.Exec(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	return nil
}
