package http

import (
	"log/slog"
	"net/http"
	"tabbybot/internal/config"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Server struct {
	Server *http.Server
	Mux    *http.ServeMux
}

func (s *Server) Start() {
	slog.Info("HTTP Server listening", "port", config.HTTP_PORT)
	go func() {
		// Start the server, after the server closes, log the close reason
		if err := s.Server.ListenAndServe(); err != nil {
			slog.Info("HTTP Server has been closed", "error", err.Error())
		}
	}()
}

func (s *Server) Close() error {
	slog.Info("Closing HTTP server...")

	// Close the HTTP server
	err := s.Server.Close()
	if err != nil {
		return err
	}
	slog.Info("HTTP Server has been closed")
	return nil
}

func NewServer() *Server {
	mux := http.NewServeMux()

	// Handle /metrics prometheus endpoint
	mux.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	// Create a new server
	server := http.Server{
		Addr:    ":" + config.HTTP_PORT,
		Handler: http.HandlerFunc(mux.ServeHTTP),
	}

	return &Server{
		Server: &server,
		Mux:    mux,
	}
}
