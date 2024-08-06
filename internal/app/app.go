package app

import (
	httpapp "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/app/http"
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/config"
	postgresDB "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/infastructura/repository/postgres"
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/postgres"
	bookservice "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/service/book"
	bookusecase "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/usecase/book"
	"log/slog"
)

type App struct {
	HTTP *httpapp.App
}

func NewApp(cfg *config.Config, logger *slog.Logger) *App {
	db, err := postgres.New(cfg)
	if err != nil {
		panic(err)
	}
	dbPostgres := postgresDB.NewSolderRepository(db, logger)
	repo := bookusecase.NewRepoBook(dbPostgres)

	service := bookservice.NewBook(logger, repo)

	handlerCase := bookusecase.NewBookUseCase(service)
	handler := httpapp.NewApp(logger, cfg.RPCPort, handlerCase)
	return &App{
		HTTP: handler,
	}
}
