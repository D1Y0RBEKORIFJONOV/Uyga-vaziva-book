package postgresDB

import (
	"context"
	"fmt"
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
		tableName: "books",
		log:       log,
	}
}

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
func (reop *BookRepository) Returning(data string) string {
	return fmt.Sprintf("RETURNING  %s", data)
}

func (r *BookRepository) UpdateBook(ctx context.Context, req *bookEntity.UpdateBookRequest) (*bookEntity.Book, error) {
	const op = "postgres.SoldiersRepository.UpdateBook"
	log := r.log.With(slog.String("method", op))
	log.Info("updating book start")
	defer log.Info("updating book end")

	date := map[string]interface{}{}
	if req.Author != "" {
		date["author_name"] = req.Author
	}
	if req.Title != "" {
		date["title"] = req.Title
	}
	if req.PublisherDay != "" {
		date["publisher_year"] = req.PublisherDay
	}
	date["id"] = req.ID

	query, args, err := r.db.Sq.Builder.Update(r.tableName).SetMap(date).
		Where(r.db.Sq.Equal("ID", req.ID)).
		Suffix(r.Returning(r.selectQuery())).ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("executing query:", query)
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var book bookEntity.Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Author, &book.PublisherDay, &book.Title)
		if err != nil {
			log.Error(err.Error())
			return nil, err
		}
	}
	if err := rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) GetBook(ctx context.Context, req *bookEntity.FieldValueReq) (*bookEntity.Book, error) {
	const op = "postgres.SoldiersRepository.GetBook"
	log := r.log.With(slog.String("method", op))
	log.Info("getting book start")
	defer log.Info("getting book end")

	query, args, err := r.db.Sq.Builder.Select(r.selectQuery()).Where(r.db.Sq.Equal(req.Field, req.Value)).ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("executing query:", query)
	var book bookEntity.Book
	err = r.db.QueryRow(ctx, query, args...).Scan(
		&book.ID,
		&book.Author,
		&book.PublisherDay,
		&book.Title)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, req *bookEntity.FieldValueReq) error {
	const op = "postgres.SoldiersRepository.DeleteBook"
	log := r.log.With(slog.String("method", op))
	log.Info("deleting book start")
	defer log.Info("deleting book end")
	query, args, err := r.db.Sq.Builder.Delete(r.tableName).Where(r.db.Sq.Equal(req.Field, req.Value)).ToSql()
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

func (r *BookRepository) GetBooks(ctx context.Context, req *bookEntity.GetAllBookReq) ([]*bookEntity.Book, error) {
	const op = "postgres.SoldiersRepository.GetBooks"
	log := r.log.With(slog.String("method", op))
	log.Info("getting books start")
	defer log.Info("getting books end")

	toSql := r.db.Sq.Builder.Select(r.selectQuery()).From(r.tableName)

	if req.Field != "" && req.Value != "" {
		toSql = toSql.Where(r.db.Sq.Equal(req.Field, req.Value))
	}
	if req.Limit > 0 {
		toSql = toSql.Limit(uint64(req.Limit))
	}
	if req.Offset > 0 {
		toSql = toSql.Offset(uint64(req.Offset))
	}

	query, args, err := toSql.ToSql()
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	log.Info("executing query:", query)
	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	defer rows.Close()
	var books []*bookEntity.Book
	for rows.Next() {
		var book bookEntity.Book
		err = rows.Scan(
			&book.ID,
			&book.Author,
			&book.PublisherDay,
			&book.Title)
		books = append(books, &book)
	}
	if err := rows.Err(); err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return books, nil
}
