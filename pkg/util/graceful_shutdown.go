package util

import (
	"log/slog"
	"tabbybot/pkg/http"
)

func GracefulShutdown(httpServer *http.Server) {
	httpServer.Close()
	slog.Info("Received signal, shutting down...")
}
