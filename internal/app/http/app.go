package httpapp

import (
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/api/router"
	bookusecase "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/usecase/book"
	"github.com/gin-gonic/gin"
	"log/slog"
)

type App struct {
	Logger *slog.Logger
	Port   string
	Server *gin.Engine
}

func NewApp(logger *slog.Logger, port string, bookCase *bookusecase.BookUseCaseImpl) *App {
	r := router.RegisterRouter(bookCase)
	return &App{
		Logger: logger,
		Port:   port,
		Server: r,
	}
}

func (app *App) Start() {

	const op = "app.Start"
	log := app.Logger.With(
		slog.String(op, "Starting server"),
		slog.String("port", app.Port))
	log.Info("Starting server")
	err := app.Server.Run(app.Port)
	if err != nil {
		log.Info("Failed to start server", "error", err)
	}
}
