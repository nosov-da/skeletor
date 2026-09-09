package app

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"github.com/yourorg/yourproject/internal/config"
	"github.com/yourorg/yourproject/internal/server"
)

type App struct {
	HTTPServer *http.Server
	log        *slog.Logger
	closers    []func() error
}

func New(ctx context.Context, cfg *config.Config, log *slog.Logger) (*App, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	a := &App{log: log}

	// Пример инициализации зависимостей:
	//
	// db, err := database.New(ctx, cfg.DB.DSN())
	// if err != nil {
	//     return nil, fmt.Errorf("connect db: %w", err)
	// }
	// a.closers = append(a.closers, db.Close)
	//
	// userRepo := repository.NewUserRepository(db)
	// userSvc := service.NewUserService(userRepo)
	// userHandler := handler.NewUserHandler(userSvc)

	router := server.NewRouter(log)

	a.HTTPServer = &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return a, nil
}

func (a *App) Shutdown(ctx context.Context) error {
	a.log.Info("shutting down http server")
	serverErr := a.HTTPServer.Shutdown(ctx)

	for _, close := range a.closers {
		if err := close(); err != nil {
			a.log.Error("failed to close resource", "err", err)
		}
	}

	return serverErr
}
