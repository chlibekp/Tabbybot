package util

import (
	"log/slog"
)

func GracefulShutdown() {

	slog.Info("Received signal, shutting down...")
}
