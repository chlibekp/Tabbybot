package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"tabbybot/internal/config"
	"tabbybot/internal/discord"
	"tabbybot/internal/http"
	"tabbybot/internal/util"
)

func main() {
	if config.ENV == "DEV" {
		slog.Info("Running in DEV mode.. Debug log level")
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

	httpServer := http.NewServer()

	httpServer.Start()

	go discord.StartBot()

	// Graceful shutdown
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-sc
	util.GracefulShutdown(httpServer)
}
