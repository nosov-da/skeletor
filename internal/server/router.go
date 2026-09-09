package server

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/yourorg/yourproject/internal/handler"
)

func NewRouter(log *slog.Logger) http.Handler {
	log.Debug("initializing router")

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	r.Get("/healthz", handler.Health)
	r.Get("/readyz", handler.Ready)

	r.Route("/api/v1", func(r chi.Router) {
		// Пример подключения бизнес-хэндлеров:
		// r.Mount("/users", userHandler.Routes())
	})

	return r
}
