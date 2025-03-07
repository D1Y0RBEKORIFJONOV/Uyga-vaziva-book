package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/config"
	err_entity "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/entity/errors"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"strings"
)

type PostgresDB struct {
	*pgxpool.Pool
	Sq Squirrel
}

func New(config *config.Config) (*PostgresDB, error) {
	var db PostgresDB

	db.Sq = *NewSquirrel()

	if err := db.connectDB(config); err != nil {
		return nil, err
	}

	return &db, nil
}

func (p *PostgresDB) connectDB(config *config.Config) error {
	pgxConfig, err := pgxpool.ParseConfig(p.configToString(config))
	if err != nil {
		return fmt.Errorf("unable to parse db conifg: %s", err.Error())
	}

	pgxPool, err := pgxpool.ConnectConfig(context.Background(), pgxConfig)
	if err != nil {
		return fmt.Errorf("unable to connect to db: %s", err.Error())
	}

	p.Pool = pgxPool

	return nil
}

func (p *PostgresDB) configToString(config *config.Config) string {
	var conn []string
	if len(config.DB.Host) != 0 {
		conn = append(conn, "host="+config.DB.Host)
	}

	if len(config.DB.Port) != 0 {
		conn = append(conn, "port="+config.DB.Port)
	}

	if len(config.DB.User) != 0 {
		conn = append(conn, "user="+config.DB.User)
	}

	if len(config.DB.Password) != 0 {
		conn = append(conn, "password="+config.DB.Password)
	}

	if len(config.DB.Name) != 0 {
		conn = append(conn, "dbname="+config.DB.Name)
	}

	if len(config.DB.SslMode) != 0 {
		conn = append(conn, "sslmode="+config.DB.SslMode)
	}

	return strings.Join(conn, " ")
}

func (p *PostgresDB) Close() {
	p.Pool.Close()
}

func (p *PostgresDB) Error(err error) error {
	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return err_entity.ErrInternal
		}
	}

	if err == pgx.ErrNoRows {
		return err_entity.ErrUserDeleted
	}
	return err
}

func (p *PostgresDB) ErrSQLBuild(err error, message string) error {
	return fmt.Errorf("error during sql build, %s: %w", message, err)
}
