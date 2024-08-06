package main

import (
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/app"
	"github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/internal/config"
	logger2 "github.com/D1Y0RBEKORIFJONOV/Uyga-vaziva-book/logger"
)

func main() {
	cfg := config.New()
	logger := logger2.SetupLogger(cfg.LogLevel)
	application := app.NewApp(cfg, logger)
	forever := make(chan bool)
	go application.HTTP.Start()
	<-forever
}
